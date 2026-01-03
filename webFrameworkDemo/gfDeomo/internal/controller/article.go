package controller

import (
	"context"
	v1 "gfDeomo/api/v1"
	"gfDeomo/internal/service"
)

var (
	Article = new(article)
)

type article struct{}

func (a *article) Index2(ctx context.Context, req *v1.ArticleIndexReq) (v1.ArticleIndexRes, error) {
	list, err := service.Article().List(ctx, v1.ContentGetListCommonReq{})
	return list, err
}
