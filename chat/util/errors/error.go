package errors

import (
	"GoLearn/chat/interfaces"
	"fmt"
)

type Error struct {
	ErrorCode uint32
	ErrMsg    string
	trace     string
}

func New(errMsg string, errCode uint32) interfaces.Error {
	return Error{
		ErrorCode: errCode,
		ErrMsg:    errMsg,
	}
}

func NewNoCode(errMsg string) interfaces.Error {
	return Error{
		ErrMsg: errMsg,
	}
}

func (err Error) String() string {
	return fmt.Sprintf("errcode %d errorMsg %s", err.ErrorCode, err.ErrMsg)
}

func (err Error) Error() string {
	return err.ErrMsg
}

func (err Error) SetTrace(trace string) {
	err.trace = trace
}

func (err Error) GetTrace() string {
	return err.trace
}

func (err Error) GetErrorCode() uint32 {
	return err.ErrorCode
}
