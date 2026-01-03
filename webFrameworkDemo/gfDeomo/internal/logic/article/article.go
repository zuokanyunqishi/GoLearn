package article

import (
	"context"
	v1 "gfDeomo/api/v1"
	"gfDeomo/internal/dao"
	"gfDeomo/internal/service"
)

type sArticle struct{}

func New() *sArticle {
	return &sArticle{}
}

func init() {
	service.RegisterArticle(New())
}

// List 文章列表
func (a *sArticle) List(ctx context.Context, req v1.ContentGetListCommonReq) (v1.ArticleIndexRes, error) {

	articles := dao.Articles.Get(ctx)

	return articles, nil
}
