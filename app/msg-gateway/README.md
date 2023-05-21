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