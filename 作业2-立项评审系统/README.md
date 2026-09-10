# 黑客松立项评审系统（徐明版）

实现学生立项申请、管理员分配评委、评委评分、学生查看结果等完整流程。

## 技术栈
- 后端：Go + Gin + GORM + MySQL
- 前端：React + Vite
- 认证：JWT

## 启动
1. `docker compose up -d` 启动数据库
2. `cd backend && go run ./cmd/server` 启动后端
3. `cd frontend && npm install && npm run dev` 启动前端

## 访问
- 前端 http://localhost:5173
- 后端 http://localhost:8080
- 管理员：admin@hackathon.com / admin123
