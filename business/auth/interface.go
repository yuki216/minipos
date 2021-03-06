package auth

import (
	"go-hexagonal-auth/api/v1/auth/request"
	"go-hexagonal-auth/business/user"
)

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	Login(username string, isAdmin bool) (*user.User, error)
	RegisterUser(request request.RegisterUserRequest) (*request.RegisterUserRequest, error)
}

