## verifyToken
1. 在usercenter/cmd/rpc/pb常见im_user.proto，内容如下：
 
    ```
    syntax = "proto3";
    
    option go_package = "./pb";
    
    package imuser;
    
    message VerifyTokenReq {
      string token = 1;
      string sendID = 2;
    }
    
    message VerifyTokenResp {
      string uid = 1;
      bool success = 2;
      string errMsg = 3;
    }
    
    service imUserService {
      rpc VerifyToken(VerifyTokenReq) returns (VerifyTokenResp);
    }
    ```
2. 在pb目录下执行
   ```
   goctl rpc protoc im_user.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../
   ```
3. 在verifylogic中粘贴下列内容,根据报错拉取，复制相关的依赖文件：
   ```
     jwtUtils.GetClaimFromToken()
     l.rep.GetTokenMap()
     rep.DeleteToken()
     rep.RenewalToken()
   ```
4. 