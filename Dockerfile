# build the server binary
FROM golang:1.15
WORKDIR /go/src/grpc-crud
COPY . .
COPY ./test test

EXPOSE 50051
ENTRYPOINT ["./test"]
