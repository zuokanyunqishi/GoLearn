package service

import (
	"speed/app/http/model"
	app "speed/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var chinaLocation *time.Location

func init() {
	// 初始化中国时区
	var err error
	chinaLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 如果加载失败，使用 UTC+8
		chinaLocation = time.FixedZone("CST", 8*3600)
	}
}

// formatTime 格式化时间为中国时区的 Y-m-d H:i:s 格式
func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(chinaLocation).Format("2006-01-02 15:04:05")
}

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

// ProductListRequest 商品列表请求参数
type ProductListRequest struct {
	Page        int    `form:"page" binding:"omitempty,min=1"`              // 页码，默认1
	PageSize    int    `form:"page_size" binding:"omitempty,min=1,max=100"` // 每页数量，默认10，最大100
	CategoryID  *int   `form:"category_id"`                                 // 分类ID
	BrandID     *int   `form:"brand_id"`                                    // 品牌ID
	Status      *int8  `form:"status"`                                      // 状态（1-上架，0-下架，2-待审核，3-审核失败）
	IsHot       *int8  `form:"is_hot"`                                      // 是否热门（1-是，0-否）
	IsNew       *int8  `form:"is_new"`                                      // 是否新品（1-是，0-否）
	IsRecommend *int8  `form:"is_recommend"`                                // 是否推荐（1-是，0-否）
	Keyword     string `form:"keyword"`                                     // 关键词搜索（商品名称、SPU编码）
	SortBy      string `form:"sort_by"`                                     // 排序字段（created_at, sale_count, view_count, sale_price）
	SortOrder   string `form:"sort_order"`                                  // 排序方式（asc, desc），默认desc
}

// ProductResponse 商品响应结构（格式化时间字段）
type ProductResponse struct {
	ID             int      `json:"id"`              // 主键
	SpuCode        string   `json:"spu_code"`        // SPU编码（商品唯一标识）
	Name           string   `json:"name"`            // 商品名称
	Subtitle       string   `json:"subtitle"`        // 商品副标题/简短描述
	Description    string   `json:"description"`     // 商品详细描述
	CategoryID     int      `json:"category_id"`     // 商品分类ID
	BrandID        *int     `json:"brand_id"`        // 品牌ID
	MainImageURL   string   `json:"main_image_url"`  // 主图URL
	ImageURLs      string   `json:"image_urls"`      // 商品图片URL列表（JSON格式存储）
	OriginalPrice  float64  `json:"original_price"`  // 商品原价
	SalePrice      float64  `json:"sale_price"`      // 商品售价
	CostPrice      *float64 `json:"cost_price"`      // 成本价
	StockQuantity  int      `json:"stock_quantity"`  // 库存数量
	SaleCount      int      `json:"sale_count"`      // 销售数量
	ViewCount      int      `json:"view_count"`      // 浏览量
	Weight         *float64 `json:"weight"`          // 商品重量（单位：kg）
	Volume         *float64 `json:"volume"`          // 商品体积（单位：立方米）
	Unit           string   `json:"unit"`            // 单位（件、台、个等）
	Status         int8     `json:"status"`          // 商品状态（1-上架，0-下架，2-待审核，3-审核失败）
	IsHot          int8     `json:"is_hot"`          // 是否热门（1-是，0-否）
	IsNew          int8     `json:"is_new"`          // 是否新品（1-是，0-否）
	IsRecommend    int8     `json:"is_recommend"`    // 是否推荐（1-是，0-否）
	SortOrder      int      `json:"sort_order"`      // 排序值（数字越小越靠前）
	SeoTitle       string   `json:"seo_title"`       // SEO标题
	SeoKeywords    string   `json:"seo_keywords"`    // SEO关键词
	SeoDescription string   `json:"seo_description"` // SEO描述
	CreatedAt      string   `json:"created_at"`      // 创建时间（格式化后）
	UpdatedAt      string   `json:"updated_at"`      // 更新时间（格式化后）
}

