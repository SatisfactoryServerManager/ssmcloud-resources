


find ./proto -type f -name "*.proto" -exec protoc --experimental_allow_proto3_optional --go_out=proto/generated --go_opt=paths=source_relative --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative --proto_path=proto {} \;
