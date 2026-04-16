#!/bin/bash

echo "🚀 Starting DevOps OS..."

# 检查后端依赖
echo "📦 Checking backend dependencies..."
cd backend
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod not found"
    exit 1
fi

# 检查前端依赖
echo "📦 Checking frontend dependencies..."
cd ../frontend
if [ ! -f "package.json" ]; then
    echo "❌ package.json not found"
    exit 1
fi

# 启动后端
echo "🔧 Starting backend server..."
cd ../backend
go run cmd/main.go &
BACKEND_PID=$!

# 等待后端启动
sleep 3

# 启动前端
echo "🎨 Starting frontend server..."
cd ../frontend
npm run dev &
FRONTEND_PID=$!

echo "✅ DevOps OS started!"
echo "🌐 Backend: http://localhost:8080"
echo "🎨 Frontend: http://localhost:3000"
echo "🔐 Default credentials: admin / admin123"

# 捕获退出信号
trap "echo '🛑 Stopping DevOps OS...'; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit 0" INT TERM

# 等待进程
wait $BACKEND_PID $FRONTEND_PID