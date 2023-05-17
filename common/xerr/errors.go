package xerr

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{
		errCode: SERVER_COMMON_ERROR,
		errMsg:  errMsg,
	}
}
