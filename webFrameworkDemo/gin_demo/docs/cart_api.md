# 购物车API接口文档

## 概述

购物车功能提供了一套完整的API接口，支持用户添加、更新、删除、查询购物车商品等操作。所有接口都需要用户登录认证（Bearer Token）。

## 基础信息

- **Base URL**: `/api/cart`
- **认证方式**: Bearer Token（JWT）
- **Content-Type**: `application/json`

## API接口列表

### 1. 添加商品到购物车

**接口地址**: `POST /api/cart`

**请求头**:
```
Authorization: Bearer {token}
Content-Type: application/json
```

**请求体**:
```json
{
  "product_id": 1,      // 商品ID（必填）
  "sku_id": 2,         // SKU ID（可选，如果商品没有SKU则不传或传null）
  "quantity": 1        // 商品数量（必填，范围：1-999）
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "添加成功"
  }
}
```

**说明**:
- 如果购物车中已存在相同的商品和SKU，会自动累加数量
- 添加前会检查商品和SKU的库存是否充足
- 添加前会验证商品和SKU是否存在且状态正常

---

### 2. 获取购物车列表

**接口地址**: `GET /api/cart`

**请求头**:
```
Authorization: Bearer {token}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "product_id": 1,
      "sku_id": 2,
      "quantity": 2,
      "is_selected": 1,
      "created_at": "2024-01-01 12:00:00",
      "updated_at": "2024-01-01 12:00:00",
      "product": {
        "id": 1,
        "spu_code": "SPU001",
        "name": "商品名称",
        "sale_price": 99.00,
        // ... 其他商品信息
      },
      "sku": {
        "id": 2,
        "sku_code": "SKU001",
        "sku_name": "红色-M码",
        "sale_price": 99.00,
        // ... 其他SKU信息
      }
    }
  ]
}
```

**说明**:
- 返回当前用户的所有购物车记录
- 包含商品和SKU的完整信息
- 按创建时间倒序排列

---

### 3. 获取购物车商品数量

**接口地址**: `GET /api/cart/count`

**请求头**:
```
Authorization: Bearer {token}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "count": 5
  }
}
```

**说明**:
- 返回当前用户购物车中的商品总数量（按记录数，不是按商品件数）

---

### 4. 更新购物车

**接口地址**: `POST /api/cart/update/{id}`

**请求头**:
```
Authorization: Bearer {token}
Content-Type: application/json
```

**路径参数**:
- `id`: 购物车记录ID

**请求体**:
```json
{
  "quantity": 3,       // 商品数量（可选，范围：1-999）
  "is_selected": 1     // 是否选中（可选，1-选中，0-未选中）
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "更新成功"
  }
}
```

**说明**:
- 可以只更新数量，或只更新选中状态，或同时更新
- 更新数量前会检查库存是否充足
- 只能更新当前用户的购物车记录

---

### 5. 删除购物车记录

**接口地址**: `POST /api/cart/del/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
- `id`: 购物车记录ID

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "删除成功"
  }
}
```

**说明**:
- 使用软删除，记录不会真正从数据库中删除
- 只能删除当前用户的购物车记录

---

### 6. 批量更新选中状态

**接口地址**: `POST /api/cart/batch-selected`

**请求头**:
```
Authorization: Bearer {token}
Content-Type: application/json
```

**请求体**:
```json
{
  "cart_ids": [1, 2, 3],  // 购物车ID列表（必填，至少一个）
  "is_selected": 1         // 是否选中（必填，1-选中，0-未选中）
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "更新成功"
  }
}
```

**说明**:
- 批量更新多个购物车记录的选中状态
- 常用于结算时选择/取消选择商品
- 只能更新当前用户的购物车记录

---

### 7. 清空购物车

**接口地址**: `POST /api/cart/clear`

**请求头**:
```
Authorization: Bearer {token}
```

**查询参数**:
- `only_unselected`: 是否只删除未选中的商品（可选，默认false，即删除所有）

**示例**:
```
POST /api/cart/clear?only_unselected=true
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "清空成功"
  }
}
```

**说明**:
- 如果不传`only_unselected`或传`false`，会删除所有购物车记录
- 如果传`true`，只删除未选中的购物车记录
- 使用软删除

---

## 错误响应格式

所有接口在发生错误时，都会返回统一的错误格式：

```json
{
  "code": 400,
  "message": "错误信息",
  "data": {}
}
```

**常见错误码**:
- `200`: 成功
- `400`: 请求参数错误
- `401`: 未授权（未登录或Token无效）
- `404`: 资源不存在
- `500`: 服务器内部错误

**常见错误信息**:
- "用户未登录"
- "商品不存在或已下架"
- "SKU不存在或已禁用"
- "商品库存不足"
- "SKU库存不足"
- "购物车记录不存在"
- "无权操作此购物车记录"

---

## 使用示例

### 示例1：添加商品到购物车

```bash
curl -X POST http://localhost:8080/api/cart \
  -H "Authorization: Bearer your_token_here" \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "sku_id": 2,
    "quantity": 1
  }'
```

### 示例2：获取购物车列表

```bash
curl -X GET http://localhost:8080/api/cart \
  -H "Authorization: Bearer your_token_here"
```

### 示例3：更新购物车商品数量

```bash
curl -X POST http://localhost:8080/api/cart/1 \
  -H "Authorization: Bearer your_token_here" \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 3
  }'
```

### 示例4：批量选中商品

```bash
curl -X POST http://localhost:8080/api/cart/batch-selected \
  -H "Authorization: Bearer your_token_here" \
  -H "Content-Type: application/json" \
  -d '{
    "cart_ids": [1, 2, 3],
    "is_selected": 1
  }'
```

---

## 注意事项

1. **库存检查**: 添加和更新购物车时，系统会自动检查商品和SKU的库存，确保不会超卖
2. **唯一性**: 同一用户、同一商品、同一SKU在购物车中只能有一条记录，重复添加会自动累加数量
3. **权限控制**: 用户只能操作自己的购物车记录，无法操作其他用户的购物车
4. **软删除**: 删除操作使用软删除，数据不会真正从数据库中删除
5. **选中状态**: 默认添加的商品都是选中状态（`is_selected=1`），可用于结算时选择商品


