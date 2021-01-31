Gin-example
===
# 一、初衷
私以为，一个优秀的代码框架，或许不能提升开发效率，或许不能提高系统性能，但一定要能改善软件开发者的幸福指数。故而，本案例将会着重在软件开发工程系的幸福指数上大作文章，着力于性能改造、效率提升的朋友可以不必在此消磨时间。

为了提升程序开发者的幸福指数，本系统主要提供以下特性：
1. 可测
2. 自动报警
3. 链路跟踪（待定）

为上手本项目，您需要掌握或之后需要了解:
1. 日志系统，例如 logrus
2. 单元测试
3. opentracing (待定)

# 二、架构
一个成熟完备的 web 项目中，大抵都需要维护一个主服务、一个后台，稍显复杂些的项目，可能还需要一些定时任务，需要聊天支持等。然而，中小型项目中，后端人力资源紧缺，可能仅为两到三人甚至更少，面对日益增长的需求，迫切需要一个简单、清晰且便于迭代快的软件架构。

本系统为中小项目而生，采用严格的分层设计，为中小型项目的 api 服务提供完备的解决方案。整体上自上而下共分为三个大层，分别是业务层、通用逻辑层以及模型层。为了最大限度的节约人力成本、减少重复，在优雅和高效之间还是做了一些取舍。

## 2.1 软件分层
分层设计中，调用链不可颠倒，本项目亦如此，只允许向下的单向调用链，上层可以跨级调用下层。同层之间减少调用，规避耦合。

### 2.1.1 业务层

业务层逻辑代码盘踞于 ./apps 目录下，存储四个子项目，依次为 api、admin、cron 以及 chat 项目。多个项目存储在一个代码仓库中可能不是最好的选择，但是为了提升代码的复用率，这样或许是成本最低的方案。

业务层中每个子项目根据路由和逻辑实现分为上下分为两层，上半层定义接口，即解析 gin.Context 请求参数，并过滤错误。考虑到gin.Context 对单元测试并不友好，逻辑实现全部转交到 service 中去。另外，上半层每一个路由方法即意味着赋予用户的一个权限，原子不可分割，故而需要谨慎定义接口，如果有割裂的权限设置，请另开辟新的接口服务业务。

下半层执行真实的逻辑，需要保证代码的可测性，可以为上层提供基本的复用，不需要和上半层保证严格一一映射关系。各个子项目对应的service应当保持隔离，不得横向调用，减少耦合.

上半层：router
下半层：services

### 2.1.2 通用逻辑层
上层的复用体 暂定为 logic 文件，命名应当多考虑

### 2.1.3 模型层
模型层中，除了数据模型外，还会有缓存模型、消息队列、第三方接口等。这里旨在把没有业务依赖的内容都抽象成模型。

上半层:    数据模型        消息队列          通知模型 
下半层：mysql redis    redis-queue    dingTalk appPush


### 2.1.4 基础设施层
工具函数、静态参数定义

## 2.2 目录结构

```
├── Readme.md
├── a.txt
├── apps
│   ├── api
│   ├── chat
│   ├── cron
│   └── dashboard
├── go.mod
├── go.sum
├── logic
│   └── config.go
├── middleware
│   ├── auth.go
│   └── rate.go
└── models
    ├── cache
    ├── mysql
    ├── posts.go
    ├── queue
    └── users.go
```

# 三、todo 清单

1. orm、原生 sql、gendry 取舍问题
2. 容器化考虑
3. swagger
