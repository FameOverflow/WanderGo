# WanderGo
宇宙超级无敌暴龙战士小分队的校园地图项目——漫Go


[TOC]
    
##*请求方式都使用POST*


## 用户注册部分
- 发送邮件接口
- 用户注册接口
##### 请求URL
- ` "xxx/register/captcha" `
- ` "xxx/register" `
  
####发送邮件
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_name |是  |String |邮箱  |
|user_password |是  |String | 密码    |


##### 返回示例 


``` 
  {
    "data": {
     "message":"验证码已发送"
    }
  }
```
``` 
  {
    "data": {
    "message": "该账号已被注册"
    }
  }
```
####用户注册
##### 参数
|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_account |是  |String |邮箱  |
|user_name |是  |String |昵称  |
|user_password |是  |String | 密码    |
|user_captcha |是  |Number | 验证码    |


##### 返回示例 
``` 
  {
    "data": {
     "message":"验证码错误"
    }
  }
```
``` 
  {
    "data": {
    "message": "注册成功"
    }
  }
```
## 用户登录部分
##### 请求URL
- ` "xxx/login" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_account |是  |String |邮箱  |
|user_password |是  |String | 密码    |


##### 返回示例 


``` 
  {
    "data": {
     "message":"登录成功"
    }
  }
```



## 退出登录部分
##### 请求URL
- ` "xxx/exit" `
  
##### 参数


无
##### 返回示例 


``` 
  {
    "data": {
     "message": "您已成功退出该账号"
    }
  }
```
## 找回密码部分
- 发送邮件接口
- 找回密码接口
##### 请求URL
- ` "xxx/passwords/forget/captcha" `
- ` "xxx/passwords/forget" `
  
####发送邮件
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_account |是  |String |邮箱  |



##### 返回示例 


``` 
  {
    "data": {
     "message":"验证码已发送"
    }
  }
```
####找回密码
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_account |是  |String |邮箱  |
|new_pwd |是  |String |新密码  |
|user_captcha |是  |Number |验证码  |



##### 返回示例 


``` 
  {
    "data": {
     "message":"您已成功修改密码,请登录"
    }
  }
```



## 改密部分
##### 请求URL
- ` "xxx/passwords/change" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|old_pwd |是  |String |旧密码  |
|current_pwd |是  |String | 新密码    |


##### 返回示例 


``` 
  {
    "data": {
     "message":"您已成功修改密码，请登录"
    }
  }
```
## 改名部分
##### 请求URL
- ` "xxx/names/change" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|user_name |是  |String |新昵称  |


##### 返回示例 


``` 
  {
    "data": {
     "message": "修改用户名成功"
    }
  }
```
## 发布评论部分
##### 请求URL
- ` "xxx/comments/post" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|text |是  |String |文本  |
|position{"x","y"} |是  |Number |发布评论时的经纬度  |
##### 请求示例 


``` 
 {
  "text":"这1是一条评论",
  "position":{
      "x":115.79941,
      "y":28.656973
  }
}
   
```
##### 返回示例 


``` 
  {
    "data": {
        "place_uid": 1,
        "comment_uuid": "445ac4d24ad8d7878077f0c5d80734d0",
        "message": "成功发布评论"
    }
  }
```
##### 返回参数说明


|参数名|类型|说明|
|:----    |:---|:----- |-----   |
|place_uid |Number |该评论所属建筑的id  |
|comment_uuid |String  |该评论的唯一编号   |


## 漫游部分
##### 请求URL
- ` "xxx/comments/roam" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|x |是  |Number |经度  |
|y |是  |Number |纬度  |


##### 返回示例 


``` 
  {
    "data": {
        "place_uid": "1",
        "text": "这是一条在图书馆的评论"
    }
  }
```
##### 返回参数说明


|参数名|类型|说明|
|:----    |:---|:----- |-----   |
|place_uid |Number |该评论所属建筑的id  |
|text |String  |文本   |


## 点赞部分
##### 请求URL
- ` "xxx/comments/like" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|comment_uuid |是  |String |该评论的唯一编号  |


##### 返回示例 


``` 
  {
    "data": {
     "message":"点赞成功"
    }
  }
```
## 获取位置信息部分
##### 请求URL
- ` "xxx/places/get" `
  
##### 参数


|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|x |是  |Number |经度  |
|y |是  |Number |纬度  |


##### 返回示例 


``` 
  {
    "data": {
     "place_id": [
        1,
        2
     ]
    }
  }
```
##### 返回参数说明


|参数名|类型|说明|
|:----    |:---|:----- |-----   |
|place_id |Number |250m范围内的地点的id  |
## 载入用户数据部分
##### 请求URL
- ` "xxx/begin/user" `
  
##### 参数


无


##### 返回示例 


``` 
  {
    "data": {
     "comments": [
        {
            "user_account": "suzimiyaharuhi@fake.com",
            "date": "2023-12-09 13:17:02",
            "text": "我还在基础实验大楼",
            "comment_uuid": "c2ffa584cefdb5cd4ac3bda8cf168428",
            "place_uid": 2,
            "star_cnt": 0
        },
        {
            "user_account": "suzimiyaharuhi@fake.com",
            "date": "2023-12-09 13:16:53",
            "text": "我在基础实验大楼",
            "comment_uuid": "3126610b117a7d946881a10d53fe1d16",
            "place_uid": 2,
            "star_cnt": 0
        },
        {
            "user_account": "suzimiyaharuhi@fake.com",
            "date": "2023-12-09 13:16:16",
            "text": "我在龙腾湖",
            "comment_uuid": "bfeaae697090da6151187b859c3c8f05",
            "place_uid": 29,
            "star_cnt": 0
        },
    ],
    "message": "正处于登录状态",
    "user_name": "haruhi"
    }
  }
```
##### 返回参数说明


|参数名|类型|说明|
|:----    |:---|:----- |-----   |
|user_name |String |用户名  |
|Comments |  |用户发过的评论  |
## 载入地点数据部分
##### 请求URL
- ` "xxx/begin/places" `
  
##### 参数


无


##### 返回示例 


``` 
  {
    "data": {
    "comments_in_place": [
        {
            "user_account": "suzumiyaharuhi@fake.com",
            "date": "2023-12-09 13:17:02",
            "text": "我还在基础实验大楼",
            "comment_uuid": "c2ffa584cefdb5cd4ac3bda8cf168428",
            "place_uid": 2,
            "star_cnt": 0
        },
        {
            "user_account": "suzumiyaharuhi@fake.com",
            "date": "2023-12-09 13:16:53",
            "text": "我在基础实验大楼",
            "comment_uuid": "3126610b117a7d946881a10d53fe1d16",
            "place_uid": 2,
            "star_cnt": 0
        }
    ]
    }
  }
```


## 点亮部分
##### 请求URL
- ` "xxx/palces/mark" `
  
##### 参数
|参数名|类型|说明|
|:----    |:---|:----- |-----   |
|id |Number |要点亮的地点的id  |


##### 返回示例 
``` 
  {
    "data": {
   "message":"已成功点亮"
    }
  }
```
## 获取STSToken令牌部分，可用于OSS服务鉴权,成功响应为sts:token
    engine.POST("/GetSTS", oss.GetSTS)
##### 请求URL
- ` "xxx/sts/get" `
##### 参数
无
##### 返回示例 
``` 
  {
    "data": {
   "sts":token
    }
  }
```