GOOS=linux CGO_ENABLED=0 go build -o client client.go

Create network to connect client and server:
```
docker network create mynetwork

Check container IP with command:  docker network inspect mynetwork | grep IPAddress

```

Run client and server:

docker run --net=mynetwork grpc-crud-client

docker run -p 50051:50051 --net=mynetwork grpc-crud