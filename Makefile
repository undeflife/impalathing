
thrift:
	thrift -r -gen go:package_prefix=impalathing/services/ interfaces/ImpalaService.thrift
	rm -rf ./services
	mv gen-go services
