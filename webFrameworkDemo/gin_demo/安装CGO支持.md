# Windows 安装 CGO 支持 - 详细步骤

## 📋 前置要求

- Windows 7 或更高版本
- 已安装 Go（1.16 或更高版本）
- 管理员权限（用于安装软件）

---

## 🚀 方法一：TDM-GCC（推荐，最简单）

### 步骤 1：下载 TDM-GCC

1. 打开浏览器，访问：**https://jmeubank.github.io/tdm-gcc/**
2. 点击 **"Download"** 按钮
3. 选择 **"tdm64-gcc-10.3.0-2.exe"** 或最新版本（64位）
4. 下载到本地（建议下载到 `Downloads` 文件夹）

### 步骤 2：安装 TDM-GCC

1. **右键点击下载的安装程序**，选择 **"以管理员身份运行"**
2. 在安装向导中：
   - 点击 **"Next"**
   - 选择安装路径（默认：`C:\TDM-GCC-64`，建议保持默认）
   - 点击 **"Next"**
   - **重要**：勾选 **"Add to PATH"** 选项
   - 点击 **"Install"**
   - 等待安装完成（约 2-5 分钟）
   - 点击 **"Finish"**

### 步骤 3：验证安装

1. **关闭所有已打开的终端窗口**（PowerShell、CMD、VS Code 终端等）
2. **重新打开** PowerShell 或 CMD
3. 运行以下命令：

```bash
gcc --version
```

**预期输出**：
```
gcc (tdm64-1) 10.3.0
Copyright (C) 2020 Free Software Foundation, Inc.
...
```

如果看到版本信息，说明安装成功！✅

### 步骤 4：启用 CGO

在终端中运行：

```bash
go env -w CGO_ENABLED=1
```

验证是否启用：

```bash
go env CGO_ENABLED
```

应该显示：`CGO_ENABLED=1`

### 步骤 5：测试编译

进入项目目录并编译：

```bash
cd D:\project\go\GoLearn\webFrameworkDemo\gin_demo
go build -o test.exe .
```

如果编译成功，说明 CGO 已正确配置！✅

---

## 🔧 方法二：MSYS2 + MinGW-w64（适合高级用户）

### 步骤 1：下载 MSYS2

1. 访问：**https://www.msys2.org/**
2. 下载 **"msys2-x86_64-latest.exe"**
3. 运行安装程序，安装到默认路径：`C:\msys64`

### 步骤 2：安装 MinGW-w64

1. 打开 **MSYS2 MSYS** 终端（不是 MinGW 终端）
2. 更新包数据库：

```bash
pacman -Syu
```

3. 安装 MinGW-w64 工具链：

```bash
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-toolchain
```

### 步骤 3：添加到系统 PATH

1. 按 `Win + R`，输入 `sysdm.cpl`，回车
2. 点击 **"高级"** 标签
3. 点击 **"环境变量"**
4. 在 **"系统变量"** 中找到 `Path`，点击 **"编辑"**
5. 点击 **"新建"**，添加：`C:\msys64\mingw64\bin`
6. 点击 **"确定"** 保存所有对话框
7. **重启电脑**（或至少重启所有终端）

### 步骤 4：验证和启用 CGO

按照方法一的步骤 3-5 进行验证。

---

## 🍫 方法三：使用 Chocolatey（如果已安装 Chocolatey）

### 步骤 1：安装 MinGW

以管理员身份打开 PowerShell，运行：

```powershell
choco install mingw -y
```

### 步骤 2：验证和启用 CGO

按照方法一的步骤 3-5 进行验证。

---

## ✅ 验证清单

完成安装后，请逐一检查：

- [ ] `gcc --version` 能显示版本信息
- [ ] `go env CGO_ENABLED` 显示 `1`
- [ ] 项目能够成功编译：`go build .`

---

## 🐛 常见问题解决

### 问题 1：`gcc: command not found`

**原因**：GCC 未添加到 PATH 或需要重启终端

**解决方法**：
1. 检查环境变量：
   - 按 `Win + R`，输入 `sysdm.cpl`
   - 环境变量 → 系统变量 → Path
   - 确认包含 `C:\TDM-GCC-64\bin` 或 `C:\msys64\mingw64\bin`
2. **完全关闭并重新打开**所有终端窗口
3. 如果还是不行，重启电脑

### 问题 2：`CGO_ENABLED=0` 但已安装 GCC

**解决方法**：
```bash
go env -w CGO_ENABLED=1
```

然后验证：
```bash
go env CGO_ENABLED
```

### 问题 3：编译时出现 "undefined reference" 错误

**原因**：缺少某些库文件

**解决方法**：
- 确保安装了完整的工具链（不只是 gcc）
- 如果使用 TDM-GCC，重新安装并确保选择完整安装
- 如果使用 MSYS2，运行：`pacman -S mingw-w64-x86_64-toolchain`

### 问题 4：32位 vs 64位不匹配

**解决方法**：
- 确保安装的是 64 位版本的 GCC
- 检查 Go 版本：`go version` 应该显示 `amd64` 或 `x86_64`

---

## 📝 快速检查脚本

项目目录中已包含检查脚本，可以直接运行：

**Windows 批处理**：
```bash
check_cgo.bat
```

**PowerShell**：
```powershell
powershell -ExecutionPolicy Bypass -File check_cgo.ps1
```

---

## 💡 重要提示

1. **安装后必须重启终端**：环境变量更改后，需要关闭并重新打开所有终端窗口
2. **推荐使用 TDM-GCC**：最简单，一键安装，自动配置
3. **如果不需要 CGO**：当前项目已配置为使用纯 Go 的 SQLite 驱动，可以不安装 CGO
4. **管理员权限**：安装软件时需要管理员权限

---

## 🎯 安装完成后

安装成功后，你可以：

1. 使用需要 CGO 的 Go 包（如 `github.com/mattn/go-sqlite3`）
2. 调用 C 代码
3. 使用其他依赖 CGO 的库

如果只是运行当前项目，**不需要安装 CGO**，因为项目已使用纯 Go 实现的 SQLite 驱动。

---

## 📞 需要帮助？

如果遇到问题：
1. 运行 `check_cgo.bat` 检查环境
2. 查看错误信息
3. 检查环境变量是否正确配置
4. 确保已重启终端或电脑

