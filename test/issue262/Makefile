regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-min-version
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/cockroachdb/gogoproto/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/cockroachdb/gogoproto/types:. timefail.proto
