FROM golang:1.22-alpine AS builder

WORKDIR /order-service

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/orderServer /order-service/cmd/server.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /build/orderServer .

EXPOSE 50055

CMD [ "./orderServer" ]