package service

import (
	"context"
	"simplebank/api/sqlc"
)

type AccountService interface {
	GetAccount(id int64) (sqlc.Accounts, error)
	CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Accounts, error)
	GetListAccount(params *sqlc.ListAccountsParams) ([]sqlc.Accounts, error)
	UpdateAccount(params *sqlc.UpdateAccountParams) (sqlc.Accounts, error)
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

func (s *accountServiceImpl) GetAccount(id int64) (sqlc.Accounts, error) {
	account, err := s.store.GetAccount(context.Background(), id)
	if err != nil {
		return sqlc.Accounts{}, err
	}
	return account, nil
}

func (s *accountServiceImpl) CreateAccount(params *sqlc.CreateAccountParams) (sqlc.Accounts, error) {
	account, err := s.store.CreateAccount(context.Background(), *params)
	if err != nil {
		return sqlc.Accounts{}, err
	}
	return account, nil
}

func (s *accountServiceImpl) GetListAccount(params *sqlc.ListAccountsParams) ([]sqlc.Accounts, error) {
	if params.Limit == 0 {
		params.Limit = 1000
	}
	accouts, err := s.store.ListAccounts(context.Background(), *params)
	if err != nil {
		return nil, err
	}
	return accouts, nil
}

func (s *accountServiceImpl) UpdateAccount(params *sqlc.UpdateAccountParams) (sqlc.Accounts, error) {
	var err error
	// accout, err := s.store.UpdateAccount(context.Background(), *params)
	if err != nil {
		return sqlc.Accounts{}, err
	}
	return sqlc.Accounts{}, nil
}

func (s *accountServiceImpl) DeleteAccount(id int64) error {
	err := s.store.DeleteAccount(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
