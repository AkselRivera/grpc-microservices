echo "Compilando Archivos .proto - gRPC"
@echo off

@REM  Product - Microservice
protoc --ts_proto_out=./product-service/src/proto --ts_proto_opt=outputServices=grpc-js,env=node --proto_path=./proto ./proto/*.proto

@REM  Order - Microservice
protoc --proto_path=./proto --go_out=./order-service/ --go-grpc_out=./order-service/ ./proto/*.proto


@REM  Auth - Microservice
protoc --proto_path=./proto --go_out=./auth-service/ --go-grpc_out=./auth-service/ ./proto/auth.proto

@REM GO - APIGateway 
protoc --proto_path=./proto --go_out=./apigateway/ --go-grpc_out=./apigateway/ ./proto/*.proto

echo "Archivos compilados con exito"
