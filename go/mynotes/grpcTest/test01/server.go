package main

import (
	"net"
	"google.golang.org/grpc"
	"mynotes/grpcTest/api"
	"log"
	"golang.org/x/net/context"
	"mynotes/grpcTest/test01/response"
	"fmt"
	"reflect"
	"io"
)

const (
	PORT = ":10023"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(
		ChainUnaryServer()))
	api.RegisterDataServer(s, &response.Server{Handler:&response.Handler{Token:lis.Addr().String()}})
	s.Serve(lis)

}

type Server struct{
	Handler *Handler
	RouteNotes    []*api.Response
}

type Handler struct {
	Token string
}

func (this *Handler) GetToken(str string) string {
	return this.Token+str+str
}

//string
func (this *Server) Test01(ctx context.Context, req *api.Request)(*api.Response,error){
	test := req.GetTest()
	fmt.Println("Test01:",this.Handler.GetToken(test))
	fmt.Println("Test01:",test)
	return &api.Response{Test : "test"},nil
}

//stream
func (this *Server) Test02(stream api.Data_Test02Server)(error){
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			return nil
		}
		if err != nil {
			fmt.Println("ERR",err)
			return err
		}
		fmt.Println("in ",in,",len:",len(this.RouteNotes))
		for _, note := range this.RouteNotes{
			fmt.Println("note:",note)
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}


func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	switch len(interceptors) {
	case 0:
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			fmt.Println("-------",reflect.TypeOf(req),req)
			return handler(ctx, req)
		}
	case 1:
		return interceptors[0]
	default:
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			buildChain := func(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return current(currentCtx, currentReq, info, next)
				}
			}
			chain := handler
			for i := len(interceptors) - 1; i >= 0; i-- {
				chain = buildChain(interceptors[i], chain)
			}
			return chain(ctx, req)
		}
	}
}