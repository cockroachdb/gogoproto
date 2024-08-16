regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogofast
	protoc-min-version --version="3.0.0" --gogofast_out=Mgoogle/protobuf/descriptor.proto=github.com/cockroachdb/gogoproto/protoc-gen-gogo/descriptor:. --proto_path=../../../../../:../../protobuf/:. *.proto
