-- 电商系统商品表 SQLite 建表语句

-- 商品分类表
-- 字段说明：id-主键, name-分类名称, parent_id-父分类ID（0表示顶级分类）, sort_order-排序值（数字越小越靠前）
-- description-分类描述, image_url-分类图片URL, status-状态（1-启用，0-禁用）
CREATE TABLE IF NOT EXISTS product_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    parent_id INTEGER DEFAULT 0,
    sort_order INTEGER DEFAULT 0,
    description TEXT,
    image_url VARCHAR(500),
    status TINYINT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 创建分类索引
CREATE INDEX idx_categories_parent_id ON product_categories(parent_id);
CREATE INDEX idx_categories_status ON product_categories(status);

-- 商品品牌表
-- 字段说明：id-主键, name-品牌名称, logo_url-品牌Logo URL, description-品牌描述
-- sort_order-排序值, status-状态（1-启用，0-禁用）
CREATE TABLE IF NOT EXISTS product_brands (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    logo_url VARCHAR(500),
    description TEXT,
    sort_order INTEGER DEFAULT 0,
    status TINYINT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 创建品牌索引
CREATE INDEX idx_brands_status ON product_brands(status);

-- 商品表（主表）
-- 字段说明：
-- spu_code-SPU编码（商品唯一标识）, name-商品名称, subtitle-商品副标题/简短描述
-- description-商品详细描述, category_id-商品分类ID, brand_id-品牌ID
-- main_image_url-主图URL, image_urls-商品图片URL列表（JSON格式存储）
-- original_price-商品原价, sale_price-商品售价, cost_price-成本价
-- stock_quantity-库存数量, sale_count-销售数量, view_count-浏览量
-- weight-商品重量（单位：kg）, volume-商品体积（单位：立方米）, unit-单位（件、台、个等）
-- status-商品状态（1-上架，0-下架，2-待审核，3-审核失败）
-- is_hot-是否热门（1-是，0-否）, is_new-是否新品（1-是，0-否）, is_recommend-是否推荐（1-是，0-否）
-- sort_order-排序值（数字越小越靠前）
-- seo_title-SEO标题, seo_keywords-SEO关键词, seo_description-SEO描述
-- deleted_at-删除时间（软删除）
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    spu_code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    subtitle VARCHAR(500),
    description TEXT,
    category_id INTEGER NOT NULL,
    brand_id INTEGER,
    main_image_url VARCHAR(500),
    image_urls TEXT,
    original_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    sale_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    cost_price DECIMAL(10, 2),
    stock_quantity INTEGER DEFAULT 0,
    sale_count INTEGER DEFAULT 0,
    view_count INTEGER DEFAULT 0,
    weight DECIMAL(10, 2),
    volume DECIMAL(10, 2),
    unit VARCHAR(20) DEFAULT '件',
    status TINYINT DEFAULT 1,
    is_hot TINYINT DEFAULT 0,
    is_new TINYINT DEFAULT 0,
    is_recommend TINYINT DEFAULT 0,
    sort_order INTEGER DEFAULT 0,
    seo_title VARCHAR(200),
    seo_keywords VARCHAR(500),
    seo_description VARCHAR(1000),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    FOREIGN KEY (category_id) REFERENCES product_categories(id),
    FOREIGN KEY (brand_id) REFERENCES product_brands(id)
);

-- 创建商品索引
CREATE INDEX idx_products_spu_code ON products(spu_code);
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_brand_id ON products(brand_id);
CREATE INDEX idx_products_status ON products(status);
CREATE INDEX idx_products_is_hot ON products(is_hot);
CREATE INDEX idx_products_is_new ON products(is_new);
CREATE INDEX idx_products_is_recommend ON products(is_recommend);
CREATE INDEX idx_products_created_at ON products(created_at);
CREATE INDEX idx_products_deleted_at ON products(deleted_at);

-- 商品SKU表（库存量单位，如：不同颜色、尺寸等）
-- 字段说明：product_id-商品ID, sku_code-SKU编码, sku_name-SKU名称（如：红色-M码）
-- spec_values-规格值（JSON格式存储，如：{"颜色":"红色","尺寸":"M"}）
-- original_price-SKU原价, sale_price-SKU售价, cost_price-SKU成本价
-- stock_quantity-SKU库存数量, image_url-SKU图片URL, status-状态（1-启用，0-禁用）
CREATE TABLE IF NOT EXISTS product_skus (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER NOT NULL,
    sku_code VARCHAR(100) UNIQUE NOT NULL,
    sku_name VARCHAR(200),
    spec_values TEXT,
    original_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    sale_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    cost_price DECIMAL(10, 2),
    stock_quantity INTEGER DEFAULT 0,
    image_url VARCHAR(500),
    status TINYINT DEFAULT 1,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- 创建SKU索引
CREATE INDEX idx_skus_product_id ON product_skus(product_id);
CREATE INDEX idx_skus_sku_code ON product_skus(sku_code);
CREATE INDEX idx_skus_status ON product_skus(status);

-- 商品规格表（用于定义商品的规格属性，如：颜色、尺寸等）
-- 字段说明：name-规格名称（如：颜色、尺寸）, sort_order-排序值, status-状态（1-启用，0-禁用）
CREATE TABLE IF NOT EXISTS product_specs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL,
    sort_order INTEGER DEFAULT 0,
    status TINYINT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 商品规格值表（规格的具体值，如：红色、蓝色、M码、L码等）
-- 字段说明：spec_id-规格ID, value-规格值, sort_order-排序值, status-状态（1-启用，0-禁用）
CREATE TABLE IF NOT EXISTS product_spec_values (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    spec_id INTEGER NOT NULL,
    value VARCHAR(100) NOT NULL,
    sort_order INTEGER DEFAULT 0,
    status TINYINT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (spec_id) REFERENCES product_specs(id) ON DELETE CASCADE
);

-- 创建规格值索引
CREATE INDEX idx_spec_values_spec_id ON product_spec_values(spec_id);

-- 商品属性表（商品的扩展属性，如：产地、材质等）
-- 字段说明：product_id-商品ID, attr_name-属性名称, attr_value-属性值, sort_order-排序值
CREATE TABLE IF NOT EXISTS product_attributes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER NOT NULL,
    attr_name VARCHAR(100) NOT NULL,
    attr_value TEXT NOT NULL,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- 创建属性索引
CREATE INDEX idx_attributes_product_id ON product_attributes(product_id);

-- 商品图片表（额外商品图片）
-- 字段说明：product_id-商品ID, image_url-图片URL, sort_order-排序值, is_main-是否主图（1-是，0-否）
CREATE TABLE IF NOT EXISTS product_images (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER NOT NULL,
    image_url VARCHAR(500) NOT NULL,
    sort_order INTEGER DEFAULT 0,
    is_main TINYINT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- 创建图片索引
CREATE INDEX idx_images_product_id ON product_images(product_id);
CREATE INDEX idx_images_is_main ON product_images(is_main);

