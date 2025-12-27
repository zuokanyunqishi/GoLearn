package model

import "time"

// ProductSku 商品SKU表（库存量单位，如：不同颜色、尺寸等）
type ProductSku struct {
	ID            int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                         // 主键
	ProductID     int       `gorm:"type:integer;not null;column:product_id;index" json:"product_id"`                      // 商品ID
	SkuCode       string    `gorm:"type:varchar(100);uniqueIndex;not null;column:sku_code" json:"sku_code"`               // SKU编码
	SkuName       string    `gorm:"type:varchar(200);column:sku_name" json:"sku_name"`                                    // SKU名称（如：红色-M码）
	SpecValues    string    `gorm:"type:text;column:spec_values" json:"spec_values"`                                      // 规格值（JSON格式存储，如：{"颜色":"红色","尺寸":"M"}）
	OriginalPrice float64   `gorm:"type:decimal(10,2);not null;default:0.00;column:original_price" json:"original_price"` // SKU原价
	SalePrice     float64   `gorm:"type:decimal(10,2);not null;default:0.00;column:sale_price" json:"sale_price"`         // SKU售价
	CostPrice     *float64  `gorm:"type:decimal(10,2);column:cost_price" json:"cost_price"`                               // SKU成本价
	StockQuantity int       `gorm:"type:integer;default:0;column:stock_quantity" json:"stock_quantity"`                   // SKU库存数量
	ImageURL      string    `gorm:"type:varchar(500);column:image_url" json:"image_url"`                                  // SKU图片URL
	Status        int8      `gorm:"type:tinyint;default:1;column:status;index" json:"status"`                             // 状态（1-启用，0-禁用）
	SortOrder     int       `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                           // 排序值
	CreatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`          // 创建时间
	UpdatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`          // 更新时间

	// 关联关系
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
}

func (p *ProductSku) TableName() string {
	return "product_skus"
}
