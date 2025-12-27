package model

import "time"

// ProductBrand 商品品牌表
type ProductBrand struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	Name        string    `gorm:"type:varchar(100);not null;column:name" json:"name"`                          // 品牌名称
	LogoURL     string    `gorm:"type:varchar(500);column:logo_url" json:"logo_url"`                           // 品牌Logo URL
	Description string    `gorm:"type:text;column:description" json:"description"`                             // 品牌描述
	SortOrder   int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值
	Status      int8      `gorm:"type:tinyint;default:1;column:status" json:"status"`                          // 状态（1-启用，0-禁用）
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"` // 更新时间
}

func (p *ProductBrand) TableName() string {
	return "product_brands"
}
