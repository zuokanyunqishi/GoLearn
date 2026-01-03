package controllers

import (
	"errors"
	"net/http"
	"speed/app/http/service"
	"speed/app/lib/validate"
	app "speed/bootstrap"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductController struct {
	Controller
	productService *service.ProductService
}

var ProductC = &ProductController{productService: &service.ProductService{}}

// List 商品列表接口
// @Summary 获取商品列表
// @Description 支持分页、筛选、排序的商品列表接口
// @Tags 商品
// @Accept json
// @Produce json
// @Param page query int false "页码，默认1"
// @Param page_size query int false "每页数量，默认10，最大100"
// @Param category_id query int false "分类ID"
// @Param brand_id query int false "品牌ID"
// @Param status query int false "状态（1-上架，0-下架，2-待审核，3-审核失败）"
// @Param is_hot query int false "是否热门（1-是，0-否）"
// @Param is_new query int false "是否新品（1-是，0-否）"
// @Param is_recommend query int false "是否推荐（1-是，0-否）"
// @Param keyword query string false "关键词搜索（商品名称、SPU编码）"
// @Param sort_by query string false "排序字段（created_at, sale_count, view_count, sale_price）"
// @Param sort_order query string false "排序方式（asc, desc），默认desc"
// @Success 200 {object} service.ProductListResponse
// @Router /api/products [get]
func (p *ProductController) List(ctx *gin.Context) {
	var req service.ProductListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		p.handleValidationError(ctx, err)
		return
	}

	result, err := p.productService.GetProductList(ctx, req)
	if err != nil {
		app.Log.Error("获取商品列表失败", zap.Error(err))
		p.ResponseError(ctx, http.StatusInternalServerError, "获取商品列表失败")
		return
	}

	p.ResponseSuccess(ctx, result)
}

// Detail 商品详情接口
// @Summary 获取商品详情
// @Description 根据商品ID获取商品详细信息，包括分类、品牌、SKU、图片、属性等
// @Tags 商品
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} service.ProductDetailResponse
// @Router /api/products/{id} [get]
func (p *ProductController) Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		p.ResponseError(ctx, http.StatusBadRequest, "商品ID格式错误")
		return
	}

	product, err := p.productService.GetProductDetail(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			p.ResponseError(ctx, http.StatusNotFound, "商品不存在")
			return
		}
		app.Log.Error("获取商品详情失败", zap.Error(err))
		p.ResponseError(ctx, http.StatusInternalServerError, "获取商品详情失败")
		return
	}

	p.ResponseSuccess(ctx, product)
}

func (p *ProductController) handleValidationError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		app.Log.Error("Invalid request format", zap.Error(err))
		p.ResponseError(ctx, http.StatusBadRequest, "请求格式错误")
		return
	}

	translatedErrs := validate.TranslateError(errs)
	app.Log.Warn("Validation failed", zap.Any("errors", translatedErrs))
	p.ResponseError(ctx, http.StatusBadRequest, translatedErrs)
}
