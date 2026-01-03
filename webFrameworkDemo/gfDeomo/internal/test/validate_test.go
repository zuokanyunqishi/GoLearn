package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestValidate(t *testing.T) {

	type req struct {
		Name string
		Age  int
		Sex  int8
	}

	rule := map[string]string{
		"Name": "required|length:20,30",
		"Age":  "required|min:1",
		"sex":  "required",
	}

	message := map[string]interface{}{
		"Age": map[string]string{
			"required": "年级必须",
			"min":      "年级最小是1",
		},
		"Name": map[string]string{
			"length":   "名字长度{min},{max}",
			"required": "姓名必须",
		},
		"Sex": "性别不能为空",
	}

	reqObj := req{
		Age:  0,
		Sex:  3,
		Name: "",
	}

	err := g.Validator().Rules(rule).Messages(message).Data(reqObj).Run(gctx.New())

	if err != nil {
		g.Dump(err.Maps())
	}

}
