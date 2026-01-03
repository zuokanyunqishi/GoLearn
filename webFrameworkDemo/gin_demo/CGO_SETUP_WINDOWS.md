# Windows 上安装 CGO 支持指南

## 方法一：使用 TDM-GCC（推荐，最简单）

### 1. 下载 TDM-GCC
- 访问：https://jmeubank.github.io/tdm-gcc/
- 下载最新版本的 TDM-GCC 64-bit 安装程序
- 推荐版本：TDM-GCC 10.3.0 或更高版本

### 2. 安装 TDM-GCC
- 运行安装程序
- 选择 "MinGW-w64/TDM64 (32-bit and 64-bit)" 选项
- 安装路径建议：`C:\TDM-GCC-64`（默认路径）
- 确保勾选 "Add to PATH" 选项

### 3. 验证安装
打开 PowerShell 或 CMD，运行：
```bash
gcc --version
```
应该显示 GCC 版本信息。

### 4. 设置环境变量（如果安装时未自动添加）
- 打开"系统属性" → "高级" → "环境变量"
- 在"系统变量"中找到 `Path`，添加：
  - `C:\TDM-GCC-64\bin`
- 重启终端或电脑使环境变量生效

### 5. 验证 CGO 支持
```bash
go env CGO_ENABLED
```
应该显示 `1`（如果显示 `0`，需要设置环境变量）

### 6. 编译测试
```bash
cd webFrameworkDemo/gin_demo
go build -o test.exe .
```

---

## 方法二：使用 MinGW-w64

### 1. 下载 MinGW-w64
- 访问：https://www.mingw-w64.org/downloads/
- 或使用 MSYS2：https://www.msys2.org/

### 2. 使用 MSYS2 安装（推荐）
```bash
# 安装 MSYS2 后，在 MSYS2 终端中运行：
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-toolchain
```

### 3. 添加到 PATH
将以下路径添加到系统 PATH：
- `C:\msys64\mingw64\bin`

---

## 方法三：使用 Chocolatey（如果已安装 Chocolatey）

```bash
choco install mingw
```

---

## 验证 CGO 是否正常工作

### 1. 检查 CGO 状态
```bash
go env CGO_ENABLED
```

### 2. 如果显示 0，手动启用
```bash
# PowerShell
$env:CGO_ENABLED="1"
go env -w CGO_ENABLED=1

# CMD
set CGO_ENABLED=1
go env -w CGO_ENABLED=1
```

### 3. 测试编译
```bash
cd webFrameworkDemo/gin_demo
go build -o test.exe .
```

---

## 常见问题

### 问题 1：`gcc: command not found`
**解决**：确保 GCC 已正确安装并添加到 PATH 环境变量中。

### 问题 2：`CGO_ENABLED=0` 但已安装 GCC
**解决**：
```bash
go env -w CGO_ENABLED=1
```

### 问题 3：编译时出现链接错误
**解决**：确保安装了完整的工具链，包括：
- gcc
- g++
- binutils

### 问题 4：32位 vs 64位
**解决**：确保安装的 GCC 版本与 Go 的架构匹配（通常是 64 位）。

---

## 注意事项

1. **推荐使用 TDM-GCC**：最简单，一键安装，自动配置 PATH
2. **重启终端**：安装后需要关闭并重新打开终端
3. **检查 PATH**：确保 GCC 的 bin 目录在 PATH 的最前面（优先级更高）
4. **Go 版本**：确保使用最新版本的 Go（1.16+）

---

## 如果不想使用 CGO

如果不想安装 CGO 支持，项目已经配置为使用 `modernc.org/sqlite`（纯 Go 实现），可以在编译时禁用 CGO：

```bash
# 编译时禁用 CGO（使用纯 Go 驱动）
set CGO_ENABLED=0
go build -o app.exe .
```

这样就不需要安装任何 C 编译器。

