# 📊 Enterprise SaaS Dashboard (WPS Elite Training Project)

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?style=flat&logo=react)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green)

> **🎯 项目背景**: 本项目的目标是实现一套对标金山办公（WPS）技术标准的云原生架构。它是从 Java 生态向 Go 微服务转型的实战演练。

## 📖 项目简介 (Introduction)

这是一个采用前后端分离架构的企业级数据看板应用（MVP版本）。我并没有将其视为简单的 CRUD 练习，而是将其作为实践 **Go 高并发模式** 与 **现代 React 工程化** 的载体。

**核心特性 (Core Features):**
- 🚀 **高性能后端**: 基于 **Gin** 框架构建 RESTful API，严格遵循 **Uber Go Style Guide** 编码规范，确保代码的可读性与高性能。
- 📊 **数据可视化**: 前端采用 **React + TypeScript + Vite** 技术栈，集成可视化图表库，模拟企业级数据展示。
- 🛡️ **容器化交付**: 支持 Docker & Docker Compose 一键编排，解决跨环境一致性问题（To Be Implemented）。

## 🛠️ 技术栈 (Tech Stack)

| 层级 | 技术选型 | 选型理由 (对标 WPS 技术体系) |
| :--- | :--- | :--- |
| **Backend** | **Go + Gin** | 利用 Goroutine 轻量级协程处理高并发，贴合云原生生态。 |
| **Frontend** | **React + TypeScript** | 强类型约束保障大型项目可维护性，组件化开发提升效率。 |
| **Database** | **MySQL + GORM** | 经典的各种业务场景下的数据持久化方案。 |
| **DevOps** | **Docker** | 模拟真实的生产环境部署流程，实现“一次构建，到处运行”。 |

## 📂 目录结构 (Structure)

```text
saas-dashboard/
├── backend/      # Go 后端服务 (Gin + RESTful API)
│   ├── main.go   # 入口文件
│   └── go.mod    # 依赖管理
├── frontend/     # React 前端应用 (Vite)
│   ├── src/      # 源代码
│   └── package.json
└── README.md     # 项目文档

