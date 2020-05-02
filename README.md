# grpctest
# install dep
```
go get google.golang.org/grpc@v1.28.1
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
#mkdir myprotoc and touch service.go
```
syntax = "proto3";
package proto;

message Request {
   int64 a = 1;
   int64 b = 2;
}
message Response {
   int64 result = 1;
}
service AddService {
   rpc Add(Request) returns (Response);
   rpc Multiply(Request) returns (Response);
}
```
#generate protobuff
```
protoc --proto_path=myproto  --go_out=plugins=grpc:myproto service.proto
```
#mkdir server and touch main.go
```
package main

import (
     "context"
     "google.golang.org/grpc"
     "net"
     "google.golang.org/grpc/reflection"
     proto "../myproto"
 )

type server struct{}

func (s *server) Add(ctx context.Context, request *proto.Request) ( *proto.Response, error) {
    a, b := request.GetA(), request.GetB()
    result := a + b
    return &proto.Response{Result: result}, nil
}

func main() {
    listener, err := net.Listen("tcp", ":4040")
    if err != nil {
        panic(err)
    }
    srv := grpc.NewServer()
    proto.RegisterAddServiceServer(srv, &server{} )
    reflection.Register(srv)
    if e := srv.Serve(listener); e != nil {
         panic(e)
    }
}
```
#mkdir client and touch main.go
```
package main
  
import (
   "google.golang.org/grpc"
   proto "../myproto"
   "context"
   "log"
   "time"

)

func main() {
    conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure() )
    if err != nil {
      panic(err)
    }
    myclient := proto.NewAddServiceClient(conn)
    defer conn.Close()
        // Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := myclient.Add(ctx, &proto.Request{A: 2,B: 3})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Add 2 and 3: %d", r.GetResult())
}
```
