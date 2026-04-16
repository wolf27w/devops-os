# DevOps OS - 研发操作系统骨架

基于 Project 的 DevOps 研发操作系统骨架，仅实现架构与框架，不包含业务逻辑。

## 🎯 系统目标

构建一个基于 Project 的研发操作系统（DevOps OS）骨架，实现完整的架构与框架，为后续业务功能开发提供基础。

## ⚠️ 核心约束

- ❌ **禁止使用数据库**
- ✔ **必须使用文件存储**（JSON / Markdown）
- ✔ **所有数据必须归属 project**
- ✔ **所有 API 必须携带 project_id**
- ✔ **必须实现 RBAC 权限骨架**
- ✔ **必须实现项目隔离机制**
- ❌ **不允许实现任何业务逻辑**（CI/CD / SQL / Workflow 等只做空壳）

## 🧱 技术栈

### 后端
- Go 1.20+
- Gin Web Framework
- 分层架构（api/service/repository/model）
- 文件存储系统（/data）

### 前端
- Vue3 + TypeScript
- Pinia
- Vue Router
- Element Plus

## 🧠 系统核心设计

### 1️⃣ Project 是唯一核心维度
所有资源必须归属于 project：user → project → resource

### 2️⃣ 登录后布局结构
统一 Layout：
- 左侧：模块菜单
- 顶部：Project Switcher（项目切换器）
- 右侧：Router View（内容区）

### 3️⃣ 模块结构（只做占位）
必须生成路由 + 空页面：
- CI/CD
- Repo
- Workflow
- Database
- Wiki
- Navigation
- Settings

## 📁 项目结构

```
devops-os/
├── backend/                    # Go 后端
│   ├── cmd/
│   │   └── main.go            # 启动文件
│   ├── internal/
│   │   ├── api/               # 路由层
│   │   ├── service/           # 业务层（空实现）
│   │   ├── repository/        # 文件存储层
│   │   ├── model/             # 数据结构
│   │   ├── middleware/        # auth + project + rbac
│   │   ├── auth/              # 登录模块
│   │   ├── project/           # 项目管理
│   │   └── common/            # 工具类
│   └── go.mod
├── frontend/                   # Vue 前端
│   ├── src/
│   │   ├── components/        # 组件
│   │   ├── views/             # 页面
│   │   ├── stores/            # Pinia Store
│   │   ├── router/            # 路由
│   │   ├── layout/            # 布局
│   │   ├── api/               # API 调用
│   │   └── utils/             # 工具函数
│   ├── package.json
│   ├── vite.config.ts
│   └── tsconfig.json
└── data/                      # 文件存储
    ├── auth/
    │   └── users.json         # 用户数据
    ├── projects/              # 项目数据
    │   └── {projectId}/
    │       ├── project.json
    │       ├── members.json
    │       ├── ci/
    │       ├── repo/
    │       ├── workflow/
    │       ├── db/
    │       ├── wiki/
    │       ├── nav/
    │       └── logs/
    └── system/
        ├── audit.log
        └── config.json
```

## 🚀 快速开始

### 后端启动

```bash
cd backend
go mod tidy
go run cmd/main.go
```

后端服务将在 `http://localhost:8080` 启动。

### 前端启动

```bash
cd frontend
npm install
npm run dev
```

前端服务将在 `http://localhost:3000` 启动。

## 🔐 默认账户

- **用户名**: admin
- **密码**: admin123

## 🧩 核心基础能力

### 1️⃣ Auth（基础）
- login
- logout
- token middleware

### 2️⃣ Project（基础）
- project create（仅 super_admin）
- project member list
- project context injection

### 3️⃣ Middleware（必须实现）
- authMiddleware（登录校验）
- projectMiddleware（project校验）
- rbacMiddleware（权限校验）

### 4️⃣ File Storage Engine（必须实现）
- ReadJSON(path)
- WriteJSON(path)
- ListFiles(path)
- AppendLog(path)

### 5️⃣ Audit System（只做框架）
- audit log 写入函数
- 不实现业务记录

## 🌐 API 设计

所有 API 必须遵循以下规范：
- 认证：Bearer Token
- 项目隔离：所有资源 API 必须携带 `project_id`
- 权限控制：基于 RBAC 的权限校验

### API 示例
```
GET    /api/projects                    # 获取项目列表
POST   /api/projects                    # 创建项目（仅 super_admin）
GET    /api/projects/{project_id}       # 获取项目详情
GET    /api/projects/{project_id}/ci    # 获取 CI/CD 列表
POST   /api/projects/{project_id}/ci    # 创建 CI/CD（需要权限）
```

## 🎨 前端设计

### 1️⃣ Layout
- Sidebar + Topbar(Project Switcher) + Content

### 2️⃣ Router
预留路由：
- `/app/ci`
- `/app/repo`
- `/app/workflow`
- `/app/db`
- `/app/wiki`
- `/app/nav`
- `/app/settings`

### 3️⃣ Store（Pinia）
必须包含：
- currentProject
- projectList
- switchProject()

### 4️⃣ 页面
所有模块只生成：
- 空页面
- 路由绑定
- 不实现业务

## ⚠️ 禁止事项

必须严格遵守：
- ❌ 不实现 CI/CD 逻辑
- ❌ 不实现 SQL 执行
- ❌ 不实现 Workflow 流程
- ❌ 不使用数据库
- ❌ 不复杂化设计

## 📝 开发指南

### 添加新模块
1. 在后端 `internal/service/` 创建空服务
2. 在 `internal/api/router.go` 添加路由
3. 在前端 `src/views/` 创建空页面
4. 在 `src/router/index.ts` 添加路由
5. 在权限常量中添加对应权限

### 权限配置
权限定义在 `backend/internal/common/constants.go`，基于角色分配：
- super_admin: 所有权限
- admin: 管理权限（除项目创建）
- member: 读写权限
- viewer: 只读权限

## 🔧 配置说明

### 后端配置
- JWT 密钥：`devops-os-secret-key-change-in-production`
- 数据存储路径：`data/`
- 端口：`8080`

### 前端配置
- 开发端口：`3000`
- API 代理：`http://localhost:8080`
- 路由模式：History Mode

## 📄 许可证

MIT License