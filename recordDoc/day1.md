# Day1

## 完成接口
api
- register

rpc
- getUserByUsername
- getUserById
- LoginById

## 主要"难点"
1. 用什么go-zero的数据库还是用xorm？
   最后用了自带的
2. 如何处理错误？
   建立common/xerror文件夹，用来处理错误

## 遗留问题
1. 错误处理有冗余
2. 日志问题
3. go-zero的数据库+缓存处理机制是什么？

## TODO
1. rpc接口GetUserByIds，LoginByPassword
2. api接口LoginByPassword