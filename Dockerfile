FROM alpine

# COPY /proto /product-service/src/proto


COPY . .

# FROM alpine as builder

# # Instalar el compilador protoc en la imagen
RUN apk add --no-cache protobuf go 

RUN chmod +x ./compile-proto.sh

CMD [ "./compile-proto.sh" ]

# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
#     go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

# WORKDIR /app
# COPY ./ /app
# COPY proto/product.proto /proto/product.proto

# # Install protoc and zip system library
# RUN mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
#     unzip protoc-3.7.0-linux-x86_64.zip

# ENV PATH="$PATH:$(go env GOPATH)/bin"
# # Copy the grpc proto file and generate the go module


# RUN protoc --proto_path=./proto --go_out=./app/apigateway/proto/ --go-grpc_out=./app/apigateway/proto/ ./proto/product.proto