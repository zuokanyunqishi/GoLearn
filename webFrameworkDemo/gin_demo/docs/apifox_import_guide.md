# Apifox 导入接口文档指南

## 导入步骤

### 1. 打开 Apifox

启动 Apifox 应用程序。

### 2. 导入 OpenAPI 文档

1. 在 Apifox 中，点击 **项目设置** 或 **导入**
2. 选择 **导入** → **OpenAPI/Swagger**
3. 选择文件：`docs/cart_api_openapi.yaml`
4. 点击 **导入**

### 3. 配置环境变量

导入后，需要配置环境变量以便测试：

1. 进入 **环境管理**
2. 创建或编辑环境，设置以下变量：
   - `baseUrl`: `http://localhost:8086/api`
   - `token`: 登录后获取的 JWT Token（可选，可以在接口中手动设置）

### 4. 配置认证

1. 进入 **项目设置** → **认证**
2. 选择 **Bearer Token** 认证方式
3. 设置 Token 变量名：`token`
4. 或者在每个接口的 **认证** 标签页中手动设置 Bearer Token

### 5. 测试接口

1. 首先调用登录接口获取 Token：
   - `POST /api/login`
   - 请求体：
     ```json
     {
       "username": "your_username",
       "password": "your_password"
     }
     ```
   - 从响应中复制 `token` 字段的值

2. 在环境变量中设置 `token`，或者在每个接口的 Header 中设置：
   ```
   Authorization: Bearer {token}
   ```

3. 测试购物车接口：
   - 添加商品到购物车
   - 获取购物车列表
   - 更新购物车
   - 删除购物车记录
   - 等等

## 接口列表

导入后，您将看到以下接口：

1. **POST /api/cart** - 添加商品到购物车
2. **GET /api/cart** - 获取购物车列表
3. **GET /api/cart/count** - 获取购物车商品数量
4. **PUT /api/cart/{id}** - 更新购物车
5. **DELETE /api/cart/{id}** - 删除购物车记录
6. **PUT /api/cart/batch-selected** - 批量更新选中状态
7. **DELETE /api/cart/clear** - 清空购物车

## 注意事项

1. 所有购物车接口都需要 JWT 认证
2. 确保服务器运行在 `http://localhost:8086`
3. 测试前需要先登录获取 Token
4. 某些接口需要先有商品数据，请确保数据库中有测试商品

## 快速测试流程

1. 启动服务器：`go run main.go`
2. 在 Apifox 中导入 `cart_api_openapi.yaml`
3. 配置环境变量和认证
4. 调用登录接口获取 Token
5. 测试购物车相关接口


