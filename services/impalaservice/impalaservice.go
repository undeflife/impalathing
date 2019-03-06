// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package impalaservice

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/koblas/impalathing/services/beeswax"
	"github.com/koblas/impalathing/services/cli_service"
	"github.com/koblas/impalathing/services/status"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__

type ImpalaService interface {
	beeswax.BeeswaxService

	// Parameters:
	//  - QueryID
	Cancel(query_id *beeswax.QueryHandle) (r *status.TStatus, err error)
	// Parameters:
	//  - Handle
	CloseInsert(handle *beeswax.QueryHandle) (r *TInsertResult_, err error)
	PingImpalaService() (err error)
}

type ImpalaServiceClient struct {
	*beeswax.BeeswaxServiceClient
}

func NewImpalaServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ImpalaServiceClient {
	return &ImpalaServiceClient{BeeswaxServiceClient: beeswax.NewBeeswaxServiceClientFactory(t, f)}
}

func NewImpalaServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ImpalaServiceClient {
	return &ImpalaServiceClient{BeeswaxServiceClient: beeswax.NewBeeswaxServiceClientProtocol(t, iprot, oprot)}
}

// Parameters:
//  - QueryID
func (p *ImpalaServiceClient) Cancel(query_id *beeswax.QueryHandle) (r *status.TStatus, err error) {
	if err = p.sendCancel(query_id); err != nil {
		return
	}
	return p.recvCancel()
}

func (p *ImpalaServiceClient) sendCancel(query_id *beeswax.QueryHandle) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Cancel", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ImpalaServiceCancelArgs{
		QueryID: query_id,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvCancel() (value *status.TStatus, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Cancel" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Cancel failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Cancel failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Cancel failed: invalid message type")
		return
	}
	result := ImpalaServiceCancelResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Error != nil {
		err = result.Error
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Handle
func (p *ImpalaServiceClient) CloseInsert(handle *beeswax.QueryHandle) (r *TInsertResult_, err error) {
	if err = p.sendCloseInsert(handle); err != nil {
		return
	}
	return p.recvCloseInsert()
}

func (p *ImpalaServiceClient) sendCloseInsert(handle *beeswax.QueryHandle) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("CloseInsert", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ImpalaServiceCloseInsertArgs{
		Handle: handle,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvCloseInsert() (value *TInsertResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "CloseInsert" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "CloseInsert failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "CloseInsert failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "CloseInsert failed: invalid message type")
		return
	}
	result := ImpalaServiceCloseInsertResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Error != nil {
		err = result.Error
		return
	} else if result.Error2 != nil {
		err = result.Error2
		return
	}
	value = result.GetSuccess()
	return
}

func (p *ImpalaServiceClient) PingImpalaService() (err error) {
	if err = p.sendPingImpalaService(); err != nil {
		return
	}
	return p.recvPingImpalaService()
}

func (p *ImpalaServiceClient) sendPingImpalaService() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("PingImpalaService", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ImpalaServicePingImpalaServiceArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvPingImpalaService() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "PingImpalaService" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "PingImpalaService failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "PingImpalaService failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "PingImpalaService failed: invalid message type")
		return
	}
	result := ImpalaServicePingImpalaServiceResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type ImpalaServiceProcessor struct {
	*beeswax.BeeswaxServiceProcessor
}

func NewImpalaServiceProcessor(handler ImpalaService) *ImpalaServiceProcessor {
	self8 := &ImpalaServiceProcessor{beeswax.NewBeeswaxServiceProcessor(handler)}
	self8.AddToProcessorMap("Cancel", &impalaServiceProcessorCancel{handler: handler})
	self8.AddToProcessorMap("CloseInsert", &impalaServiceProcessorCloseInsert{handler: handler})
	self8.AddToProcessorMap("PingImpalaService", &impalaServiceProcessorPingImpalaService{handler: handler})
	return self8
}

