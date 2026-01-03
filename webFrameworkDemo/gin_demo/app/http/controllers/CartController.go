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
)

type CartController struct {
	Controller
	cartService *service.CartService
}

var CartC = &CartController{cartService: service.NewCartService()}

// Add 添加商品到购物车
// @Summary 添加商品到购物车
// @Description 将商品添加到购物车，如果已存在相同商品和SKU，则增加数量
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param body body service.AddCartRequest true "添加购物车请求"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart [post]
func (c *CartController) Add(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	var req service.AddCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}

	err := c.cartService.AddCart(ctx, userID.(int), req)
	if err != nil {
		app.Log.Error("添加购物车失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	c.ResponseSuccess(ctx, gin.H{"message": "添加成功"})
}

// Update 更新购物车
// @Summary 更新购物车
// @Description 更新购物车商品数量或选中状态
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "购物车ID"
// @Param body body service.UpdateCartRequest true "更新购物车请求"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart/{id} [post]
func (c *CartController) Update(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	// 获取购物车ID
	idStr := ctx.Param("id")
	cartID, err := strconv.Atoi(idStr)
	if err != nil {
		c.ResponseError(ctx, http.StatusBadRequest, "购物车ID格式错误")
		return
	}

	var req service.UpdateCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}

	err = c.cartService.UpdateCart(ctx, userID.(int), cartID, req)
	if err != nil {
		app.Log.Error("更新购物车失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	c.ResponseSuccess(ctx, gin.H{"message": "更新成功"})
}

// Delete 删除购物车记录
// @Summary 删除购物车记录
// @Description 删除指定的购物车记录
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "购物车ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart/{id} [post]
func (c *CartController) Delete(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	// 获取购物车ID
	idStr := ctx.Param("id")
	cartID, err := strconv.Atoi(idStr)
	if err != nil {
		c.ResponseError(ctx, http.StatusBadRequest, "购物车ID格式错误")
		return
	}

	err = c.cartService.DeleteCart(ctx, userID.(int), cartID)
	if err != nil {
		app.Log.Error("删除购物车失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	c.ResponseSuccess(ctx, gin.H{"message": "删除成功"})
}

// List 获取购物车列表
// @Summary 获取购物车列表
// @Description 获取当前用户的购物车列表，包含商品和SKU信息
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart [get]
func (c *CartController) List(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	carts, err := c.cartService.GetCartList(ctx, userID.(int))
	if err != nil {
		app.Log.Error("获取购物车列表失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "获取购物车列表失败")
		return
	}

	c.ResponseSuccess(ctx, carts)
}

// Count 获取购物车商品数量
// @Summary 获取购物车商品数量
// @Description 获取当前用户购物车中的商品数量
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart/count [get]
func (c *CartController) Count(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	count, err := c.cartService.GetCartCount(ctx, userID.(int))
	if err != nil {
		app.Log.Error("获取购物车数量失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "获取购物车数量失败")
		return
	}

	c.ResponseSuccess(ctx, gin.H{"count": count})
}

// BatchUpdateSelected 批量更新选中状态
// @Summary 批量更新选中状态
// @Description 批量更新购物车商品的选中状态
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param body body map[string]interface{} true "请求体，包含cart_ids数组和is_selected字段"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart/batch-selected [post]
func (c *CartController) BatchUpdateSelected(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	var req struct {
		CartIDs    []int `json:"cart_ids" binding:"required,min=1" comment:"购物车ID列表"`
		IsSelected int8  `json:"is_selected" binding:"required,oneof=0 1" comment:"是否选中（1-选中，0-未选中）"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}

	err := c.cartService.BatchUpdateSelected(ctx, userID.(int), req.CartIDs, req.IsSelected)
	if err != nil {
		app.Log.Error("批量更新选中状态失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	c.ResponseSuccess(ctx, gin.H{"message": "更新成功"})
}

// Clear 清空购物车
// @Summary 清空购物车
// @Description 清空购物车，可选择只删除未选中的商品或删除所有商品
// @Tags 购物车
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param only_unselected query bool false "是否只删除未选中的商品，默认false（删除所有）"
// @Success 200 {object} map[string]interface{}
// @Router /api/cart/clear [post]
func (c *CartController) Clear(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("userId")
	if !exists {
		c.ResponseError(ctx, http.StatusUnauthorized, "用户未登录")
		return
	}

	onlyUnselected := ctx.DefaultQuery("only_unselected", "false") == "true"

	err := c.cartService.ClearCart(ctx, userID.(int), onlyUnselected)
	if err != nil {
		app.Log.Error("清空购物车失败", zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "清空购物车失败")
		return
	}

	c.ResponseSuccess(ctx, gin.H{"message": "清空成功"})
}

func (c *CartController) handleValidationError(ctx *gin.Context, err error) {
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
