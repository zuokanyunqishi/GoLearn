package model

import "time"

// ProductSpec 商品规格表（用于定义商品的规格属性，如：颜色、尺寸等）
type ProductSpec struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	Name      string    `gorm:"type:varchar(50);not null;column:name" json:"name"`                           // 规格名称（如：颜色、尺寸）
	SortOrder int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值
	Status    int8      `gorm:"type:tinyint;default:1;column:status" json:"status"`                          // 状态（1-启用，0-禁用）
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"` // 更新时间

	// 关联关系
	SpecValues []ProductSpecValue `gorm:"foreignKey:SpecID;references:ID" json:"spec_values,omitempty"`
}

func (p *ProductSpec) TableName() string {
	return "product_specs"
}
