package services

import (
	"context"
	"simplebank/api/sqlc"
	"simplebank/api/util"
	"time"
)

type UserService interface {
	GetUser(username string) (sqlc.User, error)
	CreateUser(params *sqlc.CreateUserParams) (*userResponse, error)
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

func (s *userServiceImpl) CreateUser(params *sqlc.CreateUserParams) (*userResponse, error) {
	//TODO:ここではまだhashedされてないのでPasswordで扱いたい
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return &userResponse{}, err
	}
	newParams := sqlc.CreateUserParams{
		Username:       params.Username,
		HashedPassword: hashedPassword,
		FullName:       params.FullName,
		Email:          params.Email,
	}
	user, err := s.store.CreateUser(context.Background(), newParams)
	if err != nil {
		return &userResponse{}, err
	}

	res := NewUserResponse(user)
	return res, nil
}

func NewUserResponse(user sqlc.User) *userResponse {
	return &userResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

