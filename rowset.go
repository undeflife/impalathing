package impalathing

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/koblas/impalathing/services/beeswax"
	impala "github.com/koblas/impalathing/services/impalaservice"
)

type rowSet struct {
	client  *impala.ImpalaServiceClient
	handle  *beeswax.QueryHandle
	options Options

	// columns    []*tcliservice.TColumnDesc
	columnNames []string

	offset  int
	rowSet  *beeswax.Results
	hasMore bool
	ready   bool

	metadata *beeswax.ResultsMetadata

	nextRow []string
}

// A RowSet represents an asyncronous hive operation. You can
// Reattach to a previously submitted hive operation if you
// have a valid thrift client, and the serialized Handle()
// from the prior operation.
type RowSet interface {
	Columns() []string
	ColumnsWithContext(ctx context.Context) []string
	Next() bool
	NextWithContext(ctx context.Context) bool
	Scan(dest ...interface{}) error
	Poll() (*Status, error)
	PollWithContext(ctx context.Context) (*Status, error)
	Wait() (*Status, error)
	WaitWithContext(ctx context.Context) (*Status, error)
	FetchAll() []map[string]interface{}
	FetchAllWithContext(ctx context.Context) []map[string]interface{}
	MapScan(dest map[string]interface{}) error
}

// Represents job status, including success state and time the
// status was updated.
type Status struct {
	state beeswax.QueryState
	Error error
}

func newRowSet(client *impala.ImpalaServiceClient, handle *beeswax.QueryHandle, options Options) RowSet {
	return &rowSet{client: client, handle: handle, options: options, columnNames: nil, offset: 0, rowSet: nil,
		hasMore: true, ready: false, metadata: nil, nextRow: nil}
}

//
//
//
func (s *Status) IsSuccess() bool {
	return s.state != beeswax.QueryState_EXCEPTION
}

func (s *Status) IsComplete() bool {
	return s.state == beeswax.QueryState_FINISHED
}

func (r *rowSet) poll(ctx context.Context) (*Status, error) {
	state, err := r.client.GetState(ctx, r.handle)
	if err != nil {
		return nil, fmt.Errorf("Error getting status: %v", err)
	}

	if state == beeswax.QueryState_EXCEPTION {
		return nil, fmt.Errorf("Exception on Impala side")
	}

	return &Status{state, nil}, nil
}

// Issue a thrift call to check for the job's current status.
func (r *rowSet) Poll() (*Status, error) {
	return r.poll(context.Background())
}

// Same as the Poll() but additionally takes the context.
func (r *rowSet) PollWithContext(ctx context.Context) (*Status, error) {
	return r.poll(ctx)
}

func (r *rowSet) wait(ctx context.Context) (*Status, error) {
	for {
		status, err := r.poll(ctx)

		if err != nil {
			return nil, err
		}

		if status.IsComplete() {
			if status.IsSuccess() {
				r.ready = true
				return status, nil
			}
			return nil, fmt.Errorf("Query failed execution: %s", status.state.String())
		}

		time.Sleep(time.Duration(r.options.PollIntervalSeconds) * time.Second)
	}
}

// Wait until the job is complete, one way or another, returning Status and error.
func (r *rowSet) Wait() (*Status, error) {
	return r.wait(context.Background())
}

// Same as the Wait() but additionally takes the context.
func (r *rowSet) WaitWithContext(ctx context.Context) (*Status, error) {
	return r.wait(ctx)
}

func (r *rowSet) waitForSuccess(ctx context.Context) error {
	if !r.ready {
		status, err := r.wait(ctx)
		if err != nil {
			return err
		}
		if !status.IsSuccess() || !r.ready {
			return fmt.Errorf("Unsuccessful query execution: %+v", status)
		}
	}

	return nil
}

func (r *rowSet) next(ctx context.Context) bool {
	if err := r.waitForSuccess(ctx); err != nil {
		return false
	}

	if r.rowSet == nil || r.offset >= len(r.rowSet.Data) {
		if !r.hasMore {
			return false
		}

		resp, err := r.client.Fetch(ctx, r.handle, false, 1000000)
		if err != nil {
			log.Printf("FetchResults failed: %v\n", err)
			return false
		}

		if r.metadata == nil {
			r.metadata, err = r.client.GetResultsMetadata(ctx, r.handle)
			if err != nil {
				log.Printf("GetResultsMetadata failed: %v\n", err)
			}
		}
		if len(r.columnNames) == 0 {
			r.columnNames = resp.Columns
		}

		r.hasMore = resp.HasMore

		r.rowSet = resp
		r.offset = 0

		// We assume that if we get 0 results back, that's the end
		if len(resp.Data) == 0 {
			return false
		}
	}

	r.nextRow = strings.Split(r.rowSet.Data[r.offset], "\t")
	r.offset++

	return true
}

