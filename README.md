# go-micro-service

go微服务学习示例


## 目录结构

```
.
├── cmd                启动命令
│   └── logservice    
│       └── main.go    日志服务启动脚本
│   └── registryservice    
│       └── main.go    注册服务启动脚本
├── log                日志服务
│   └── service.go
├── registry           注册服务
│   └── registration.go
│   └── service.go
├── service            服务
│   └── service.go
├── go.mod   
└── README.md  
```

## 视频
bilibili [Go语言编写简单分布式系统](https://www.bilibili.com/video/BV1ZU4y1577q?p=1)



## 用法

### 启动
```shell
go run cmd/registryservice/main.go
```
### 注册服务
```shell
curl -d '{"serviceName":"Business services","ServiceUrl":"http://localhost:5000"}' -i -X POST http://localhost:3000/services
```