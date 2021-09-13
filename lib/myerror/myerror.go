package myerror

import "net/http"

type MyError interface {
	Error() string
}

type myError struct {
	CodeMessage string `json:"error"`
	Message     string `json:"message"`
}

func (e *myError) Error() string {
	return e.Message
}

//return errをするとlogrusがstatusCodeを常に200で受け取ってしまうため、c.JSONを使う必要がある
//しかしc.JSON(500,err)の形だとメッセージを意図した通りに表示してくれない
//logrusを正しく動かす＋return errと同じメッセージをフロントで受け取るためにmyerrorがある
func Set(code int, err error) (int, myError) {
	return code, myError{
		CodeMessage: http.StatusText(code),
		Message:     err.Error(),
	}
}
