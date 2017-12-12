# Agenda Service

[![Build Status](https://travis-ci.org/Mensu/Agenda-cli-service-Go.svg?branch=master)](https://travis-ci.org/Mensu/Agenda-cli-service-Go)

## 概述
	利用命令行 或 web 客户端调用远端服务是服务开发的重要内容。其中，要点是如何实现 API First 开发，使得团队协作变得更有效率。
## 任务目标
	1. 熟悉 API 设计工具，实现从资源（领域）建模，到 API 设计的过程
	2. 使用 Github ，通过 API 文档，实现 agenda 命令行项目 与 RESTful 服务项目同步开发
	3. 使用 API 设计工具提供 Mock 服务，两个团队独立测试 API
	4. 使用 travis 测试相关模块
	5. 利用 dockerfile 在 docker hub 上构建一个镜像，同时包含 agenda cli 和 agenda service， 如果 mysql 包含 服务器 和 客户端一样
## agenda 开发项目基本要求
	cli 目录
	service 目录
	.travis
	apiary.apib
	dockerfile
	LICENSE
	README.md
	README-yourid.md 记录你的工作摘要（个人评分依据）
	API 开发 
	使用 API Blueprint 设计 API
	资源 URL 命名符合 RESTful 设计标准
	资源 CRUD 基本完整
	API 客户端开发 
	可用命令 5 个以上
	必须有 XXX-test.go 文件
	服务端开发 
	使用 sqlite3 作为数据库
	建议使用课程提供的服务端框架
	必须有 XXX-test.go 文件
	容器镜像制作 
	在 docker hub 上生成镜像
	base 镜像 go-1.8
	需要加载 sqlite3
	同时包含客户端与服务器
	README.md 
	有 build pass 标签
	有简短使用说明
	有系统测试的结果（包含如何下载镜像，如何启动服务器，如何使用命令行，cli 的 mock 测试结果， 综合系统测试结果）
	README-yourid.md 
	fork 项目的位置
	个人工作摘要（每次提交）
	项目小结
	注意
	开发过程中会有许多问题，如 golang/x 包本地构建镜像无法下载等。
	yourid 建议不要使用个人真实姓名
## 实验过程运行截图
### 开始时没有用户
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E5%BC%80%E5%A7%8B%E6%97%B6%E6%B2%A1%E6%9C%89%E7%94%A8%E6%88%B7.png)
### 建立了用户后查询，以及创立多个用户
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7.png)
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E5%A4%9A%E4%B8%AA%E7%94%A8%E6%88%B7.png)
### 开始时没有会议
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AE%AE.png)
### 成功创立会议
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E6%88%90%E5%8A%9F%E5%88%9B%E5%BB%BA%E4%BC%9A%E8%AE%AE.png)
### 用用户名字查询其相关信息
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E6%9F%A5%E8%AF%A2%E7%94%A8%E6%88%B7.png)
### 用会议名字自查询相关信息
![](https://github.com/453326526/agenda-go-server/blob/master/photos/%E5%90%8D%E5%AD%97%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AE%AE.png)

## Docker安装过程
### 使用 APT 镜像源 安装

	由于官方源使用 HTTPS 以确保软件下载过程中不被篡改。因此，我们首先需要添加使用 HTTPS 传输的软件包以及 CA 证书。

	```bash
	$ sudo apt-get update

	$ sudo apt-get install \
	    apt-transport-https \
	    ca-certificates \
	    curl \
	    software-properties-common
	```

	鉴于国内网络问题，强烈建议使用国内源，官方源请在注释中查看。

	为了确认所下载软件包的合法性，需要添加软件源的 GPG 密钥。

	```bash
	$ curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -


### 官方源

### $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

	然后，我们需要向 `source.list` 中添加 Docker 软件源

	```bash
	$ sudo add-apt-repository \
	    "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu \
	    $(lsb_release -cs) \
	    stable"


	更新 apt 软件包缓存，并安装 `docker-ce`：

	```bash
	$ sudo apt-get update

	$ sudo apt-get install docker-ce
	```

### 启动 Docker CE

	```bash
	$ sudo systemctl enable docker
	$ sudo systemctl start docker
	```

	Ubuntu 14.04 请使用以下命令启动：

	```bash
	$ sudo service docker start
	```
