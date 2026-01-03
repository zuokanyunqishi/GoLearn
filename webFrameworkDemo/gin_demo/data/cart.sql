-- 购物车表 SQLite 建表语句

-- 购物车表
-- 字段说明：
-- id-主键
-- user_id-用户ID（关联users表）
-- product_id-商品ID（关联products表）
-- sku_id-SKU ID（关联product_skus表，可选，如果商品没有SKU则为NULL）
-- quantity-商品数量
-- is_selected-是否选中（1-选中，0-未选中），用于结算时选择商品
-- created_at-创建时间
-- updated_at-更新时间
-- deleted_at-删除时间
CREATE TABLE IF NOT EXISTS carts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    sku_id INTEGER,
    quantity INTEGER NOT NULL DEFAULT 1,
    is_selected TINYINT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (sku_id) REFERENCES product_skus(id) ON DELETE CASCADE
);

-- 创建购物车索引
-- 用户ID索引，用于快速查询用户的购物车
CREATE INDEX idx_carts_user_id ON carts(user_id);
-- 商品ID索引，用于快速查询商品在购物车中的情况
CREATE INDEX idx_carts_product_id ON carts(product_id);
-- SKU ID索引，用于快速查询SKU在购物车中的情况
CREATE INDEX idx_carts_sku_id ON carts(sku_id);
-- 用户ID和商品ID组合索引，用于快速查询用户是否已添加该商品到购物车
CREATE INDEX idx_carts_user_product ON carts(user_id, product_id);
-- 用户ID、商品ID和SKU ID组合唯一索引，确保同一用户、同一商品、同一SKU只有一条记录
CREATE UNIQUE INDEX idx_carts_user_product_sku ON carts(user_id, product_id, sku_id);
-- 选中状态索引，用于快速查询选中的商品
CREATE INDEX idx_carts_is_selected ON carts(is_selected);

-- 说明：
-- 1. 购物车表支持商品和SKU两种粒度，如果商品没有SKU，则sku_id为NULL
-- 2. 通过唯一索引 idx_carts_user_product_sku 确保同一用户、同一商品、同一SKU只能有一条记录
--    - 如果商品没有SKU，则sku_id为NULL，此时同一用户、同一商品只能有一条记录
--    - 如果商品有SKU，则同一用户、同一商品、同一SKU只能有一条记录
-- 3. 当用户添加商品到购物车时，如果已存在相同记录，应该更新数量而不是新增记录
-- 4. is_selected字段用于结算时选择商品，默认值为1（选中）
-- 5. 外键约束使用ON DELETE CASCADE，当用户、商品或SKU被删除时，自动删除购物车中的相关记录

