package model

import "time"

// ProductAttribute 商品属性表（商品的扩展属性，如：产地、材质等）
type ProductAttribute struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                // 主键
	ProductID int       `gorm:"type:integer;not null;column:product_id;index" json:"product_id"`             // 商品ID
	AttrName  string    `gorm:"type:varchar(100);not null;column:attr_name" json:"attr_name"`                // 属性名称
	AttrValue string    `gorm:"type:text;not null;column:attr_value" json:"attr_value"`                      // 属性值
	SortOrder int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                  // 排序值
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"` // 创建时间

	// 关联关系
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
}

func (p *ProductAttribute) TableName() string {
	return "product_attributes"
}
