package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"speed/app/http/model"
	"speed/app/lib/log"
	"speed/app/lib/validate"
	app "speed/bootstrap"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // 使用纯 Go 实现的 SQLite 驱动，不需要 CGO
	"moul.io/zapgorm2"
)

func init() {
	// 从文件系统读取配置文件
	workDir, _ := os.Getwd()
	// 如果当前在 cmd 目录，需要回到项目根目录
	if filepath.Base(workDir) == "cmd" {
		workDir = filepath.Dir(workDir)
	}

	configPath := filepath.Join(workDir, ".config.json")
	dbPath := filepath.Join(workDir, "data", "shop.sqlite3")

	// 初始化配置
	app.Config = viper.New()
	app.Config.SetConfigType("json")
	app.Config.SetConfigFile(configPath)
	if err := app.Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("无法读取配置文件: %w", err))
	}

	// 初始化应用名称等
	app.AppName = app.Config.GetString("appName")
	app.AppEnv = app.Config.GetString("appEnv")
	app.AppKey = app.Config.GetString("appKey")

	// 初始化日志
	app.Log0 = initLog()
	app.Log = app.Log0.Sugar()

	// 初始化验证器
	validate.Init()

	// 初始化数据库
	initSqliteDbFromFile(dbPath)
}

func initLog() *zap.Logger {
	fileName := "storage/logs/zap.log"
	level := zapcore.DebugLevel

	var write zapcore.WriteSyncer
	if app.AppEnv == "prod" {
		write = zapcore.AddSync(&lumberjack.Logger{
			Filename:  fileName,
			MaxSize:   1 << 30, //1G
			LocalTime: true,
			Compress:  true,
		})
	} else {
		write = os.Stdout
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(i.Format("2006-01-02 15:04:05.000"))
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), write, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(log.Stack{}), zap.Development())
	log.Log = logger.Sugar()
	return logger
}

func initSqliteDbFromFile(dbPath string) {
	// 使用直接的文件路径连接数据库
	Logger := zapgorm2.New(app.Log0)
	Logger.Context = func(ctx context.Context) []zapcore.Field {
		traceId := ctx.Value("traceId")
		if traceId != nil {
			return []zapcore.Field{{Key: "traceId", Type: zapcore.StringType, String: traceId.(string)}}
		}
		return []zapcore.Field{}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: Logger.LogMode(gormlogger.Info)})
	if err != nil {
		panic(fmt.Errorf("无法连接数据库: %w", err))
	}
	app.Db = db
}

// 商品分类数据
var categories = []struct {
	Name        string
	Description string
	Children    []struct {
		Name        string
		Description string
	}
}{
	{
		Name:        "电子产品",
		Description: "各类电子数码产品",
		Children: []struct {
			Name        string
			Description string
		}{
			{Name: "手机", Description: "智能手机、功能手机"},
			{Name: "电脑", Description: "笔记本电脑、台式机"},
			{Name: "平板", Description: "平板电脑"},
			{Name: "耳机", Description: "有线耳机、无线耳机"},
		},
	},
	{
		Name:        "服装鞋帽",
		Description: "时尚服装和鞋类",
		Children: []struct {
			Name        string
			Description string
		}{
			{Name: "男装", Description: "男士服装"},
			{Name: "女装", Description: "女士服装"},
			{Name: "运动鞋", Description: "运动休闲鞋"},
			{Name: "皮鞋", Description: "商务正装皮鞋"},
		},
	},
	{
		Name:        "食品饮料",
		Description: "各类食品和饮品",
		Children: []struct {
			Name        string
			Description string
		}{
			{Name: "零食", Description: "休闲零食"},
			{Name: "饮料", Description: "各类饮品"},
			{Name: "生鲜", Description: "新鲜食材"},
			{Name: "粮油", Description: "米面粮油"},
		},
	},
	{
		Name:        "家居用品",
		Description: "家居生活用品",
		Children: []struct {
			Name        string
			Description string
		}{
			{Name: "家具", Description: "各类家具"},
			{Name: "家纺", Description: "床上用品"},
			{Name: "厨具", Description: "厨房用品"},
			{Name: "收纳", Description: "收纳整理用品"},
		},
	},
	{
		Name:        "美妆护肤",
		Description: "化妆品和护肤品",
		Children: []struct {
			Name        string
			Description string
		}{
			{Name: "护肤", Description: "面部护理"},
			{Name: "彩妆", Description: "彩妆产品"},
			{Name: "香水", Description: "各类香水"},
			{Name: "男士护理", Description: "男士护肤"},
		},
	},
}

