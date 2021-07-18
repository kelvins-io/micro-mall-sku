# micro-mall-sku

#### 介绍
微商城-商品系统

#### 软件架构
软件架构说明

#### 框架，库依赖
kelvins框架支持（gRPC，cron，queue，web支持）：https://gitee.com/kelvins-io/kelvins   
g2cache缓存库支持（两级缓存）：https://gitee.com/kelvins-io/g2cache   

#### 安装教程

1.仅构建  sh build.sh   
2 运行  sh build-run.sh   

#### 使用说明
配置参考
```toml
[kelvins-server]
EndPoint = 8080
IsRecordCallResponse = true

[kelvins-logger]
RootPath = "./logs"
Level = "debug"

[kelvins-mysql]
Host = "127.0.0.1:3306"
UserName = "root"
Password = "ddd"
DBName = "micro_mall_sku"
Charset = "utf8mb4"
PoolNum =  10
MaxIdleConns = 5
ConnMaxLifeSecond = 3600
MultiStatements = true
ParseTime = true

[kelvins-redis]
Host = "127.0.0.1:6379"
Password = "xxx"
DB = 1
PoolNum = 10

[email-config]
User = "ddd@qq.com"
Password = "dddd"
Host = "smtp.qq.com"
Port = "465"

[mongodb-config]
Uri = "mongodb://localhost:27017"
Username = "admin"
Password = "ddddd"
Database = "micro_mall_sku"
AuthSource = "admin"
MaxPoolSize = 9
MinPoolSize = 3
```
#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
