package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id           uint   `json:"id"          `  //
	Title        string `json:"title"        ` // 文章标题
	Keyword      string `json:"keyword"      ` // keywords
	Desc         string `json:"desc"         ` // 描述
	Content      string `json:"content"      ` // 文章内容,markdown格式
	UserId       int    `json:"userId"      `  // 文章编写人,对应users表
	CateId       int    `json:"cateId"      `  // 文章分类
	CommentCount int    `json:"commentCount" ` // 评论数量
	ReadCount    int    `json:"readCount"    ` // 阅读数量
	Status       int    `json:"status"       ` // 文章状态:0-公开;1-私密
	Sort         int    `json:"sort"         ` // 排序
	HtmlContent  string `json:"htmlContent"  ` // 文章内容,html格式
	ListPic      string `json:"listPic"      ` // 文章列表图
}

func init() {
	orm.RegisterModel(new(Article))
}

func (a *Article) GetAll() []Article {
	builder, err := orm.NewQueryBuilder("mysql")
	var articles []Article

	sql := builder.Select("*").From("articles").Where("id > ?").OrderBy("id").Desc().Limit(20).String()

	_, err = orm.NewOrm().Raw(sql, 0).QueryRows(&articles)

	if err != nil {
		panic(err)
	}

	return articles

}