// 品牌数据
var brands = []struct {
	Name        string
	Description string
}{
	{Name: "苹果", Description: "Apple Inc. 全球知名科技品牌"},
	{Name: "华为", Description: "华为技术有限公司"},
	{Name: "小米", Description: "小米科技有限责任公司"},
	{Name: "三星", Description: "Samsung 韩国电子品牌"},
	{Name: "耐克", Description: "Nike 全球运动品牌"},
	{Name: "阿迪达斯", Description: "Adidas 德国运动品牌"},
	{Name: "优衣库", Description: "UNIQLO 日本快时尚品牌"},
	{Name: "可口可乐", Description: "Coca-Cola 全球饮料品牌"},
	{Name: "美的", Description: "美的集团 家电品牌"},
	{Name: "欧莱雅", Description: "L'Oreal 法国美妆品牌"},
}

// 商品名称模板
var productNames = []struct {
	Category string
	Names    []string
}{
	{
		Category: "手机",
		Names:    []string{"智能手机", "5G手机", "拍照手机", "游戏手机", "商务手机"},
	},
	{
		Category: "电脑",
		Names:    []string{"笔记本电脑", "游戏本", "商务本", "轻薄本", "工作站"},
	},
	{
		Category: "男装",
		Names:    []string{"T恤", "衬衫", "牛仔裤", "休闲裤", "外套", "夹克"},
	},
	{
		Category: "女装",
		Names:    []string{"连衣裙", "T恤", "牛仔裤", "半身裙", "外套", "毛衣"},
	},
	{
		Category: "零食",
		Names:    []string{"薯片", "坚果", "巧克力", "饼干", "糖果", "果干"},
	},
	{
		Category: "饮料",
		Names:    []string{"碳酸饮料", "果汁", "茶饮料", "功能饮料", "矿泉水"},
	},
}

// 规格名称
var specNames = []string{"颜色", "尺寸", "容量", "规格", "版本"}

// 规格值
var specValues = map[string][]string{
	"颜色": {"红色", "蓝色", "绿色", "黑色", "白色", "灰色", "粉色", "紫色"},
	"尺寸": {"S", "M", "L", "XL", "XXL"},
	"容量": {"64GB", "128GB", "256GB", "512GB", "1TB"},
	"规格": {"标准版", "豪华版", "旗舰版", "青春版"},
	"版本": {"WiFi版", "4G版", "5G版"},
}

func main() {
	// 初始化数据库连接
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("开始生成商品数据...")

	// 1. 生成商品分类
	categoryMap := seedCategories(ctx)
	fmt.Printf("✓ 已生成 %d 个分类\n", len(categoryMap))

	// 2. 生成品牌
	brandMap := seedBrands(ctx)
	fmt.Printf("✓ 已生成 %d 个品牌\n", len(brandMap))

	// 3. 生成规格和规格值
	specMap := seedSpecs(ctx)
	fmt.Printf("✓ 已生成 %d 个规格\n", len(specMap))

	// 4. 生成商品
	seedProducts(ctx, categoryMap, brandMap, specMap)
	fmt.Printf("✓ 商品数据生成完成\n")

	fmt.Println("\n所有数据生成完成！")
}

// 生成商品分类
func seedCategories(ctx context.Context) map[int]int {
	categoryMap := make(map[int]int) // 子分类ID -> 父分类ID

	for _, cat := range categories {
		// 创建父分类
		parent := model.ProductCategory{
			Name:        cat.Name,
			Description: cat.Description,
			ParentID:    0,
			SortOrder:   rand.Intn(100),
			Status:      1,
		}
		app.Db.WithContext(ctx).Create(&parent)

		// 创建子分类
		for i, child := range cat.Children {
			childCat := model.ProductCategory{
				Name:        child.Name,
				Description: child.Description,
				ParentID:    parent.ID,
				SortOrder:   i,
				Status:      1,
			}
			app.Db.WithContext(ctx).Create(&childCat)
			categoryMap[childCat.ID] = parent.ID
		}
	}

	return categoryMap
}

// 生成品牌
func seedBrands(ctx context.Context) map[int]*model.ProductBrand {
	brandMap := make(map[int]*model.ProductBrand)

	for i, brand := range brands {
		b := &model.ProductBrand{
			Name:        brand.Name,
			Description: brand.Description,
			SortOrder:   i,
			Status:      1,
		}
		app.Db.WithContext(ctx).Create(b)
		brandMap[b.ID] = b
	}

	return brandMap
}

// 生成规格和规格值
func seedSpecs(ctx context.Context) map[string]*model.ProductSpec {
	specMap := make(map[string]*model.ProductSpec)

	for _, specName := range specNames {
		spec := &model.ProductSpec{
			Name:      specName,
			SortOrder: rand.Intn(100),
			Status:    1,
		}
		app.Db.WithContext(ctx).Create(spec)
		specMap[specName] = spec

		// 生成规格值
		if values, ok := specValues[specName]; ok {
			for i, value := range values {
				specValue := model.ProductSpecValue{
					SpecID:    spec.ID,
					Value:     value,
					SortOrder: i,
					Status:    1,
				}
				app.Db.WithContext(ctx).Create(&specValue)
			}
		}
	}

	return specMap
}

