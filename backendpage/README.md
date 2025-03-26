# Storage System Backend

这是一个使用 Go 语言和 MySQL 构建的后端系统，提供用户认证和授权功能。

## 功能特性

- 用户注册
- 用户登录
- JWT 认证
- 用户信息获取
- 跨域支持

## 技术栈

- Go 1.20+
- Gin Web 框架
- GORM ORM
- MySQL 数据库
- JWT 认证

## 安装和运行

1. 确保已安装 Go 1.20 或更高版本
2. 克隆项目
3. 安装依赖：
   ```bash
   go mod tidy
   ```
4. 配置环境变量：
   - 复制 `.env.example` 为 `.env`
   - 修改数据库配置和 JWT 密钥

5. 创建 MySQL 数据库：
   ```sql
   CREATE DATABASE storage_system;
   ```

6. 运行项目：
   ```bash
   go run main.go
   ```

## API 端点

### 公开接口

- POST `/api/register` - 用户注册
- POST `/api/login` - 用户登录

### 需要认证的接口

- GET `/api/user/profile` - 获取用户信息

## 数据库表结构

### users 表

```sql
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

## 前端集成

前端需要在请求头中添加 JWT token：

```javascript
headers: {
  'Authorization': `Bearer ${token}`
}
``` 