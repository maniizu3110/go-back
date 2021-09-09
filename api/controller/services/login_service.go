package services

import (
	"context"
	"simplebank/api/sqlc"
	"simplebank/api/util"
	"simplebank/token"
)

type LoginService interface {
	LoginUser(params *LoginUserRequest) (*loginUserResponse, error)
}

type loginServiceImpl struct {
	store      sqlc.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewLoginService(store sqlc.Store, tokenMaker token.Maker, config util.Config) LoginService {
	res := &loginServiceImpl{}
	res.store = store
	res.tokenMaker = tokenMaker
	res.config = config
	return res
}

func (s *loginServiceImpl) LoginUser(params *LoginUserRequest) (*loginUserResponse, error) {
	user, err := s.store.GetUser(context.Background(),params.Username)
	if err != nil {
		return nil, err
	}
	err = util.CheckPassword(params.Password, user.HashedPassword)
	if err != nil {
		return nil, err
	}
	accessToken, err := s.tokenMaker.CreateToken(user.Username,s.config.AccessTokenDuration)
	if err != nil {
		return nil,err
	}
	createdUser := newUserResponse(user)
	res := &loginUserResponse{
		AccessToken:accessToken,
		User:*createdUser,
	}
	return res,nil
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}
