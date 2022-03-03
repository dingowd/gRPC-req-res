#Simple gRPC

sending string request to gRPC server and getting response

GOOS=linux
GOARCH=amd64

docker run -it -p 50051:50051 -d grpc