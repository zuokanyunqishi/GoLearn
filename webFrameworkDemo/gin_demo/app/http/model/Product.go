package model

import (
	"time"

	"gorm.io/gorm"
)

// Product 商品表（主表）
type Product struct {
	ID             int            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                         // 主键
	SpuCode        string         `gorm:"type:varchar(100);uniqueIndex;not null;column:spu_code" json:"spu_code"`               // SPU编码（商品唯一标识）
	Name           string         `gorm:"type:varchar(200);not null;column:name" json:"name"`                                   // 商品名称
	Subtitle       string         `gorm:"type:varchar(500);column:subtitle" json:"subtitle"`                                    // 商品副标题/简短描述
	Description    string         `gorm:"type:text;column:description" json:"description"`                                      // 商品详细描述
	CategoryID     int            `gorm:"type:integer;not null;column:category_id;index" json:"category_id"`                    // 商品分类ID
	BrandID        *int           `gorm:"type:integer;column:brand_id;index" json:"brand_id"`                                   // 品牌ID
	MainImageURL   string         `gorm:"type:varchar(500);column:main_image_url" json:"main_image_url"`                        // 主图URL
	ImageURLs      string         `gorm:"type:text;column:image_urls" json:"image_urls"`                                        // 商品图片URL列表（JSON格式存储）
	OriginalPrice  float64        `gorm:"type:decimal(10,2);not null;default:0.00;column:original_price" json:"original_price"` // 商品原价
	SalePrice      float64        `gorm:"type:decimal(10,2);not null;default:0.00;column:sale_price" json:"sale_price"`         // 商品售价
	CostPrice      *float64       `gorm:"type:decimal(10,2);column:cost_price" json:"cost_price"`                               // 成本价
	StockQuantity  int            `gorm:"type:integer;default:0;column:stock_quantity" json:"stock_quantity"`                   // 库存数量
	SaleCount      int            `gorm:"type:integer;default:0;column:sale_count" json:"sale_count"`                           // 销售数量
	ViewCount      int            `gorm:"type:integer;default:0;column:view_count" json:"view_count"`                           // 浏览量
	Weight         *float64       `gorm:"type:decimal(10,2);column:weight" json:"weight"`                                       // 商品重量（单位：kg）
	Volume         *float64       `gorm:"type:decimal(10,2);column:volume" json:"volume"`                                       // 商品体积（单位：立方米）
	Unit           string         `gorm:"type:varchar(20);default:件;column:unit" json:"unit"`                                   // 单位（件、台、个等）
	Status         int8           `gorm:"type:tinyint;default:1;column:status;index" json:"status"`                             // 商品状态（1-上架，0-下架，2-待审核，3-审核失败）
	IsHot          int8           `gorm:"type:tinyint;default:0;column:is_hot;index" json:"is_hot"`                             // 是否热门（1-是，0-否）
	IsNew          int8           `gorm:"type:tinyint;default:0;column:is_new;index" json:"is_new"`                             // 是否新品（1-是，0-否）
	IsRecommend    int8           `gorm:"type:tinyint;default:0;column:is_recommend;index" json:"is_recommend"`                 // 是否推荐（1-是，0-否）
	SortOrder      int            `gorm:"type:integer;default:0;column:sort_order" json:"sort_order"`                           // 排序值（数字越小越靠前）
	SeoTitle       string         `gorm:"type:varchar(200);column:seo_title" json:"seo_title"`                                  // SEO标题
	SeoKeywords    string         `gorm:"type:varchar(500);column:seo_keywords" json:"seo_keywords"`                            // SEO关键词
	SeoDescription string         `gorm:"type:varchar(1000);column:seo_description" json:"seo_description"`                     // SEO描述
	CreatedAt      time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at;index" json:"created_at"`    // 创建时间
	UpdatedAt      time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`          // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"type:datetime;column:deleted_at;index" json:"deleted_at,omitempty"`                    // 删除时间（软删除）

	// 关联关系
	Category ProductCategory    `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"` // 商品分类
	Brand    *ProductBrand      `gorm:"foreignKey:BrandID;references:ID" json:"brand,omitempty"`       // 商品品牌
	Skus     []ProductSku       `gorm:"foreignKey:ProductID;references:ID" json:"skus,omitempty"`      // 商品SKU列表
	Images   []ProductImage     `gorm:"foreignKey:ProductID;references:ID" json:"images,omitempty"`    // 商品图片列表
	Attrs    []ProductAttribute `gorm:"foreignKey:ProductID;references:ID" json:"attrs,omitempty"`     // 商品属性列表
}

func (p *Product) TableName() string {
	return "products"
}
