package filter

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	switch len(interceptors) {
	case 0:
		// do not want to return nil interceptor since this function was never defined to do so/for backwards compatibility
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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

func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	switch len(interceptors) {
	case 0:
		// do not want to return nil interceptor since this function was never defined to do so/for backwards compatibility
		return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, stream)
		}
	case 1:
		return interceptors[0]
	default:
		return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			buildChain := func(current grpc.StreamServerInterceptor, next grpc.StreamHandler) grpc.StreamHandler {
				return func(currentSrv interface{}, currentStream grpc.ServerStream) error {
					return current(currentSrv, currentStream, info, next)
				}
			}
			chain := handler
			for i := len(interceptors) - 1; i >= 0; i-- {
				chain = buildChain(interceptors[i], chain)
			}
			return chain(srv, stream)
		}
	}
}
