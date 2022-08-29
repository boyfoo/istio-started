protoc --proto_path=gsrc/protos --go_out=./../ prod_model.proto

protoc --proto_path=gsrc/protos --go_out=plugins=grpc:./../ prod_service.proto

::protoc --proto_path=gsrc/protos --go-grpc_out=./../ prod_service.proto