regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/cockroachdb/gogoproto/protobuf/ \
	--gogo_out=. events.proto 
