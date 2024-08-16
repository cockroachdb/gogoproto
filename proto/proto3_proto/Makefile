regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/cockroachdb/gogoproto/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/cockroachdb/gogoproto/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
