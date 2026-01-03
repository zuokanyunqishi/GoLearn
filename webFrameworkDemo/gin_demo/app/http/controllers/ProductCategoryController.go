package controllers

import (
	"errors"
	"net/http"
	"speed/app/http/service"
	"speed/app/lib/validate"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type ProductCategoryController struct {
	Controller
	categoryService *service.ProductCategoryService
}

var ProductCategoryC = &ProductCategoryController{
	categoryService: &service.ProductCategoryService{},
}

// List 商品分类列表接口
// @Summary 获取商品分类列表
// @Description 支持按父分类ID查询子分类，支持状态筛选。如果 parent_id 为 0 或不传，返回树形结构的分类列表
// @Tags 商品分类
// @Accept json
// @Produce json
// @Param parent_id query int false "父分类ID（0或不传表示查询顶级分类）"
// @Param status query int false "状态（1-启用，0-禁用），不传则查询所有"
// @Success 200 {object} service.CategoryListResponse
// @Router /api/categories [get]
func (c *ProductCategoryController) List(ctx *gin.Context) {
	var req service.CategoryListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}

	result, err := c.categoryService.GetCategoryList(ctx, req)
	if err != nil {
		app.Log.Error("获取商品分类列表失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "获取商品分类列表失败")
		return
	}

	c.ResponseSuccess(ctx, result)
}

// Tree 获取完整分类树接口
// @Summary 获取完整的商品分类树
// @Description 返回所有分类的完整树形结构（包括禁用的分类）
// @Tags 商品分类
// @Accept json
// @Produce json
// @Success 200 {object} service.CategoryListResponse
// @Router /api/categories/tree [get]
func (c *ProductCategoryController) Tree(ctx *gin.Context) {
	result, err := c.categoryService.GetCategoryTree(ctx)
	if err != nil {
		app.Log.Error("获取商品分类树失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "获取商品分类树失败")
		return
	}

	c.ResponseSuccess(ctx, result)
}

func (c *ProductCategoryController) handleValidationError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		app.Log.Error("Invalid request format", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, "请求格式错误")
		return
	}

	translatedErrs := validate.TranslateError(errs)
	app.Log.Warn("Validation failed", zap.Any("errors", translatedErrs))
	c.ResponseError(ctx, http.StatusBadRequest, translatedErrs)
}
