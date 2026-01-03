package httpError

import (
	"net/http"
	"os"
)

const ERROR_CODE_FAILE int = 5000
const ERROR_CODE_SUCCESS int = 200
const ERROR_CODE_NOTALLOW int = 4003

type errorClass struct {
	codeMsg map[int]string
}

func (self errorClass) init() *errorClass {
	return &errorClass{map[int]string{
		ERROR_CODE_FAILE:    `abc`,
		ERROR_CODE_SUCCESS:  `OK`,
		ERROR_CODE_NOTALLOW: `不允许访问`,
	}}
}

// 异常处理
func ErrorHandle(writer http.ResponseWriter, err error) {

	code := http.StatusOK
	switch {
	case os.IsNotExist(err):
		code = http.StatusNotFound
	case os.IsPermission(err):
		code = http.StatusForbidden
	default:
		code = http.StatusInternalServerError
	}
	http.Error(writer, http.StatusText(code), http.StatusNotFound)
}