// toProductResponse 将 model.Product 转换为 ProductResponse
func toProductResponse(p model.Product) ProductResponse {
	return ProductResponse{
		ID:             p.ID,
		SpuCode:        p.SpuCode,
		Name:           p.Name,
		Subtitle:       p.Subtitle,
		Description:    p.Description,
		CategoryID:     p.CategoryID,
		BrandID:        p.BrandID,
		MainImageURL:   p.MainImageURL,
		ImageURLs:      p.ImageURLs,
		OriginalPrice:  p.OriginalPrice,
		SalePrice:      p.SalePrice,
		CostPrice:      p.CostPrice,
		StockQuantity:  p.StockQuantity,
		SaleCount:      p.SaleCount,
		ViewCount:      p.ViewCount,
		Weight:         p.Weight,
		Volume:         p.Volume,
		Unit:           p.Unit,
		Status:         p.Status,
		IsHot:          p.IsHot,
		IsNew:          p.IsNew,
		IsRecommend:    p.IsRecommend,
		SortOrder:      p.SortOrder,
		SeoTitle:       p.SeoTitle,
		SeoKeywords:    p.SeoKeywords,
		SeoDescription: p.SeoDescription,
		CreatedAt:      formatTime(p.CreatedAt),
		UpdatedAt:      formatTime(p.UpdatedAt),
	}
}

// ProductListResponse 商品列表响应
type ProductListResponse struct {
	Total      int64             `json:"total"`       // 总记录数
	Page       int               `json:"page"`        // 当前页码
	PageSize   int               `json:"page_size"`   // 每页数量
	TotalPages int               `json:"total_pages"` // 总页数
	Products   []ProductResponse `json:"products"`    // 商品列表
}

