gen: 
	protoc --proto_path=./grpc/protos --go_out=./grpc/protos --go-grpc_out=./grpc/protos ./grpc/protos/admin.proto
	protoc --proto_path=./grpc/protos --go_out=./grpc/protos --go-grpc_out=./grpc/protos ./grpc/protos/customer.proto