// 生成商品
func seedProducts(ctx context.Context, categoryMap map[int]int, brandMap map[int]*model.ProductBrand, specMap map[string]*model.ProductSpec) {
	// 获取所有子分类
	var subCategories []model.ProductCategory
	app.Db.WithContext(ctx).Where("parent_id > 0").Find(&subCategories)

	if len(subCategories) == 0 {
		fmt.Println("警告: 没有找到子分类，无法生成商品")
		return
	}

	// 为每个子分类生成3-8个商品
	for _, subCat := range subCategories {
		productCount := 3 + rand.Intn(6) // 3-8个商品

		for i := 0; i < productCount; i++ {
			product := generateProduct(ctx, subCat.ID, brandMap, specMap)
			if product != nil {
				// 生成SKU
				generateSkus(ctx, product.ID, specMap)

				// 生成图片
				generateImages(ctx, product.ID)

				// 生成属性
				generateAttributes(ctx, product.ID)
			}
		}
	}
}

// 生成单个商品
func generateProduct(ctx context.Context, categoryID int, brandMap map[int]*model.ProductBrand, specMap map[string]*model.ProductSpec) *model.Product {
	// 获取分类信息
	var category model.ProductCategory
	app.Db.WithContext(ctx).First(&category, categoryID)

	// 随机选择品牌（70%概率有品牌）
	var brandID *int
	if rand.Float32() < 0.7 {
		// 随机选择一个品牌
		var brands []model.ProductBrand
		app.Db.WithContext(ctx).Where("status = ?", 1).Find(&brands)
		if len(brands) > 0 {
			bID := brands[rand.Intn(len(brands))].ID
			brandID = &bID
		}
	}

	// 生成商品名称
	productName := generateProductName(category.Name)
	subtitle := generateSubtitle()

	// 生成价格
	originalPrice := float64(50+rand.Intn(5000)) + rand.Float64()*100
	salePrice := originalPrice * (0.7 + rand.Float64()*0.2) // 7-9折
	costPrice := salePrice * (0.5 + rand.Float64()*0.3)     // 成本价

	// 生成SPU编码
	spuCode := fmt.Sprintf("SPU%06d", rand.Intn(999999))

	product := &model.Product{
		SpuCode:        spuCode,
		Name:           productName,
		Subtitle:       subtitle,
		Description:    generateDescription(productName),
		CategoryID:     categoryID,
		BrandID:        brandID,
		MainImageURL:   generateImageURL(),
		ImageURLs:      generateImageURLs(),
		OriginalPrice:  originalPrice,
		SalePrice:      salePrice,
		CostPrice:      &costPrice,
		StockQuantity:  rand.Intn(1000) + 10,
		SaleCount:      rand.Intn(5000),
		ViewCount:      rand.Intn(10000),
		Weight:         floatPtr(float64(rand.Intn(5000)+100) / 1000.0), // 0.1-5.1kg
		Volume:         floatPtr(float64(rand.Intn(100)+1) / 1000.0),    // 0.001-0.1立方米
		Unit:           "件",
		Status:         1,
		IsHot:          int8(rand.Intn(2)),
		IsNew:          int8(rand.Intn(2)),
		IsRecommend:    int8(rand.Intn(2)),
		SortOrder:      rand.Intn(100),
		SeoTitle:       productName + " - 优质商品",
		SeoKeywords:    productName + "," + category.Name,
		SeoDescription: subtitle,
	}

	app.Db.WithContext(ctx).Create(product)
	return product
}

// 生成SKU
func generateSkus(ctx context.Context, productID int, specMap map[string]*model.ProductSpec) {
	// 随机生成1-5个SKU
	skuCount := 1 + rand.Intn(5)

	for i := 0; i < skuCount; i++ {
		// 随机选择规格组合
		specValuesJSON := generateSpecValuesJSON(ctx, specMap)

		skuCode := fmt.Sprintf("SKU%06d", rand.Intn(999999))
		skuName := generateSkuName(specValuesJSON)

		// 获取商品价格
		var product model.Product
		app.Db.WithContext(ctx).First(&product, productID)

		// SKU价格在商品价格基础上浮动±10%
		priceVariation := 0.9 + rand.Float64()*0.2
		originalPrice := product.OriginalPrice * priceVariation
		salePrice := product.SalePrice * priceVariation
		costPrice := product.CostPrice
		if costPrice != nil {
			cp := *costPrice * priceVariation
			costPrice = &cp
		}

		sku := model.ProductSku{
			ProductID:     productID,
			SkuCode:       skuCode,
			SkuName:       skuName,
			SpecValues:    specValuesJSON,
			OriginalPrice: originalPrice,
			SalePrice:     salePrice,
			CostPrice:     costPrice,
			StockQuantity: rand.Intn(500) + 5,
			ImageURL:      generateImageURL(),
			Status:        1,
			SortOrder:     i,
		}

		app.Db.WithContext(ctx).Create(&sku)
	}
}

