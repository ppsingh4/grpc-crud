# build the server binary
FROM golang:1.15
WORKDIR /go/src/grpc-crud/client
COPY . .
COPY ./client client

EXPOSE 50051
ENTRYPOINT ["./client"]
