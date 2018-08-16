package main


import (
	"fmt"
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"mynotes/grpcTest/api"
	"google.golang.org/grpc/status"
	"time"
)

const (
	port = ":50051"
)

// server实现服务端
type server struct {
	mu    sync.Mutex
	count map[string]int
}

//实现服务端方法
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Track the number of times the user has been greeted.
	fmt.Println("aaa:",s.count[in.Name])
	s.count[in.Name]++
	if s.count[in.Name] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
		ds, err := st.WithDetails(
			&epb.QuotaFailure{
				Violations: []*epb.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "Limit one greeting per person",
				}},
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	time.Sleep(time.Second*3)
	return &api.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	log.Printf("server starting on port %s...", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHelloServer(s, &server{count:make(map[string]int)})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}