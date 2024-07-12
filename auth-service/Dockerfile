FROM golang:1.22-alpine AS builder

WORKDIR /auth-service

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/authServer /auth-service/cmd/server.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /build/authServer .

EXPOSE 50056

CMD [ "./authServer" ]