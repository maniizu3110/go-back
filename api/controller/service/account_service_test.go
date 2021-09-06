package service_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"simplebank/api"
	"simplebank/api/sqlc"
	"simplebank/api/util"
	mockdb "simplebank/db/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)


func TestGetAccount(t *testing.T){
	account := randomAccount()
	ctrl := gomock.NewController(t)
	//これをNew関数に渡せばうまくserviceの形にできそう
	store := mockdb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().GetAccount(gomock.Any(),gomock.Eq(account.ID)).Times(1).Return(account,nil)

	server := api.NewServer(store)
	recoder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/v1/account/%d",account.ID)
	fmt.Println(url)
	request,err := http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)

	server.SetRouter().ServeHTTP(recoder,request)
	//check response
	require.Equal(t,http.StatusOK,recoder.Code)
	requireBodyMatchAccount(t,recoder.Body,account)

}


func randomAccount() sqlc.Account{
	return sqlc.Account{
		ID: util.RandomInt(1,1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account sqlc.Account){
	data,err := ioutil.ReadAll(body)
	require.NoError(t,err)

	var gotAccount sqlc.Account
	err = json.Unmarshal(data,&gotAccount)
	require.NoError(t,err)
	require.Equal(t,account,gotAccount)
}