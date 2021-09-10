package services

import (
	"context"
	"simplebank/api/sqlc"
	"time"
)

type UserService interface {
	GetUser(username string) (sqlc.User, error)
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

func newUserResponse(user sqlc.User) *userResponse {
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

