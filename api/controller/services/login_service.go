package services

import (
	"simplebank/api/sqlc"
	"simplebank/api/util"
)

type LoginService interface {
	LoginUser(params *LoginUserRequest)(*LoginUserResponse,error) 
}

type loginServiceImpl struct {
	store sqlc.Store
}

func NewLoginService(store sqlc.Store) LoginService {
	res := &userServiceImpl{}
	res.store = store
	return res
}


func (s *userServiceImpl) LoginUser(params *LoginUserRequest)(*LoginUserResponse,error) {
	user,err := s.GetUser(params.Username)
	if err != nil {
		return nil,err
	}
	err = util.CheckPassword(params.Password,user.HashedPassword)
	if err != nil {
		return nil,err
	}
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,alphanium"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}