type impalaServiceProcessorCancel struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorCancel) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ImpalaServiceCancelArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Cancel", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ImpalaServiceCancelResult{}
	var retval *status.TStatus
	var err2 error
	if retval, err2 = p.handler.Cancel(args.QueryID); err2 != nil {
		switch v := err2.(type) {
		case *beeswax.BeeswaxException:
			result.Error = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Cancel: "+err2.Error())
			oprot.WriteMessageBegin("Cancel", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Cancel", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type impalaServiceProcessorCloseInsert struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorCloseInsert) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ImpalaServiceCloseInsertArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CloseInsert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ImpalaServiceCloseInsertResult{}
	var retval *TInsertResult_
	var err2 error
	if retval, err2 = p.handler.CloseInsert(args.Handle); err2 != nil {
		switch v := err2.(type) {
		case *beeswax.QueryNotFoundException:
			result.Error = v
		case *beeswax.BeeswaxException:
			result.Error2 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CloseInsert: "+err2.Error())
			oprot.WriteMessageBegin("CloseInsert", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CloseInsert", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type impalaServiceProcessorPingImpalaService struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorPingImpalaService) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ImpalaServicePingImpalaServiceArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("PingImpalaService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ImpalaServicePingImpalaServiceResult{}
	var err2 error
	if err2 = p.handler.PingImpalaService(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing PingImpalaService: "+err2.Error())
		oprot.WriteMessageBegin("PingImpalaService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("PingImpalaService", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - QueryID
type ImpalaServiceCancelArgs struct {
	QueryID *beeswax.QueryHandle `thrift:"query_id,1" json:"query_id"`
}

func NewImpalaServiceCancelArgs() *ImpalaServiceCancelArgs {
	return &ImpalaServiceCancelArgs{}
}

var ImpalaServiceCancelArgs_QueryID_DEFAULT *beeswax.QueryHandle

func (p *ImpalaServiceCancelArgs) GetQueryID() *beeswax.QueryHandle {
	if !p.IsSetQueryID() {
		return ImpalaServiceCancelArgs_QueryID_DEFAULT
	}
	return p.QueryID
}
func (p *ImpalaServiceCancelArgs) IsSetQueryID() bool {
	return p.QueryID != nil
}

func (p *ImpalaServiceCancelArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServiceCancelArgs) readField1(iprot thrift.TProtocol) error {
	p.QueryID = &beeswax.QueryHandle{}
	if err := p.QueryID.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.QueryID), err)
	}
	return nil
}

func (p *ImpalaServiceCancelArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Cancel_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServiceCancelArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("query_id", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:query_id: ", p), err)
	}
	if err := p.QueryID.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.QueryID), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:query_id: ", p), err)
	}
	return err
}

func (p *ImpalaServiceCancelArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServiceCancelArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Error
type ImpalaServiceCancelResult struct {
	Success *status.TStatus           `thrift:"success,0" json:"success,omitempty"`
	Error   *beeswax.BeeswaxException `thrift:"error,1" json:"error,omitempty"`
}

func NewImpalaServiceCancelResult() *ImpalaServiceCancelResult {
	return &ImpalaServiceCancelResult{}
}

var ImpalaServiceCancelResult_Success_DEFAULT *status.TStatus

func (p *ImpalaServiceCancelResult) GetSuccess() *status.TStatus {
	if !p.IsSetSuccess() {
		return ImpalaServiceCancelResult_Success_DEFAULT
	}
	return p.Success
}

var ImpalaServiceCancelResult_Error_DEFAULT *beeswax.BeeswaxException

func (p *ImpalaServiceCancelResult) GetError() *beeswax.BeeswaxException {
	if !p.IsSetError() {
		return ImpalaServiceCancelResult_Error_DEFAULT
	}
	return p.Error
}
func (p *ImpalaServiceCancelResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ImpalaServiceCancelResult) IsSetError() bool {
	return p.Error != nil
}

func (p *ImpalaServiceCancelResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServiceCancelResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &status.TStatus{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *ImpalaServiceCancelResult) readField1(iprot thrift.TProtocol) error {
	p.Error = &beeswax.BeeswaxException{
		SQLState: "     ",
	}
	if err := p.Error.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error), err)
	}
	return nil
}

func (p *ImpalaServiceCancelResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Cancel_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServiceCancelResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ImpalaServiceCancelResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetError() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err)
		}
		if err := p.Error.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err)
		}
	}
	return err
}

func (p *ImpalaServiceCancelResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServiceCancelResult(%+v)", *p)
}

// Attributes:
//  - Handle
type ImpalaServiceCloseInsertArgs struct {
	Handle *beeswax.QueryHandle `thrift:"handle,1" json:"handle"`
}

func NewImpalaServiceCloseInsertArgs() *ImpalaServiceCloseInsertArgs {
	return &ImpalaServiceCloseInsertArgs{}
}

var ImpalaServiceCloseInsertArgs_Handle_DEFAULT *beeswax.QueryHandle

