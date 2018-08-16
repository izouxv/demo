package main

import (
	"log"
	"net"
	"time"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	cache "mynotes/testZipkin/cache/client"
	pb "mynotes/testZipkin/proto/add"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type AddServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *AddServer) DoAdd(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	log.Printf("input %d %d", in.GetNum1(), in.GetNum2())
	time.Sleep(time.Duration(10) * time.Millisecond)

	tracer := opentracing.GlobalTracer()
	val := cache.GetCache(ctx, tracer, in.GetNum1())
	log.Printf("cache value %d", val)

	return &pb.AddReply{Result: val + in.GetNum2()}, nil
}

func main() {
	collector, err := zipkin.NewHTTPCollector("httpReq://192.168.1.6:9411/api/v1/spans")
	if err != nil {
		log.Fatal(err)
		return
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "localhost:0", "grpc_server"),
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	opentracing.InitGlobalTracer(tracer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	pb.RegisterAddServer(s, &AddServer{})
	s.Serve(lis)
}
