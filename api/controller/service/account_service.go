package service

import (
	"context"
	"simplebank/api/sqlc"
)

type AccountService interface {
	GetAccountByID(id int64) (sqlc.Account, error)
	CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Account, error)
}

type accountServiceImpl struct {
	store *sqlc.Store
}

func NewAccountService(store *sqlc.Store) AccountService {
	res := &accountServiceImpl{}
	res.store = store
	return res
}

func (s *accountServiceImpl) GetAccountByID(id int64) (sqlc.Account, error) {
	account, err := s.store.GetAccount(context.Background(), id)
	if err != nil {
		return sqlc.Account{}, err
	}
	return account, nil
}

func (s *accountServiceImpl) CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Account, error) {
	account, err := s.store.CreateAccount(context.Background(),*params)
	if err != nil {
		return sqlc.Account{}, err
	}
	return account, nil
}
