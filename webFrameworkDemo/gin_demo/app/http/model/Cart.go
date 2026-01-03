package model

import (
	app "speed/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Cart 购物车表
type Cart struct {
	ID         int            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                    // 主键
	UserID     int            `gorm:"type:integer;not null;column:user_id;index" json:"user_id"`       // 用户ID
	ProductID  int            `gorm:"type:integer;not null;column:product_id;index" json:"product_id"` // 商品ID
	SkuID      *int           `gorm:"type:integer;column:sku_id;index" json:"sku_id"`                  // SKU ID（可选）
	Quantity   int            `gorm:"type:integer;not null;default:1;column:quantity" json:"quantity"` // 商品数量
	IsSelected int8           `gorm:"type:tinyint;default:1;column:is_selected" json:"is_selected"`    // 是否选中（1-选中，0-未选中）
	CreatedAt  time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime;column:deleted_at;index" json:"deleted_at,omitempty"` // 删除时间（软删除）

	// 关联关系
	User    User        `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Product Product     `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
	Sku     *ProductSku `gorm:"foreignKey:SkuID;references:ID" json:"sku,omitempty"`
}

func (c *Cart) TableName() string {
	return "carts"
}

// GetCartByID 根据ID获取购物车记录
func (c *Cart) GetCartByID(ctx *gin.Context, id int) error {
	tx := app.Db.WithContext(ctx).Where("id = ?", id).First(&c)
	return tx.Error
}

// GetCartByUserProductSku 根据用户ID、商品ID和SKU ID获取购物车记录
func (c *Cart) GetCartByUserProductSku(ctx *gin.Context, userID, productID int, skuID *int) error {
	query := app.Db.WithContext(ctx).Where("user_id = ? AND product_id = ?", userID, productID)
	if skuID != nil {
		query = query.Where("sku_id = ?", *skuID)
	} else {
		query = query.Where("sku_id IS NULL")
	}
	tx := query.First(&c)
	return tx.Error
}

// Add 添加购物车记录
func (c *Cart) Add(ctx *gin.Context) error {
	tx := app.Db.WithContext(ctx).Create(&c)
	return tx.Error
}

// Update 更新购物车记录
func (c *Cart) Update(ctx *gin.Context) error {
	tx := app.Db.WithContext(ctx).Model(&c).Where("id = ?", c.ID).Updates(&c)
	return tx.Error
}

// Delete 删除购物车记录（软删除）
func (c *Cart) Delete(ctx *gin.Context) error {
	tx := app.Db.WithContext(ctx).Delete(&c)
	return tx.Error
}

// GetCartListByUserID 获取用户的购物车列表
func GetCartListByUserID(ctx *gin.Context, userID int) ([]Cart, error) {
	var carts []Cart
	tx := app.Db.WithContext(ctx).
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Brand").
		Preload("Sku").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&carts)
	return carts, tx.Error
}

// GetCartCountByUserID 获取用户购物车商品数量
func GetCartCountByUserID(ctx *gin.Context, userID int) (int64, error) {
	var count int64
	tx := app.Db.WithContext(ctx).Model(&Cart{}).Where("user_id = ?", userID).Count(&count)
	return count, tx.Error
}
