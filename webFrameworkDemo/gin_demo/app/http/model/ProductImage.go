package model

import "time"

// ProductImage 商品图片表（额外商品图片）
type ProductImage struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	ProductID int       `gorm:"type:integer;not null;column:product_id;index" json:"product_id"`             // 商品ID
	ImageURL  string    `gorm:"type:varchar(500);not null;column:image_url" json:"image_url"`                // 图片URL
	SortOrder int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值
	IsMain    int8      `gorm:"type:tinyint;default:0;column:is_main;index" json:"is_main"`                  // 是否主图（1-是，0-否）
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间

	// 关联关系
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
}

func (p *ProductImage) TableName() string {
	return "product_images"
}
