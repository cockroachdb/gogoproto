regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-min-version
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. \
	--proto_path=../../../../../:../../protobuf/:. merge.proto
