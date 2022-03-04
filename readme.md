#Simple gRPC

Creates 2 docker containers, bridge subnet inside docker and links containers.
Client sending string request to gRPC server and getting response.

Server container from scratch

Client container from scratch

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

server/go build -a -installsuffix cgo -o server .

client/go build -a -installsuffix cgo -o client .

server/docker build . --tag grpc-small

client/docker build . --tag grpc-client

docker-compose up -d

docker attach grpc-client