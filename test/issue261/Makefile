regenerate:
	go install github.com/cockroachdb/gogoproto/protoc-min-version
	go install github.com/cockroachdb/gogoproto/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --gogoslick_out=\
	Mgoogle/protobuf/duration.proto=github.com/cockroachdb/gogoproto/types,\
	:. \
	--proto_path=../../../../../:../../protobuf/:. issue261.proto
