#Simple gRPC
Creates 2 docker containers, bridge subnet inside docker and links containers.
Client sending string request to gRPC server and getting response.

Server container from scratch

Client container from ubuntu:20.04

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

server/go build -a -installsuffix cgo -o server .

client/go build -a -installsuffix cgo -o client .

server/docker build . --tag grpc-small

docker run -it grpc-small

client/docker build . --tag grpc-client

docker run -it grpc-client

docker network create --subnet 10.1.0.0/16 --gateway=10.1.0.1 --ip-range 10.1.4.0/24 --driver=bridge --label=dockerhost grpcbridge

docker network connect --alias host grpcbridge (grpc-small)

docker network connect grpcbridge (grpc-client)

Run client inside grpc-client container