protoc \
	--experimental_allow_proto3_optional \
	--go_out=proto/ \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/ \
	--go-grpc_opt=paths=source_relative \
	--proto_path=proto/definitions \
	proto/definitions/*.proto