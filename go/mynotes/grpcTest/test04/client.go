package main


import (
	"log"
	"os"
	"time"
	"golang.org/x/net/context"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"mynotes/grpcTest/api"
	"google.golang.org/grpc/status"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	client := api.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &api.HelloRequest{Name: "world"})
	if err != nil {
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *epb.QuotaFailure:
				log.Printf("Quota failure: %s", info)
			default:
				log.Printf("Unexpected type: %s", info)
			}
		}
		os.Exit(1)
	}
	log.Printf("Greeting: %s", r.Message)
}
