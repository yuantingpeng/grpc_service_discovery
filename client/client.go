package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb_echo "../pb/echo"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func main() {
    // Set up a connection to the server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect:%v", err)
    }
    defer conn.Close()
    c := pb_echo.NewEchoSvrClient(conn)
    
    // contact the server and print out its response
    str := defaultName
    if len(os.Args) > 1 {
      	str = os.Args[1]
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.Echo(ctx, &pb_echo.EchoReq{Str: str})
    if err != nil {
        log.Fatalf("could not greet:%v", err)
    }

    log.Printf("Get: %s", r.Str)
}
