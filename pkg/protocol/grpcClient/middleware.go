package grpcClient

import (
	"context"
	"fmt"

	tkn "getitqec.com/server/user/pkg/commons/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	SERVICE = "user_service"
)

func UnaryAuthInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// short circuit for simplicity, and avoiding allocations.
		fmt.Println("\tset metadata")
		md := metadata.MD{}
		t, e := tkn.GenerateOneTimeToken(SERVICE)
		if e != nil {
			fmt.Println("\terror set token")
			return e
		}
		md.Set("Token", t)
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamAuthInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		md := metadata.MD{}
		t, e := tkn.GenerateOneTimeToken(SERVICE)
		if e != nil {
			return nil, e
		}
		md.Set("Token", t)
		ctx = metadata.NewOutgoingContext(ctx, md)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
