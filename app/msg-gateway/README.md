1. 在msg-gateway/cmd/wsrpc/pb中创建msg-gateway.proto文件
2. 在msg-gateway/cmd/wsrpc/pb中创建文件ws.proto文件（从msg中复制过来）
3. 在pb目录下执行
   ```shell
    goctl rpc protoc msg-gateway.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../
    ```
4. 可将ws.proto删掉
5. 点进msg-gateway.pb.go文件，在import中添加  pb "goChat/app/msg/cmd/rpc/pb"，然后在
    在msgData前面添加pb.，再找到file_ws_proto_init()，将其删除
6. 进入msg-gateway/cmd/wsrpc中的internal文件夹，依次为其子目录名添加前缀rpc
   - config->rpcconfig
   - server->rpcserver
   - svc->rpcsvc
   - logic->rpclogic
7. 将msg-gateway/cmd/wsrpc/etc目录下的配置文件命名为：msggateway-rpc.yaml


1. 在msg-gateway/cmd/wrcapi/desc中创建msg-gateway.api文件，文件内容如下
   ```syntax = "v1"

   syntax = "v1"
   
   info(
     title: "msg-gateway.api"
     author: "rosy"
     email: "xxx@qq.com"
     version: "v1.0"
   )
   
   type Request {
     Token      string `json:"token" form:"token"`
     SendID     string `json:"sendId" form:"sendId"`
     PlatformID string `json:"platformId" form:"platformId"`
   }
   
   type Response {
     Uid     string `json:"uid"`
     ErrMsg  string `json:"errMsg"`
     Success string `json:"success"`
   }
   
   service msgGateway {
     @doc "msg gateway api"
     @handler msggateway
     post / (Request) returns (Response)
   }
   
   ```
2. 在msg-gateway/cmd/wrcapi/desc目录下执行
   ```shell
   goctl api go -api msg-gateway.api -dir ../  -style=goZero
   ```
3. 将msg-gateway/cmd/wrcapi/etc重命名为msggateway-ws.yaml,然后移动到wsrpc/etc目录下
4. 将internal中的文件重命名
   - config -> wsconfig
   - svc -> wssvc
   - logic -> wslogic
5. 将wsapi/internal中的所有文件移动到wsrpc/internal中
6. 进入文件中修改import的报错
7. 将wsapi目录下的msggateway.go文件中的main函数复制到wsrpc中的msggateway中,最终为：

   ```
   package main
   
   import (
       "flag"
       "fmt"
       "github.com/zeromicro/go-zero/core/conf"
       "github.com/zeromicro/go-zero/core/logx"
       "github.com/zeromicro/go-zero/core/service"
       "github.com/zeromicro/go-zero/rest"
       "github.com/zeromicro/go-zero/zrpc"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/handler"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/rpcconfig"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/rpcserver"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/wsconfig"
       "goChat/app/msg-gateway/cmd/wsrpc/internal/wssvc"
       "goChat/app/msg-gateway/cmd/wsrpc/pb"
       "google.golang.org/grpc"
       "google.golang.org/grpc/reflection"
       "time"
   )
   
   var wsConfigFile = flag.String("w", "etc/msggateway-ws.yaml", "ws config file")
   var rpcConfigFile = flag.String("r", "etc/msggateway-rpc.yaml", "rpc config file")
   
   func ws() {
       flag.Parse()
   
       var c wsconfig.Config
       conf.MustLoad(*wsConfigFile, &c)
   
       server := rest.MustNewServer(c.RestConf)
       defer server.Stop()
   
       ctx := wssvc.NewServiceContext(c)
       handler.RegisterHandlers(server, ctx)
   
       fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
       server.Start()
   }
   
   func rpc() {
       flag.Parse()
   
       var c rpcconfig.Config
       conf.MustLoad(*rpcConfigFile, &c)
       ctx := rpcsvc.NewServiceContext(c)
   
       s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
           pb.RegisterOnlineMessageRelayServiceServer(grpcServer, rpcserver.NewOnlineMessageRelayServiceServer(ctx))
   
           if c.Mode == service.DevMode || c.Mode == service.TestMode {
               reflection.Register(grpcServer)
           }
       })
       defer s.Stop()
   
       fmt.Printf("Starting rpc rpcserver at %s...\n", c.ListenOn)
       s.Start()
   }
   
   func main() {
       go ws()
       logx.Info("ws 启动成功 等待1秒启动 rpc")
       time.Sleep(time.Second)
       rpc()
   }
   ```
8. 在wsrpc目录下运行go run msggateway.go显示如下信息表示运行成功。
    ```shell
   {"@timestamp":"2023-05-23T22:55:44.942+08:00","caller":"wsrpc/msggateway.go:64","content":"ws 启动成功 等待1秒启动 rpc","level":"info"}
   Starting server at 0.0.0.0:8888...
   Starting rpc rpcserver at 0.0.0.0:8080...
    ```
9. 可将wsapi文件删去