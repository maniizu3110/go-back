package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"simplebank/api/sqlc"

	"github.com/labstack/echo/v4"
)

type TransferService interface {
	GetTransfer(id int64) (sqlc.Transfer, error)
	CreateTransfer(params *sqlc.TransferTxParams) (sqlc.TransferTxResult, error)
	GetListTransfer(params *sqlc.ListTransferParams) ([]sqlc.Transfer, error)
}

type transferServiceImpl struct {
	store sqlc.Store
}

func NewTransferService(store sqlc.Store) TransferService {
	res := &transferServiceImpl{}
	res.store = store
	return res
}

func (s *transferServiceImpl) GetTransfer(id int64) (sqlc.Transfer, error) {
	transfer, err := s.store.GetTransfer(context.Background(), id)
	if err != nil {
		return sqlc.Transfer{}, err
	}
	return transfer, nil
}

func (s *transferServiceImpl) CreateTransfer(params *sqlc.TransferTxParams) (sqlc.TransferTxResult, error) {
	var err error
	if !s.validAccount(params.FromAccountID, params.Currency) {
		err := fmt.Errorf("Invalid FromAccountID %d", params.FromAccountID)
		return sqlc.TransferTxResult{}, err
	}
	if !s.validAccount(params.ToAccountID, params.Currency) {
		err := fmt.Errorf("Invalid ToAccountID %d", params.ToAccountID)
		return sqlc.TransferTxResult{}, err
	}
	transfer, err := s.store.TransferTx(context.Background(), *params)
	if err != nil {
		return sqlc.TransferTxResult{}, err
	}
	return transfer, nil
}

func (s *transferServiceImpl) GetListTransfer(params *sqlc.ListTransferParams) ([]sqlc.Transfer, error) {
	if params.Limit == 0 {
		params.Limit = 1000
	}
	accouts, err := s.store.ListTransfer(context.Background(), *params)
	if err != nil {
		return nil, err
	}
	return accouts, nil
}

func (s *transferServiceImpl) validAccount(accountID int64, currency string) bool {
	var c echo.Context
	account, err := s.store.GetAccount(context.Background(), accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err)
			return false
		}
		c.JSON(http.StatusInternalServerError, err)
	}
	if account.Currency != currency {
		//TODO:handlerから返されるエラーからこのメッセージを見つけられるようにする
		// err := fmt.Errorf("account [%d] currency mismatch : %s vs %s", account.ID, account.Currency, currency)
		// c.JSON(http.StatusBadRequest, err)
		return false
	}
	return true
}
