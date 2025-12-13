package validate

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func Init() {
	Validate = validator.New()
	uni := ut.New(zh.New(), zh.New())
	Trans, _ = uni.GetTranslator("zh")

	// 注册中文翻译到Gin的绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zhtranslations.RegisterDefaultTranslations(v, Trans)
	}
}

// TranslateError 将验证错误翻译为中文字符串
func TranslateError(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var result string
		for _, e := range errs {
			result += e.Translate(Trans) + "; "
		}
		return result
	}
	return err.Error()
}
