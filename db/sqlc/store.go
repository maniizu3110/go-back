package db

import (
	"context"
	"database/sql"
	"fmt"
)

//自動生成コードは編集することはできないのでここでNewの処理を書いているQueriesをmapしている
//Store provides all functions to execure db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

//txはトランザクション
//ececTx executes a function within a database transacsion
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//このNewはNewStoreとは別物（*sql.DB,*sql.Txはどちらもインターフェースを満たしている）
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v,rb err: %v",err,rbErr)
		}
		return err
	}
	return tx.Commit()
}
