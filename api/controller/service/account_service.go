package service

import (
	"context"
	"simplebank/api/sqlc"
)

type AccountService interface {
	GetAccount(id int64) (sqlc.Account, error)
	CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Account, error)
	GetListAccount(params *sqlc.ListAccountParams) ([]sqlc.Account, error)
	UpdateAccount(params *sqlc.UpdateAccountParams) (sqlc.Account, error)
	DeleteAccount(id int64) error 
}

type accountServiceImpl struct {
	store sqlc.Store
}

func NewAccountService(store sqlc.Store) AccountService {
	res := &accountServiceImpl{}
	res.store = store
	return res
}

func (s *accountServiceImpl) GetAccount(id int64) (sqlc.Account, error) {
	account, err := s.store.GetAccount(context.Background(), id)
	if err != nil {
		return sqlc.Account{}, err
	}
	return account, nil
}

func (s *accountServiceImpl) CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Account, error) {
	account, err := s.store.CreateAccount(context.Background(), *params)
	if err != nil {
		return sqlc.Account{}, err
	}
	return account, nil
}

func (s *accountServiceImpl) GetListAccount(params *sqlc.ListAccountParams) ([]sqlc.Account, error) {
	if params.Limit == 0 {
		params.Limit = 1000
	}
	accouts, err := s.store.ListAccount(context.Background(), *params)
	if err != nil {
		return nil, err
	}
	return accouts, nil
}

func (s *accountServiceImpl) UpdateAccount(params *sqlc.UpdateAccountParams) (sqlc.Account, error) {
	var err error
	// accout, err := s.store.UpdateAccount(context.Background(), *params)
	if err != nil {
		return sqlc.Account{}, err
	}
	return sqlc.Account{}, nil
}

func (s *accountServiceImpl) DeleteAccount(id int64) error {
	err := s.store.DeleteAccount(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