// Prepares a row for scanning into memory, by reading data from hive if
// the operation is successful, blocking until the operation is
// complete, if necessary.
// Returns true is a row is available to Scan(), and false if the
// results are empty or any other error occurs.
func (r *rowSet) Next() bool {
	return r.next(context.Background())
}

// Same as the Next() but additionally takes the context.
func (r *rowSet) NextWithContext(ctx context.Context) bool {
	return r.next(ctx)
}

// Scan the last row prepared via Next() into the destination(s) provided,
// which must be pointers to value types, as in database.sql. Further,
// only pointers of the following types are supported:
//  - int, int16, int32, int64
//  - string, []byte
//  - float64
//   - bool
func (r *rowSet) Scan(dest ...interface{}) error {
	// TODO: Add type checking and conversion between compatible
	// types where possible, as well as some common error checking,
	// like passing nil. database/sql's method is very convenient,
	// for example: http://golang.org/src/pkg/database/sql/convert.go, like 85
	if r.nextRow == nil {
		return errors.New("No row to scan! Did you call Next() first?")
	}

	if len(dest) != len(r.nextRow) {
		return fmt.Errorf("Can't scan into %d arguments with input of length %d", len(dest), len(r.nextRow))
	}

	for i, val := range r.nextRow {
		d := dest[i]
		switch dt := d.(type) {
		case *string:
			*dt = val
		case *int:
			i, _ := strconv.ParseInt(val, 10, 0)
			*dt = int(i)
		case *int64:
			i, _ := strconv.ParseInt(val, 10, 0)
			*dt = int64(i)
		case *int32:
			i, _ := strconv.ParseInt(val, 10, 0)
			*dt = int32(i)
		case *int16:
			i, _ := strconv.ParseInt(val, 10, 0)
			*dt = int16(i)
		case *float64:
			*dt, _ = strconv.ParseFloat(val, 64)
			/*
			   case *[]byte:
			       *dt = []byte(val.(string))
			   case *bool:
			       *dt = val.(bool)
			*/
		default:
			return fmt.Errorf("Can't scan value of type %T with value %v", dt, val)
		}
	}

	return nil
}

//Convert from a hive column type to a Go type
func (r *rowSet) convertRawValue(raw string, hiveType string) (interface{}, error) {
	if strings.ToLower(raw) == "null" {
		return raw,nil
	}
	switch hiveType {
	case "string":
		return raw, nil
	case "int", "tinyint", "smallint":
		i, err := strconv.ParseInt(raw, 10, 0)
		return int32(i), err
	case "bigint":
		i, err := strconv.ParseInt(raw, 10, 0)
		return int64(i), err
	case "float", "double", "decimal":
		i, err := strconv.ParseFloat(raw, 64)
		return i, err
	case "timestamp":
		i, err := time.Parse("2006-01-02 15:04:05", raw)
		return i, err
	case "boolean":
		return raw == "true", nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalid hive type %v", hiveType))
	}
}

func (r *rowSet) fetchAll(ctx context.Context) []map[string]interface{} {
	response := make([]map[string]interface{}, 0)
	for r.next(ctx) {
		row := make(map[string]interface{})
		for i, val := range r.nextRow {
			conv, err := r.convertRawValue(val, r.metadata.Schema.FieldSchemas[i].Type)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			row[r.metadata.Schema.FieldSchemas[i].Name] = conv
		}
		response = append(response, row)
	}
	return response
}

//Fetch all rows and convert to a []map[string]interface{} with
//appropriate type conversion already carried out
func (r *rowSet) FetchAll() []map[string]interface{} {
	return r.fetchAll(context.Background())
}

// Same as the FetchAll() but additionally takes the context.
func (r *rowSet) FetchAllWithContext(ctx context.Context) []map[string]interface{} {
	return r.fetchAll(ctx)
}

func (r *rowSet) columns(ctx context.Context) []string {
	if r.columnNames == nil {
		if err := r.waitForSuccess(ctx); err != nil {
			return nil
		}
	}

	return r.columnNames
}

// Returns the names of the columns for the given operation,
// blocking if necessary until the information is available.
func (r *rowSet) Columns() []string {
	return r.columns(context.Background())
}

// Same as the Columns() but additionally takes the context.
func (r *rowSet) ColumnsWithContext(ctx context.Context) []string {
	return r.columns(ctx)
}

// MapScan scans a single Row into the dest map[string]interface{}.
func (r *rowSet) MapScan(row map[string]interface{}) error {
	for i, val := range r.nextRow {
		conv, err := r.convertRawValue(val, r.metadata.Schema.FieldSchemas[i].Type)
		if err != nil {
			return err
		}
		row[r.metadata.Schema.FieldSchemas[i].Name] = conv
	}
	return nil
}
