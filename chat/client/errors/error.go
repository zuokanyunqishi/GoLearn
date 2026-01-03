package errors

import "GoLearn/chat/util/errors"

const ErrCodeRequestMarshal = 2001
const ErrCodePackRequestMsg = 2002
const ErrCodeLinkServerNotInit = 3000
const ErrCodeLoginResponse = 4000

const ErrCodeRegisterResponse = 4001

var ErrLoginRequestMarshal = errors.New("login.Request Marshal error", ErrCodeRequestMarshal)
var ErrGetResponseTimeOut = errors.NewNoCode("get response content time out")
