package service

import (
	"context"
	"simplebank/api/sqlc"
	"simplebank/api/util"
)

type UserService interface {
	GetUser(username string) (sqlc.User, error)
	CreateUser(params *sqlc.CreateUserParams) (sqlc.User, error)
}

type userServiceImpl struct {
	store sqlc.Store
}

func NewUserService(store sqlc.Store) UserService {
	res := &userServiceImpl{}
	res.store = store
	return res
}

func (s *userServiceImpl) GetUser(username string) (sqlc.User, error) {
	user, err := s.store.GetUser(context.Background(), username)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

func (s *userServiceImpl) CreateUser(params *sqlc.CreateUserParams) (sqlc.User, error) {
	//TODO:ここではまだhashedされてないのでPasswordで扱いたい
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return sqlc.User{},err
	}
	newParams := sqlc.CreateUserParams{
		Username:       params.Username,
		HashedPassword: hashedPassword,
		FullName:       params.FullName,
		Email:          params.Email,
	}
	user, err := s.store.CreateUser(context.Background(), newParams)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}
