package main

import (
	"log"
	"time"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "mynotes/testZipkin/proto/add"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	address = "localhost:50051"
)

func main() {
	num1 := flag.Int("num1", 1, "")
	num2 := flag.Int("num2", 2, "")
	flag.Parse()

	collector, err := zipkin.NewHTTPCollector("httpReq://192.168.1.6:9411/api/v1/spans")
	if err != nil {
		log.Fatal(err)
		return
	}

	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "localhost:0", "grpc_client"),
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	opentracing.InitGlobalTracer(tracer)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)

	// Create Root Span for duration of the interaction with svc1
	span := opentracing.StartSpan("Start")

	// Put root span in context so it will be used in our calls to the client.
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	time.Sleep(time.Duration(20) * time.Millisecond)
	// Contact the server and print out its response.
	r, err := c.DoAdd(ctx, &pb.AddRequest{Num1: int32(*num1), Num2: int32(*num2)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("add(%d,%d), Result: %d", *num1, *num2, r.GetResult())

	span.Finish()
	collector.Close()
}
