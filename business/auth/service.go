package auth

import (
	"errors"
	"go-hexagonal-auth/api/v1/auth/request"
	"go-hexagonal-auth/business"
	"go-hexagonal-auth/business/user"
	"go-hexagonal-auth/config"
	utils "go-hexagonal-auth/util"
	"go-hexagonal-auth/util/validator"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService user.Service
	userRepo   user.Repository
	cfg        config.Config
}

//NewService Construct user service object
func NewService(userService user.Service, userRepo user.Repository, cfg config.Config) Service {
	user := &service{
		userService,
		userRepo,
		cfg,
	}
	return user
}

//Login by given user Username and Password, return error if not exist
func (s *service) Login(username string, isAdmin bool) (*user.User, error) {
	var result user.User

		userData, err := s.userService.FindUserByUsernameAndPassword(username, "")
		if err != nil {
			return nil, err
		}
		if userData == nil {
			return nil, errors.New("data not found")
		}
		result = user.User{
			Name:       userData.Name,
			Username:   userData.Username,
			Password:   userData.Password,
		}




	return &result, nil
}

func (s *service) RegisterUser(request request.RegisterUserRequest) (*request.RegisterUserRequest, error)  {
	err := validator.GetValidator().Struct(request)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}

	pass, err :=utils.HashPassword(request.Password)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}

	UserReq := user.User{
		Name:       request.Name,
		Username:   request.Username,
		Password:   pass,
		Address: request.Address,
	}
	err = s.userRepo.InsertUser(UserReq)
	if err != nil {
		return nil,err
	}

	return &request,nil
}