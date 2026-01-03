package service

import (
	"speed/app/http/model"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
)

type ProductCategoryService struct {
}

func NewProductCategoryService() *ProductCategoryService {
	return &ProductCategoryService{}
}

// CategoryListRequest 商品分类列表请求参数
type CategoryListRequest struct {
	ParentID *int  `form:"parent_id"` // 父分类ID（0或不传表示查询顶级分类）
	Status   *int8 `form:"status"`    // 状态（1-启用，0-禁用），不传则查询所有
}

// CategoryResponse 商品分类响应结构（格式化时间字段）
type CategoryResponse struct {
	ID          int                `json:"id"`                 // 主键
	Name        string             `json:"name"`               // 分类名称
	ParentID    int                `json:"parent_id"`          // 父分类ID（0表示顶级分类）
	SortOrder   int                `json:"sort_order"`         // 排序值（数字越小越靠前）
	Description string             `json:"description"`        // 分类描述
	ImageURL    string             `json:"image_url"`          // 分类图片URL
	Status      int8               `json:"status"`             // 状态（1-启用，0-禁用）
	CreatedAt   string             `json:"created_at"`         // 创建时间（格式化后）
	UpdatedAt   string             `json:"updated_at"`         // 更新时间（格式化后）
	Children    []CategoryResponse `json:"children,omitempty"` // 子分类列表（树形结构）
}

// toCategoryResponse 将 model.ProductCategory 转换为 CategoryResponse
func toCategoryResponse(category model.ProductCategory) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		ParentID:    category.ParentID,
		SortOrder:   category.SortOrder,
		Description: category.Description,
		ImageURL:    category.ImageURL,
		Status:      category.Status,
		CreatedAt:   formatTime(category.CreatedAt),
		UpdatedAt:   formatTime(category.UpdatedAt),
	}
}

// CategoryListResponse 商品分类列表响应
type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"` // 分类列表
}

// GetCategoryList 获取商品分类列表
// 如果 parent_id 为 0 或不传，返回树形结构的分类列表
// 如果 parent_id 有值，返回该父分类下的直接子分类列表
func (s *ProductCategoryService) GetCategoryList(ctx *gin.Context, req CategoryListRequest) (*CategoryListResponse, error) {
	var categories []model.ProductCategory
	query := app.Db.WithContext(ctx).Model(&model.ProductCategory{})

	// 如果指定了父分类ID，查询该父分类下的直接子分类
	if req.ParentID != nil {
		query = query.Where("parent_id = ?", *req.ParentID)
	} else {
		// 默认查询顶级分类（parent_id = 0）
		query = query.Where("parent_id = ?", 0)
	}

	// 状态筛选
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 按排序值排序
	query = query.Order("sort_order ASC, id ASC")

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	// 转换为响应格式
	categoryResponses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = toCategoryResponse(category)
	}

	// 如果查询的是顶级分类，构建树形结构
	if req.ParentID == nil || *req.ParentID == 0 {
		categoryResponses = s.buildCategoryTree(ctx, categoryResponses, req.Status)
	}

	return &CategoryListResponse{
		Categories: categoryResponses,
	}, nil
}

// buildCategoryTree 构建分类树形结构
func (s *ProductCategoryService) buildCategoryTree(ctx *gin.Context, topCategories []CategoryResponse, status *int8) []CategoryResponse {
	// 获取所有分类（根据状态筛选）
	var allCategories []model.ProductCategory
	query := app.Db.WithContext(ctx).Model(&model.ProductCategory{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query.Order("sort_order ASC, id ASC").Find(&allCategories)

	// 构建分类映射表（按父分类ID分组）
	categoryMap := make(map[int][]CategoryResponse)
	for _, cat := range allCategories {
		if cat.ParentID == 0 {
			continue // 跳过顶级分类，已经在 topCategories 中
		}
		catResp := toCategoryResponse(cat)
		categoryMap[cat.ParentID] = append(categoryMap[cat.ParentID], catResp)
	}

	// 递归构建树形结构
	var buildTree func([]CategoryResponse) []CategoryResponse
	buildTree = func(cats []CategoryResponse) []CategoryResponse {
		for i := range cats {
			if children, ok := categoryMap[cats[i].ID]; ok {
				cats[i].Children = buildTree(children)
			}
		}
		return cats
	}

	return buildTree(topCategories)
}

// GetCategoryTree 获取完整的分类树（所有分类，包括禁用的）
func (s *ProductCategoryService) GetCategoryTree(ctx *gin.Context) (*CategoryListResponse, error) {
	var allCategories []model.ProductCategory
	if err := app.Db.WithContext(ctx).
		Model(&model.ProductCategory{}).
		Order("sort_order ASC, id ASC").
		Find(&allCategories).Error; err != nil {
		return nil, err
	}

	// 转换为响应格式
	categoryMap := make(map[int]CategoryResponse)
	var topCategories []CategoryResponse

	for _, cat := range allCategories {
		catResp := toCategoryResponse(cat)
		categoryMap[cat.ID] = catResp

		if cat.ParentID == 0 {
			topCategories = append(topCategories, catResp)
		}
	}

	// 构建树形结构
	for i := range topCategories {
		s.addChildrenToCategory(&topCategories[i], categoryMap)
	}

	return &CategoryListResponse{
		Categories: topCategories,
	}, nil
}

// addChildrenToCategory 递归添加子分类
func (s *ProductCategoryService) addChildrenToCategory(category *CategoryResponse, categoryMap map[int]CategoryResponse) {
	for _, cat := range categoryMap {
		if cat.ParentID == category.ID {
			s.addChildrenToCategory(&cat, categoryMap)
			category.Children = append(category.Children, cat)
		}
	}
}
