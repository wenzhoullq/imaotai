# 配置文件
* AK

在百度地图中注册AK,输入至配置文件中的AK

* 定时任务

imaotai订购时间在9:00 - 10:00,reservation的配置在这时间之内,其他的定时任务在非该时段即可 

# 数据库

* 建库语句

```sql
  CREATE DATABASE `imaotai`  DEFAULT CHARACTER SET utf8mb4 ;
```

# 接口使用

* 发送验证码

```cmd
curl --location 'http://localhost:8080/imaotai/user/getVerifyCode' \
--header 'Content-Type: application/json' \
--data '{
    "mobile":"138xxxxxxxx"
}'



```
* 登录

```cmd
```cmd
curl --location 'http://localhost:8080/imaotai/user/login' \
--header 'Content-Type: application/json' \
--data '{
"code":"*******",
"mobile":"138******"
}'
* 
```

* 更新地址

地址越详细越好

```cmd
curl --location 'http://localhost:8080/imaotai/user/updateAddress' \
--header 'Content-Type: application/json' \
--data '{
    "address":"天津市东丽区",
    "mobile":"138*******"
}'
```

* 更新Token

Token一个月过期,需要获得验证码重新更新

```cmd
curl --location 'http://localhost:8080/imaotai/user/updateToken' \
--header 'Content-Type: application/json' \
--data '{
    "code":"******",
    "mobile":"138********"
}'
```

* 查看用户状态

查看用户是否在定时任务中执行

```cmd
curl --location 'http://localhost:8080/imaotai/user/getUserStatus' \
--header 'Content-Type: application/json' \
--data '{
    "mobile":"138********"
}'
```