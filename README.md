Go Distributed Student System

一个基于 Go + Gin 实现的简单微服务学生成绩管理系统。
系统将业务拆分为多个服务，通过 HTTP 进行通信，实现基础的服务注册与发现机制，并提供 Web 界面进行管理。

该项目主要用于学习 Go Web 开发、微服务架构设计以及服务间通信。

项目功能

系统主要提供以下功能：

用户登录（JWT认证）

学生信息管理

成绩录入与查询

学生成绩统计

服务注册与发现

日志服务统一处理

Web 管理界面

Swagger API 文档

技术栈
技术	说明
Go	后端开发语言
Gin	Web 框架
GORM	ORM 框架
MySQL	数据库
JWT	用户认证
Docker	容器化部署
Swagger	API 文档
系统架构

项目采用简单的微服务架构，系统拆分为 4 个主要服务：

           +----------------+
           |    Portal      |
           |  (Web 前端)    |
           +--------+-------+
                    |
                    v
           +----------------+
           |   Grades       |
           |  成绩管理服务   |
           +--------+-------+
                    |
                    v
                 MySQL

           +----------------+
           |  Registry      |
           | 服务注册中心    |
           +----------------+

           +----------------+
           |   Log Service  |
           |   日志服务     |
           +----------------+

各服务职责：

Portal Service

提供 Web 管理界面

调用 Grades 服务获取数据

实现用户登录功能

Grades Service

学生信息管理

成绩管理

数据持久化

Registry Service

服务注册

服务发现

心跳检测

Log Service

统一日志记录

通过 HTTP 接收日志信息

目录结构
distributed
│
├── cmd
│   ├── gradingservice
│   ├── logservice
│   ├── portal
│   └── registryservice
│
├── grades
│   ├── handler.go
│   ├── service.go
│   └── repository.go
│
├── portal
│   ├── templates
│   └── handler.go
│
├── registry
│   └── registry.go
│
├── log
│   └── log.go
│
├── pkg
│   ├── jwt
│   ├── middleware
│   └── response
│
├── docker-compose.yml
└── README.md
启动方式
1 安装依赖

确保已安装：

Go 1.20+

Docker

MySQL

2 启动 MySQL
docker-compose up mysql

默认数据库配置：

user: root
password: 123456
database: grades
3 启动 Registry Service
go run cmd/registryservice/main.go

默认端口：

3000
4 启动 Log Service
go run cmd/logservice/main.go
5 启动 Grades Service
go run cmd/gradingservice/main.go
6 启动 Portal Service
go run cmd/portal/main.go

访问：

http://localhost:8000
API 文档

项目集成 Swagger。

启动服务后访问：

http://localhost:8000/swagger/index.html

即可查看 API 文档。

登录账号

默认账号：

username: admin
password: 123456
示例接口
获取所有学生
GET /students
获取学生成绩
GET /students/{id}/grades
新增学生
POST /students
新增成绩
POST /students/{id}/grades
项目亮点

基于 Gin 框架实现 RESTful API

使用 GORM + MySQL 进行数据持久化

实现 JWT 用户认证

设计 Registry 服务实现服务注册与发现

将系统拆分为 Portal、Grades、Registry、Log 四个服务

使用 Docker Compose 管理数据库环境

提供 Swagger API 文档

项目目的

该项目主要用于学习：

Go Web 开发

微服务架构设计

服务间 HTTP 通信

JWT 认证

Docker 基础部署