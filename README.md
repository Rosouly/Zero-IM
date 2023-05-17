# GoChat

```text
基于go-zero即时通信系统
```

使用到的命令
```text
# 创建API服务
goctl api new usercenter

# 启动api/rpc服务
go run usercenter.go -f etc/usercenter-api.yaml

# 使用api文件生成代码（在desc目录下）goctl api go -api usercenter.api -dir ../  -style=goZero

# 使用proto文件生成代码（在pb目录下）
goctl rpc protoc usercenter.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../

# 生成model
goctl model mysql datasource -url="root:PXDN93VRKUm8TeE7@tcp(127.0.0.1:33069)/zeroim" -table="user" -c -dir .

# 依赖搭建：docker-compose-env(在根目录下)
docker-compose -f docker-compose-env.yml up -d
```

技术栈
- go-zero
- redis
- mysql or mongodb?
- kafka
- etcd
- k8s
- prometheus
- jaeger
- jenkins
- docker

Todo

用户模块
- [ ] 用户注册
- [ ] 用户登录
- [ ] 获取用户信息
- [ ] 修改用户信息
- [ ] 好友申请
- [ ] 好友申请列表
- [ ] 同意好友申请
- [ ] 拒绝好友申请
- [ ] 好友列表
msg模块
- [ ] 发送消息
msg-transfer模块
- [ ] 消息转发
- [ ] 消息持久化
msg-push模块
- [ ] 消息推送