// GetProductList 获取商品列表
func (s *ProductService) GetProductList(ctx *gin.Context, req ProductListRequest) (*ProductListResponse, error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.SortBy == "" {
		req.SortBy = "created_at"
	}
	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	// 构建查询
	query := app.Db.WithContext(ctx).Model(&model.Product{})

	// 筛选条件
	if req.CategoryID != nil {
		query = query.Where("category_id = ?", *req.CategoryID)
	}
	if req.BrandID != nil {
		query = query.Where("brand_id = ?", *req.BrandID)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	} else {
		// 默认只查询上架商品
		query = query.Where("status = ?", 1)
	}
	if req.IsHot != nil {
		query = query.Where("is_hot = ?", *req.IsHot)
	}
	if req.IsNew != nil {
		query = query.Where("is_new = ?", *req.IsNew)
	}
	if req.IsRecommend != nil {
		query = query.Where("is_recommend = ?", *req.IsRecommend)
	}
	if req.Keyword != "" {
		query = query.Where("name LIKE ? OR spu_code LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 排序
	orderBy := req.SortBy
	if req.SortOrder == "asc" {
		orderBy += " ASC"
	} else {
		orderBy += " DESC"
	}
	query = query.Order(orderBy)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var products []model.Product
	if err := query.Offset(offset).Limit(req.PageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := int(total) / req.PageSize
	if int(total)%req.PageSize > 0 {
		totalPages++
	}

	// 转换为响应格式（格式化时间）
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = toProductResponse(product)
	}

	return &ProductListResponse{
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
		Products:   productResponses,
	}, nil
}

// ProductDetailResponse 商品详情响应结构（格式化时间字段）
type ProductDetailResponse struct {
	ProductResponse
	Category *model.ProductCategory   `json:"category,omitempty"` // 商品分类
	Brand    *model.ProductBrand      `json:"brand,omitempty"`    // 商品品牌
	Skus     []ProductSkuResponse     `json:"skus,omitempty"`     // 商品SKU列表
	Images   []ProductImageResponse   `json:"images,omitempty"`   // 商品图片列表
	Attrs    []model.ProductAttribute `json:"attrs,omitempty"`    // 商品属性列表
}

// ProductSkuResponse SKU响应结构（格式化时间字段）
type ProductSkuResponse struct {
	ID            int      `json:"id"`             // 主键
	ProductID     int      `json:"product_id"`     // 商品ID
	SkuCode       string   `json:"sku_code"`       // SKU编码
	SkuName       string   `json:"sku_name"`       // SKU名称
	SpecValues    string   `json:"spec_values"`    // 规格值
	OriginalPrice float64  `json:"original_price"` // SKU原价
	SalePrice     float64  `json:"sale_price"`     // SKU售价
	CostPrice     *float64 `json:"cost_price"`     // SKU成本价
	StockQuantity int      `json:"stock_quantity"` // SKU库存数量
	ImageURL      string   `json:"image_url"`      // SKU图片URL
	Status        int8     `json:"status"`         // 状态
	SortOrder     int      `json:"sort_order"`     // 排序值
	CreatedAt     string   `json:"created_at"`     // 创建时间（格式化后）
	UpdatedAt     string   `json:"updated_at"`     // 更新时间（格式化后）
}

// ProductImageResponse 商品图片响应结构（格式化时间字段）
type ProductImageResponse struct {
	ID        int    `json:"id"`         // 主键
	ProductID int    `json:"product_id"` // 商品ID
	ImageURL  string `json:"image_url"`  // 图片URL
	SortOrder int    `json:"sort_order"` // 排序值
	IsMain    int8   `json:"is_main"`    // 是否主图
	CreatedAt string `json:"created_at"` // 创建时间（格式化后）
}

// toProductSkuResponse 将 model.ProductSku 转换为 ProductSkuResponse
func toProductSkuResponse(sku model.ProductSku) ProductSkuResponse {
	return ProductSkuResponse{
		ID:            sku.ID,
		ProductID:     sku.ProductID,
		SkuCode:       sku.SkuCode,
		SkuName:       sku.SkuName,
		SpecValues:    sku.SpecValues,
		OriginalPrice: sku.OriginalPrice,
		SalePrice:     sku.SalePrice,
		CostPrice:     sku.CostPrice,
		StockQuantity: sku.StockQuantity,
		ImageURL:      sku.ImageURL,
		Status:        sku.Status,
		SortOrder:     sku.SortOrder,
		CreatedAt:     formatTime(sku.CreatedAt),
		UpdatedAt:     formatTime(sku.UpdatedAt),
	}
}

// toProductImageResponse 将 model.ProductImage 转换为 ProductImageResponse
func toProductImageResponse(img model.ProductImage) ProductImageResponse {
	return ProductImageResponse{
		ID:        img.ID,
		ProductID: img.ProductID,
		ImageURL:  img.ImageURL,
		SortOrder: img.SortOrder,
		IsMain:    img.IsMain,
		CreatedAt: formatTime(img.CreatedAt),
	}
}

// GetProductDetail 获取商品详情
func (s *ProductService) GetProductDetail(ctx *gin.Context, id int) (*ProductDetailResponse, error) {
	var product model.Product
	err := app.Db.WithContext(ctx).
		Preload("Category").
		Preload("Brand").
		Preload("Skus").
		Preload("Images").
		Preload("Attrs").
		Where("id = ? AND status = ?", id, 1).
		First(&product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// 增加浏览量
	app.Db.WithContext(ctx).Model(&product).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1))

	// 转换为响应格式（格式化时间）
	response := &ProductDetailResponse{
		ProductResponse: toProductResponse(product),
		Category:        &product.Category,
		Brand:           product.Brand,
		Attrs:           product.Attrs,
	}

	// 转换SKU列表
	if len(product.Skus) > 0 {
		response.Skus = make([]ProductSkuResponse, len(product.Skus))
		for i, sku := range product.Skus {
			response.Skus[i] = toProductSkuResponse(sku)
		}
	}

	// 转换图片列表
	if len(product.Images) > 0 {
		response.Images = make([]ProductImageResponse, len(product.Images))
		for i, img := range product.Images {
			response.Images[i] = toProductImageResponse(img)
		}
	}

	return response, nil
}
