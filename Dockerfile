FROM alpine

COPY . .

RUN apk add --no-cache protobuf go 

RUN chmod +x ./compile-proto.sh

CMD [ "./compile-proto.sh" ]

