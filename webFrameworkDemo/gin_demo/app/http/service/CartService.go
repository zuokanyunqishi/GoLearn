package service

import (
	"errors"
	"speed/app/http/model"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartService struct {
}

func NewCartService() *CartService {
	return &CartService{}
}

// AddCartRequest 添加购物车请求参数
type AddCartRequest struct {
	ProductID int  `json:"product_id" binding:"required,min=1" comment:"商品ID"`
	SkuID     *int `json:"sku_id" binding:"omitempty,min=1" comment:"SKU ID（可选）"`
	Quantity  int  `json:"quantity" binding:"required,min=1,max=999" comment:"商品数量"`
}

// UpdateCartRequest 更新购物车请求参数
type UpdateCartRequest struct {
	Quantity   *int  `json:"quantity" binding:"omitempty,min=1,max=999" comment:"商品数量"`
	IsSelected *int8 `json:"is_selected" binding:"omitempty,oneof=0 1" comment:"是否选中（1-选中，0-未选中）"`
}

// CartResponse 购物车响应结构
type CartResponse struct {
	ID         int    `json:"id"`          // 主键
	UserID     int    `json:"user_id"`     // 用户ID
	ProductID  int    `json:"product_id"`  // 商品ID
	SkuID      *int   `json:"sku_id"`      // SKU ID
	Quantity   int    `json:"quantity"`    // 商品数量
	IsSelected int8   `json:"is_selected"` // 是否选中
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 更新时间
	// 商品信息
	Product ProductResponse `json:"product"` // 商品信息
	// SKU信息
	Sku *ProductSkuResponse `json:"sku,omitempty"` // SKU信息
}

// toCartResponse 将 model.Cart 转换为 CartResponse
func toCartResponse(cart model.Cart) CartResponse {
	response := CartResponse{
		ID:         cart.ID,
		UserID:     cart.UserID,
		ProductID:  cart.ProductID,
		SkuID:      cart.SkuID,
		Quantity:   cart.Quantity,
		IsSelected: cart.IsSelected,
		CreatedAt:  formatTime(cart.CreatedAt),
		UpdatedAt:  formatTime(cart.UpdatedAt),
	}

	// 转换商品信息
	if cart.Product.ID > 0 {
		response.Product = toProductResponse(cart.Product)
	}

	// 转换SKU信息
	if cart.Sku != nil && cart.Sku.ID > 0 {
		skuResp := toProductSkuResponse(*cart.Sku)
		response.Sku = &skuResp
	}

	return response
}

// AddCart 添加商品到购物车
func (s *CartService) AddCart(ctx *gin.Context, userID int, req AddCartRequest) error {
	// 检查商品是否存在
	var product model.Product
	if err := app.Db.WithContext(ctx).Where("id = ? AND status = ?", req.ProductID, 1).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("商品不存在或已下架")
		}
		return err
	}

	// 如果提供了SKU ID，检查SKU是否存在
	if req.SkuID != nil {
		var sku model.ProductSku
		if err := app.Db.WithContext(ctx).Where("id = ? AND product_id = ? AND status = ?", *req.SkuID, req.ProductID, 1).First(&sku).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("SKU不存在或已禁用")
			}
			return err
		}

		// 检查SKU库存
		if sku.StockQuantity < req.Quantity {
			return errors.New("SKU库存不足")
		}
	} else {
		// 检查商品库存
		if product.StockQuantity < req.Quantity {
			return errors.New("商品库存不足")
		}
	}

	// 检查购物车中是否已存在相同商品和SKU
	var existingCart model.Cart
	err := existingCart.GetCartByUserProductSku(ctx, userID, req.ProductID, req.SkuID)
	if err == nil {
		// 已存在，更新数量
		existingCart.Quantity += req.Quantity
		// 再次检查库存
		if req.SkuID != nil {
			var sku model.ProductSku
			app.Db.WithContext(ctx).Where("id = ?", *req.SkuID).First(&sku)
			if sku.StockQuantity < existingCart.Quantity {
				return errors.New("SKU库存不足")
			}
		} else {
			if product.StockQuantity < existingCart.Quantity {
				return errors.New("商品库存不足")
			}
		}
		return existingCart.Update(ctx)
	} else if err != gorm.ErrRecordNotFound {
		// 其他错误
		return err
	}

	// 不存在，创建新记录
	cart := model.Cart{
		UserID:     userID,
		ProductID:  req.ProductID,
		SkuID:      req.SkuID,
		Quantity:   req.Quantity,
		IsSelected: 1,
	}

	return cart.Add(ctx)
}

