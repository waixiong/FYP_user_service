package middleware

import (
	"context"
	"fmt"

	authproto "getitqec.com/server/user/pkg/api/clients/auth/v1"
	"getitqec.com/server/user/pkg/commons"
	tkn "getitqec.com/server/user/pkg/commons/token"
	"getitqec.com/server/user/pkg/logger"
	"getitqec.com/server/user/pkg/protocol/grpcClient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

// func logger(format string, a ...interface{}) {
// 	fmt.Printf("LOG:\t"+format+"\n", a...)
// }

func AuthFunction(ctx context.Context) (context.Context, error) {
	// authentication (token verification)

	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("\tAuth Func Get metadata...\n")
	if !ok {
		return nil, commons.ErrMissingMetadata
	}
	fmt.Printf("\t%v\n", md.Len)
	tokens := md.Get("Token")
	if len(tokens) != 1 {
		return ctx, commons.ErrInvalidToken
	}
	token := tokens[0]
	// if services or admin
	getToken, err := tkn.RawTokenFromAccessToken(token)
	if err != nil {
		return ctx, commons.ErrInvalidToken
	}
	logger.Log.Debug(getToken.User)
	if commons.IsGetItService(getToken.User) {
		logger.Log.Debug("sudo token")
		md.Set("User", getToken.User)
		return ctx, nil
	}
	logger.Log.Debug("not admin or service")
	// if services or admin

	authClient, conn, err := grpcClient.GetAuthClient()
	if err != nil {
		return ctx, err
	}
	v, err := authClient.VerifyToken(ctx, &authproto.Token{AccessToken: token})
	conn.Close()
	if err != nil {
		return ctx, err
	}
	if v.Result == authproto.VerifyState_ILLEGAL {
		return ctx, commons.ErrInvalidToken
	}
	if v.Result == authproto.VerifyState_REFRESH {
		return ctx, commons.ErrExpiredToken
	}
	md.Set("User", v.UserId)
	return ctx, nil
}

func UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return UnaryServerInterceptor(AuthFunction)
}

func StreamAuthInterceptor() grpc.StreamServerInterceptor {
	return StreamServerInterceptor(AuthFunction)
}

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	// md, ok := metadata.FromIncomingContext(ctx)
	// fmt.Printf("\tGet metadata...\n")
	// if !ok {
	// 	return nil, commons.ErrMissingMetadata
	// }
	// fmt.Printf("\t%v\n", md)
	// if u, e := "", errors.New("Not implement yet"); e != nil {
	// 	return nil, e
	// } else if len(u) == 0 {
	// 	return nil, ErrInvalidToken
	// }
	fmt.Println(info.FullMethod)
	m, err := handler(ctx, req)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("RPC failed with error %v", err))
	}
	return m, err
}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func UnaryServerInterceptor(authFunc grpc_auth.AuthFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var newCtx context.Context
		// var err error

		// check method name
		fmt.Println(info.FullMethod)
		// if commons.IsAuthException(info.FullMethod) {
		// 	fmt.Println("Pass Auth")
		// 	return handler(ctx, req)
		// }

		// if overrideSrv, ok := info.Server.(grpc_auth.ServiceAuthFuncOverride); ok {
		// 	newCtx, err = overrideSrv.AuthFuncOverride(ctx, info.FullMethod)
		// } else {
		// 	newCtx, err = authFunc(ctx)
		// }
		// if err != nil {
		// 	return nil, err
		// }
		return handler(newCtx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptors that performs per-request auth.
func StreamServerInterceptor(authFunc grpc_auth.AuthFunc) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var newCtx context.Context
		var err error

		// check method name
		if info.FullMethod == "/userproto.UserService/signIn" {
			wrapped := grpc_middleware.WrapServerStream(stream)
			wrapped.WrappedContext = stream.Context()
			return handler(srv, wrapped)
		}

		if overrideSrv, ok := srv.(grpc_auth.ServiceAuthFuncOverride); ok {
			newCtx, err = overrideSrv.AuthFuncOverride(stream.Context(), info.FullMethod)
		} else {
			newCtx, err = authFunc(stream.Context())
		}
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}
