package auth

import (
	"context"

	pb "getitqec.com/server/user/pkg/api/v1"
	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/model"
	// "github.com/golang/protobuf/ptypes/empty"
	//pb "./proto"
)

// var httpClient = &http.Client{}

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
// func logger(format string, a ...interface{}) {
// 	fmt.Printf("LOG:\t"+format+"\n", a...)
// }

// var (
// 	//port = flag.Int("port", 50051, "the port to serve on")

// 	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
// 	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
// )

// Server class
type Server struct {
	model model.UserModelI

	pb.UnimplementedUserServiceServer
}

// SignIn function
func (s *Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInReponse, error) {
	// google sign in
	user, exist, err := s.model.GoogleSignIn(ctx, req.IdToken, req.AccessToken)
	if err != nil {
		return nil, err
	}
	// TODO: return user
	return &pb.SignInReponse{
		Exist: exist,
		User: &pb.User{
			Id:    user.UserId,
			Name:  user.UserName,
			Email: user.Email,
			Img:   user.Img,
		},
	}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

// GetUser function
func (s *Server) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// TODO: check authorization (no need for hackathon)
	user, err := s.model.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	token, err := commons.VerifyGoogleAccessToken(ctx)
	if err == nil && token.UserId == user.UserId {
		return &pb.User{
			Id:    user.UserId,
			Email: user.Email,
			Name:  user.UserName,
			Img:   user.Img,
		}, nil
	}
	return &pb.User{
		Id:    user.UserId,
		Email: "",
		Name:  user.UserName,
		Img:   user.Img,
	}, nil
}

func (s *Server) SearchUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// TODO: check authorization (no need for hackathon)
	user, err := s.model.SearchUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	token, err := commons.VerifyGoogleAccessToken(ctx)
	if err == nil && token.UserId == user.UserId {
		return &pb.User{
			Id:    user.UserId,
			Email: user.Email,
			Name:  user.UserName,
			Img:   user.Img,
		}, nil
	}
	return &pb.User{
		Id:    user.UserId,
		Email: "",
		Name:  user.UserName,
		Img:   user.Img,
	}, nil
}

// func (s *Server) UpdateUser(ctx context.Context, request *pb.User) (*pb.Acknowledgement, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
// }

// NewServer return new auth server service
func NewServer(model model.UserModelI) *Server {
	server := &Server{}
	server.model = model
	return server
}
