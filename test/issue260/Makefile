regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/cockroachdb/gogoproto/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
