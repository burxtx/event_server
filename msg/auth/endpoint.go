package auth

import (
	"context"

	"github.com/burxtx/car/users"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Register endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Register: makeRegisterEndpoint(s),
	}
}

type RegisterRequest struct {
	username string
	password string
}

type RegisterResponse struct {
	User *users.User
	Err  error
}

func makeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		u, err := s.Register(ctx, req.username, req.password)
		if err != nil {
			return nil, err
		}
		return RegisterResponse{User: u, Err: err}, nil
	}
}
