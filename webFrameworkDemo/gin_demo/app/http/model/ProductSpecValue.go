package model

import "time"

// ProductSpecValue 商品规格值表（规格的具体值，如：红色、蓝色、M码、L码等）
type ProductSpecValue struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	SpecID    int       `gorm:"type:integer;not null;column:spec_id;index" json:"spec_id"`                   // 规格ID
	Value     string    `gorm:"type:varchar(100);not null;column:value" json:"value"`                        // 规格值
	SortOrder int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值
	Status    int8      `gorm:"type:tinyint;default:1;column:status" json:"status"`                          // 状态（1-启用，0-禁用）
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间

	// 关联关系
	Spec ProductSpec `gorm:"foreignKey:SpecID;references:ID" json:"spec,omitempty"`
}

func (p *ProductSpecValue) TableName() string {
	return "product_spec_values"
}
