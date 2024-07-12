# 🗄️ Microservices with gRPC

## 🧑🏻‍💻 _Author: Aksel Rivera_

This is an implementation of microservice architecture, gRPC Protocols, API Gateway.
Powered with Golang and Nest JS.

## 📜 Description

This project demonstrates a microservices architecture using an API Gateway built with Golang and Fiber, along with three microservices:

- API Gateway (Go - Fiber)
  - Handles every request from the client and deliver the data to the corresponding.
- Product Service (NestJS)
  - Handles product-related functionality.
  - Communicates with other services via gRPC.
- Order Service (Go)
  - Manages order creation and retrieval.
  - Communicates with other services via gRPC.
- Auth Service (Go)
  - Provides authentication functionality.
  - Integrates with the client-side authentication

## ✨ Installation and Setup

### Prerequisites

- Docker
- Node.js
- Go
- Protobuf
- ✨Magic ✨

## 👣 Steps

### Clone the Repository

```
git clone https://github.com/your-username/microservices-project.git
cd microservices-project
```

## 🤖 Compile Protobuf Files

Compile the `.proto` files to generate gRPC code:

> Please install [proto bins files](https://grpc.io/docs/protoc-installation) because we are going to use them to compile protocols

```
 For Windows:
./compile-proto.bat
```

```
For linux:
./compile-proto.sh
```

## 🐋 Deploy

To deploy this proyect you have to execute this command at the project root:

```
docker compose up --build
```

Once docker compose finish you could visit this [Postman Workspace](https://www.postman.com/orbital-module-geoscientist-17997070/workspace/grpc-aksel-rivera/overview) to interact with the endpoint.
Or you can access to the API Gateway at http://localhost:8080.
Explore the endpoints for products, orders, and authentication.

## 🖥 ️Technologies

The current technology stack for this project:

- [Golang](https://go.dev/)
- [Typescript](https://www.typescriptlang.org/)
- [Fiber](https://docs.gofiber.io/)
- [Nest JS](https://docs.nestjs.com/microservices/grpc)
- [gRPC](https://grpc.io/docs/protoc-installation/)
- [Github](https://github.com/AkselRivera)
- [Postman](https://www.postman.com/orbital-module-geoscientist-17997070/workspace/grpc-aksel-rivera/overview)

## 🚀 Contributing

Feel free to contribute by opening issues or pull requests!
Good luck with your microservices project!

**Happy hacking!**
