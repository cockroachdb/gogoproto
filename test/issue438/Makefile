regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogo
	go install github.com/cockroachdb/gogoproto/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mgoogle/protobuf/any.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/api.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/duration.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/empty.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/field_mask.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/source_context.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/struct.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/type.proto=github.com/cockroachdb/gogoproto/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/cockroachdb/gogoproto/types,\
	:. --proto_path=../../../../../:../../protobuf/:. issue438.proto
