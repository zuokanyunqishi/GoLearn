package util

import (
	clientError "GoLearn/chat/client/errors"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/errors"
	"encoding/json"
	"time"
)

type BasRouter struct {
	handleError interfaces.Error
	handleOk    bool
	okChan      chan struct{}
}

// 基类Router
//
// 处理业务之前钩子
func (b *BasRouter) PreHandle(request interfaces.Request) {
}

// 处理业务中前钩子
func (b *BasRouter) Handle(request interfaces.Request) {
}

// 处理业务后钩子
func (b *BasRouter) PostHandle(request interfaces.Request) {
}

func (b *BasRouter) SetHandleOK(ok bool) {
	b.handleOk = ok
}

func (b *BasRouter) HandleIsOk() bool {
	return b.handleOk
}

func (b *BasRouter) SetHandleErr(error interfaces.Error) {
	b.handleError = error
}

func (b *BasRouter) GetHandleErr() interfaces.Error {
	return b.handleError
}

func (b *BasRouter) HandleResponse() interfaces.Error {
	after := b.GetTimer()
	defer after.Stop()
	if b.okChan == nil {
		b.okChan = make(chan struct{})
	}
	for {
		if b.handleError != nil {
			after.Stop()
			return b.handleError
		}

		select {
		case <-b.okChan:
			return nil
		case <-after.C:
			return clientError.ErrGetResponseTimeOut
		default:
			continue
		}
	}
}

func (b *BasRouter) AfterBusinessHandle() {
	b.handleOk = false
	b.handleError = nil
}

func (b *BasRouter) GetTimer() *time.Timer {
	return time.NewTimer(time.Second * 4)
}

func (b *BasRouter) UnmarshalRequest(request interfaces.Request, response interface{}) {
	err := json.Unmarshal(request.GetData(), response)
	if err != nil {
		b.SetHandleErr(errors.New(
			"json.Unmarshal(request.GetData(), &response "+err.Error(),
			clientError.ErrCodePackRequestMsg))
		return
	}
}

func (b *BasRouter) SendHandleOk() {
	b.okChan <- struct{}{}
}
