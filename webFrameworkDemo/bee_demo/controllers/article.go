package controllers

import (
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Add() {

	type user struct {
		Id     int
		Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
		Age    int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140，超出此范围即为不合法
		Email  string `valid:"Email; MaxSize(100)"`      // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
		Mobile string `valid:"Mobile"`                   // Mobile 必须为正确的手机号
		IP     string `valid:"IP"`                       // IP 必须为一个正确的 IPv4 地址
	}

	u := user{
		Id:     0,
		Name:   "",
		Age:    3,
		Email:  "dev@.com",
		Mobile: "186",
		IP:     "0.00.2",
	}
	v := validation.Validation{}
	v.Valid(u)

	if v.HasErrors() {
		a.JSONResp(v.Errors)
		return
	}

	return

}
