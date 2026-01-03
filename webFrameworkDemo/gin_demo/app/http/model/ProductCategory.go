package model

import "time"

// ProductCategory 商品分类表
type ProductCategory struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	Name        string    `gorm:"type:varchar(100);not null;column:name" json:"name"`                          // 分类名称
	ParentID    int       `gorm:"type:integer;default:0;column:parent_id" json:"parent_id"`                    // 父分类ID（0表示顶级分类）
	SortOrder   int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值（数字越小越靠前）
	Description string    `gorm:"type:text;column:description" json:"description"`                             // 分类描述
	ImageURL    string    `gorm:"type:varchar(500);column:image_url" json:"image_url"`                         // 分类图片URL
	Status      int8      `gorm:"type:tinyint;default:1;column:status" json:"status"`                          // 状态（1-启用，0-禁用）
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"` // 更新时间
}

func (p *ProductCategory) TableName() string {
	return "product_categories"
}