// 生成商品图片
func generateImages(ctx context.Context, productID int) {
	// 生成3-8张图片
	imageCount := 3 + rand.Intn(6)

	for i := 0; i < imageCount; i++ {
		image := model.ProductImage{
			ProductID: productID,
			ImageURL:  generateImageURL(),
			SortOrder: i,
			IsMain: int8(func() int {
				if i == 0 {
					return 1
				} else {
					return 0
				}
			}()),
		}
		app.Db.WithContext(ctx).Create(&image)
	}
}

// 生成商品属性
func generateAttributes(ctx context.Context, productID int) {
	attributes := []struct {
		Name  string
		Value string
	}{
		{"产地", []string{"中国", "美国", "日本", "德国", "韩国"}[rand.Intn(5)]},
		{"保质期", fmt.Sprintf("%d天", 30+rand.Intn(365))},
		{"材质", []string{"纯棉", "涤纶", "真皮", "金属", "塑料", "玻璃"}[rand.Intn(6)]},
		{"适用人群", []string{"通用", "儿童", "成人", "老人"}[rand.Intn(4)]},
		{"包装规格", fmt.Sprintf("%d件/箱", 10+rand.Intn(50))},
	}

	for i, attr := range attributes {
		productAttr := model.ProductAttribute{
			ProductID: productID,
			AttrName:  attr.Name,
			AttrValue: attr.Value,
			SortOrder: i,
		}
		app.Db.WithContext(ctx).Create(&productAttr)
	}
}

// 辅助函数
func generateProductName(categoryName string) string {
	for _, pn := range productNames {
		if pn.Category == categoryName {
			if len(pn.Names) > 0 {
				return pn.Names[rand.Intn(len(pn.Names))] + " " + categoryName
			}
		}
	}
	return "优质" + categoryName
}

func generateSubtitle() string {
	subtitles := []string{
		"品质保证，值得信赖",
		"热销爆款，限时优惠",
		"新品上市，抢先体验",
		"经典款式，永不过时",
		"精选材质，舒适体验",
		"时尚设计，个性选择",
	}
	return subtitles[rand.Intn(len(subtitles))]
}

func generateDescription(name string) string {
	return fmt.Sprintf("这是一款优质的%s，采用精选材料制作，工艺精湛，品质保证。适合日常使用，是您理想的选择。", name)
}

func generateImageURL() string {
	// 使用占位图片服务
	return fmt.Sprintf("https://picsum.photos/800/800?random=%d", rand.Intn(10000))
}

func generateImageURLs() string {
	urls := []string{}
	for i := 0; i < 3+rand.Intn(5); i++ {
		urls = append(urls, generateImageURL())
	}
	jsonBytes, _ := json.Marshal(urls)
	return string(jsonBytes)
}

func generateSpecValuesJSON(ctx context.Context, specMap map[string]*model.ProductSpec) string {
	// 随机选择1-3个规格
	selectedSpecs := []string{}
	specKeys := []string{}
	for k := range specMap {
		specKeys = append(specKeys, k)
	}

	count := 1 + rand.Intn(3)
	if count > len(specKeys) {
		count = len(specKeys)
	}

	// 随机选择规格
	used := make(map[string]bool)
	for len(selectedSpecs) < count {
		key := specKeys[rand.Intn(len(specKeys))]
		if !used[key] {
			selectedSpecs = append(selectedSpecs, key)
			used[key] = true
		}
	}

	// 构建JSON
	specValues := make(map[string]string)
	for _, specName := range selectedSpecs {
		if spec, ok := specMap[specName]; ok {
			// 获取该规格的所有值
			var values []model.ProductSpecValue
			app.Db.WithContext(ctx).Where("spec_id = ?", spec.ID).Find(&values)
			if len(values) > 0 {
				specValues[specName] = values[rand.Intn(len(values))].Value
			}
		}
	}

	jsonBytes, _ := json.Marshal(specValues)
	return string(jsonBytes)
}

func generateSkuName(specValuesJSON string) string {
	var specValues map[string]string
	json.Unmarshal([]byte(specValuesJSON), &specValues)

	parts := []string{}
	for _, v := range specValues {
		parts = append(parts, v)
	}

	if len(parts) > 0 {
		return parts[0]
	}
	return "标准版"
}

func floatPtr(f float64) *float64 {
	return &f
}