// UpdateCart 更新购物车
func (s *CartService) UpdateCart(ctx *gin.Context, userID, cartID int, req UpdateCartRequest) error {
	// 获取购物车记录
	var cart model.Cart
	if err := cart.GetCartByID(ctx, cartID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("购物车记录不存在")
		}
		return err
	}

	// 验证是否为当前用户的购物车
	if cart.UserID != userID {
		return errors.New("无权操作此购物车记录")
	}

	// 更新数量
	if req.Quantity != nil {
		// 检查库存
		if cart.SkuID != nil {
			var sku model.ProductSku
			if err := app.Db.WithContext(ctx).Where("id = ?", *cart.SkuID).First(&sku).Error; err != nil {
				return errors.New("SKU不存在")
			}
			if sku.StockQuantity < *req.Quantity {
				return errors.New("SKU库存不足")
			}
		} else {
			var product model.Product
			if err := app.Db.WithContext(ctx).Where("id = ?", cart.ProductID).First(&product).Error; err != nil {
				return errors.New("商品不存在")
			}
			if product.StockQuantity < *req.Quantity {
				return errors.New("商品库存不足")
			}
		}
		cart.Quantity = *req.Quantity
	}

	// 更新选中状态
	if req.IsSelected != nil {
		cart.IsSelected = *req.IsSelected
	}

	return cart.Update(ctx)
}

// DeleteCart 删除购物车记录
func (s *CartService) DeleteCart(ctx *gin.Context, userID, cartID int) error {
	// 获取购物车记录
	var cart model.Cart
	if err := cart.GetCartByID(ctx, cartID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("购物车记录不存在")
		}
		return err
	}

	// 验证是否为当前用户的购物车
	if cart.UserID != userID {
		return errors.New("无权操作此购物车记录")
	}

	return cart.Delete(ctx)
}

// GetCartList 获取购物车列表
func (s *CartService) GetCartList(ctx *gin.Context, userID int) ([]CartResponse, error) {
	carts, err := model.GetCartListByUserID(ctx, userID)
	if err != nil {
		app.Log.Error("获取购物车列表失败", zap.Error(err))
		return nil, err
	}

	// 转换为响应格式
	cartResponses := make([]CartResponse, len(carts))
	for i, cart := range carts {
		cartResponses[i] = toCartResponse(cart)
	}

	return cartResponses, nil
}

// GetCartCount 获取购物车商品数量
func (s *CartService) GetCartCount(ctx *gin.Context, userID int) (int64, error) {
	return model.GetCartCountByUserID(ctx, userID)
}

// BatchUpdateSelected 批量更新选中状态
func (s *CartService) BatchUpdateSelected(ctx *gin.Context, userID int, cartIDs []int, isSelected int8) error {
	if len(cartIDs) == 0 {
		return errors.New("购物车ID列表不能为空")
	}

	// 更新选中状态
	tx := app.Db.WithContext(ctx).
		Model(&model.Cart{}).
		Where("user_id = ? AND id IN ?", userID, cartIDs).
		Update("is_selected", isSelected)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("未找到可更新的购物车记录")
	}

	return nil
}

// ClearCart 清空购物车（删除所有未选中的记录，或删除所有记录）
func (s *CartService) ClearCart(ctx *gin.Context, userID int, onlyUnselected bool) error {
	query := app.Db.WithContext(ctx).Where("user_id = ?", userID)
	if onlyUnselected {
		query = query.Where("is_selected = ?", 0)
	}
	tx := query.Delete(&model.Cart{})
	return tx.Error
}
