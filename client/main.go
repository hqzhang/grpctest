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
    client := proto.NewAddServiceClient(conn)
    
    ///
            // Set up a connection to the server.
        defer conn.Close()
        // Contact the server and print out its response.
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        
        r, err := client.Add(ctx, &proto.Request{A: 2,B: 3})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Add 2 and 3: %d", r.GetResult())


         r, err = client.Multiply(ctx, &proto.Request{A: 2,B: 3})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Times 2 and 3: %d", r.GetResult())
}
