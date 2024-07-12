FROM golang:1.22-alpine AS builder

WORKDIR /apigateway

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/server /apigateway/cmd/server.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /build/server .

EXPOSE 8080
EXPOSE 50051

CMD [ "./server" ]