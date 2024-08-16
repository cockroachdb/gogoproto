regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