func (p *ImpalaServiceCloseInsertArgs) GetHandle() *beeswax.QueryHandle {
	if !p.IsSetHandle() {
		return ImpalaServiceCloseInsertArgs_Handle_DEFAULT
	}
	return p.Handle
}
func (p *ImpalaServiceCloseInsertArgs) IsSetHandle() bool {
	return p.Handle != nil
}

func (p *ImpalaServiceCloseInsertArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertArgs) readField1(iprot thrift.TProtocol) error {
	p.Handle = &beeswax.QueryHandle{}
	if err := p.Handle.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Handle), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CloseInsert_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("handle", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:handle: ", p), err)
	}
	if err := p.Handle.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Handle), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:handle: ", p), err)
	}
	return err
}

func (p *ImpalaServiceCloseInsertArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServiceCloseInsertArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Error
//  - Error2
type ImpalaServiceCloseInsertResult struct {
	Success *TInsertResult_                 `thrift:"success,0" json:"success,omitempty"`
	Error   *beeswax.QueryNotFoundException `thrift:"error,1" json:"error,omitempty"`
	Error2  *beeswax.BeeswaxException       `thrift:"error2,2" json:"error2,omitempty"`
}

func NewImpalaServiceCloseInsertResult() *ImpalaServiceCloseInsertResult {
	return &ImpalaServiceCloseInsertResult{}
}

var ImpalaServiceCloseInsertResult_Success_DEFAULT *TInsertResult_

func (p *ImpalaServiceCloseInsertResult) GetSuccess() *TInsertResult_ {
	if !p.IsSetSuccess() {
		return ImpalaServiceCloseInsertResult_Success_DEFAULT
	}
	return p.Success
}

var ImpalaServiceCloseInsertResult_Error_DEFAULT *beeswax.QueryNotFoundException

func (p *ImpalaServiceCloseInsertResult) GetError() *beeswax.QueryNotFoundException {
	if !p.IsSetError() {
		return ImpalaServiceCloseInsertResult_Error_DEFAULT
	}
	return p.Error
}

var ImpalaServiceCloseInsertResult_Error2_DEFAULT *beeswax.BeeswaxException

func (p *ImpalaServiceCloseInsertResult) GetError2() *beeswax.BeeswaxException {
	if !p.IsSetError2() {
		return ImpalaServiceCloseInsertResult_Error2_DEFAULT
	}
	return p.Error2
}
func (p *ImpalaServiceCloseInsertResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ImpalaServiceCloseInsertResult) IsSetError() bool {
	return p.Error != nil
}

func (p *ImpalaServiceCloseInsertResult) IsSetError2() bool {
	return p.Error2 != nil
}

func (p *ImpalaServiceCloseInsertResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TInsertResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertResult) readField1(iprot thrift.TProtocol) error {
	p.Error = &beeswax.QueryNotFoundException{}
	if err := p.Error.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertResult) readField2(iprot thrift.TProtocol) error {
	p.Error2 = &beeswax.BeeswaxException{
		SQLState: "     ",
	}
	if err := p.Error2.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error2), err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CloseInsert_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServiceCloseInsertResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ImpalaServiceCloseInsertResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetError() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err)
		}
		if err := p.Error.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err)
		}
	}
	return err
}

func (p *ImpalaServiceCloseInsertResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetError2() {
		if err := oprot.WriteFieldBegin("error2", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:error2: ", p), err)
		}
		if err := p.Error2.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error2), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:error2: ", p), err)
		}
	}
	return err
}

func (p *ImpalaServiceCloseInsertResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServiceCloseInsertResult(%+v)", *p)
}

type ImpalaServicePingImpalaServiceArgs struct {
}

func NewImpalaServicePingImpalaServiceArgs() *ImpalaServicePingImpalaServiceArgs {
	return &ImpalaServicePingImpalaServiceArgs{}
}

func (p *ImpalaServicePingImpalaServiceArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServicePingImpalaServiceArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PingImpalaService_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServicePingImpalaServiceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServicePingImpalaServiceArgs(%+v)", *p)
}

type ImpalaServicePingImpalaServiceResult struct {
}

func NewImpalaServicePingImpalaServiceResult() *ImpalaServicePingImpalaServiceResult {
	return &ImpalaServicePingImpalaServiceResult{}
}

func (p *ImpalaServicePingImpalaServiceResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ImpalaServicePingImpalaServiceResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PingImpalaService_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ImpalaServicePingImpalaServiceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ImpalaServicePingImpalaServiceResult(%+v)", *p)
}
