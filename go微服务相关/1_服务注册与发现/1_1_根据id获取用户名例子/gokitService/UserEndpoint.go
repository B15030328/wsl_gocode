package gokitService

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*
2、endpoint层用于实现response,request规范，包裹
*/

type UserRequest struct {
	UserId int `json:"user_id"`
}

type UserResponse struct {
	UserName string `json:"user_name"`
}

func GenUserEndpoint(userservice IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		username := userservice.GetName(r.UserId)
		return UserResponse{
			UserName: username,
		}, nil
	}
}
