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

  

| 参数名         | 必选 | 类型   | 说明       |
| ------------- | ---- | ------ | ---------- |
| user_name      | 是   | String | 邮箱       |
| user_password  | 是   | String | 密码       |


  

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

| 参数名         | 必选 | 类型   | 说明       |
| -------------- | ---- | ------ | ---------- |
| user_account   | 是   | String | 邮箱       |
| user_name      | 是   | String | 昵称       |
| user_password  | 是   | String | 密码       |
| user_captcha   | 是   | Number | 验证码     |


  

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

  

| 参数名         | 必选 | 类型   | 说明       |
| -------------- | ---- | ------ | ---------- |
| user_account   | 是   | String | 邮箱       |
| user_password  | 是   | String | 密码       |

  

##### 返回示例

  

```

  {

    "data": {

     "message":"登录成功"

    }

  }

```
Headers内有响应token,例如：
_token=Bearer+eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoiMTQwNDI4NzE0N0BxcS5jb20iLCJ0aW1lIjoxNzAyMTg1NzUwLCJleHAiOjE3MDMzOTUzNTAsImlzcyI6IkZsMFJlbmNFc3MiLCJuYmYiOjE3MDIxODU3NTB9.fL_3UKUvwCOu70pk0By_FkMhw9ITlfHLS0IfCxtxxcA; Path=/; Domain=localhost; Max-Age=2592000; HttpOnly
在鉴权请求中设置Authorization:Bearer token,
例如Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoiMTQwNDI4NzE0N0BxcS5jb20iLCJ0aW1lIjoxNzAyMTg1NzUwLCJleHAiOjE3MDMzOTUzNTAsImlzcyI6IkZsMFJlbmNFc3MiLCJuYmYiOjE3MDIxODU3NTB9.fL_3UKUvwCOu70pk0By_FkMhw9ITlfHLS0IfCxtxxcA
  

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

  


| 参数名         | 必选 | 类型   | 说明       |
| -------------- | ---- | ------ | ---------- |
| user_account   | 是   | String | 邮箱       |
  
  

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

  
| 参数名       | 必选 | 类型   | 说明     |
| ------------ | ---- | ------ | -------- |
| user_account | 是   | String | 邮箱     |
| new_pwd      | 是   | String | 新密码   |
| user_captcha | 是   | Number | 验证码   |


  
  

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

  

| 参数名       | 必选 | 类型   | 说明     |
| ------------ | ---- | ------ | -------- |
| old_pwd      | 是   | String | 旧密码   |
| current_pwd  | 是   | String | 新密码   |

  

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

  

| 参数名     | 必选 | 类型   | 说明   |
| ---------- | ---- | ------ | ------ |
| user_name  | 是   | String | 新昵称 |


  

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

  

| 参数名             | 必选 | 类型   | 说明                   |
| ------------------ | ---- | ------ | ---------------------- |
| text               | 是   | String | 文本                   |
| position{"x","y"} | 是   | Number | 发布评论时的经纬度     |


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

  
| 参数名         | 类型   | 说明               |
| -------------- | ------ | ------------------ |
| place_uid      | Number | 该评论所属建筑的id |
| comment_uuid   | String | 该评论的唯一编号   |

  

## 漫游部分

##### 请求URL

- ` "xxx/comments/roam" `

##### 参数

  
| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| x      | 是   | Number | 经度   |
| y      | 是   | Number | 纬度   |


  

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

  
| 参数名      | 类型   | 说明               |
| ----------- | ------ | ------------------ |
| place_uid   | Number | 该评论所属建筑的id |
| text        | String | 文本               |

  

## 点赞部分

##### 请求URL

- ` "xxx/comments/like" `

##### 参数

  

| 参数名        | 必选 | 类型   | 说明               |
| ------------- | ---- | ------ | ------------------ |
| comment_uuid  | 是   | String | 该评论的唯一编号   |

  

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

  

  
| 参数名 | 必选 | 类型   | 说明   |
| ------ | ---- | ------ | ------ |
| x      | 是   | Number | 经度   |
| y      | 是   | Number | 纬度   |

  

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

  

| 参数名    | 类型   | 说明                   |
| --------- | ------ | ---------------------- |
| place_id  | Number | 250m范围内的地点的id   |


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

  
| 参数名     | 类型   | 说明         |
| ---------- | ------ | ------------ |
| user_name  | String | 用户名       |
| Comments   |        | 用户发过的评论|


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

##### 返回参数说明

  

| 参数名     | 类型   | 说明         |
| ---------- | ------ | ------------ |
| user_name  | String | 用户名       |
| Comments   |        | 用户发过的评论|

  

## 点亮部分

##### 请求URL

- ` "xxx/palces/mark" `

##### 参数

| 参数名  | 类型   | 说明               |
| ------- | ------ | ------------------ |
| id      | Number | 要点亮的地点的id   |


  

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
   "sts": "CAIS5QF1q6Ft5B2yfSjIr5bFJvTCprtz8pCeYV/YrjNkf7ZIp4jztzz2IHhMe3FvBO8XsPQwm21U6fgblqpoV4QdrI1y/GcpvPpt6gqET9frma7ctM4p6vCMHWyUFGSIvqv7aPn4S9XwY+qkb0u++AZ43br9c0fJPTXnS+rr76RqddMKRAK1QCNbDdNNXGtYpdQdKGHaOITGUHeooBKJUBQz5Vsn1Dolt//imJLE0HeE0g2mkN1yjp/qP52pY/NrOJpCSNqv1IR0DPGZiHEItEUVpfkm0/MZpWiZ58v2GEVK8/tZumBDc3kFGoABIpgyUCKAhlRQfALPOrgoW0+fkIHd/O0oK9p7J9At8Y0uMN8wHKymfS1ua1PSBqlgu7JzslPjZGbNWKOZnBm7JUWrhlTASLpgD559gNR1nQm+UFD8fCaikULaoAVmVPeiRHC69kSztGoiy+QLNBt+Ggmz4zVDk+8S0yZPFjzLHfEgAA=="
    }

  }

```
内置建筑
Certainly! Here's a table with the requested columns extracted from the provided data:

| id | place_name      | top_left_point              | bottom_right_point           | center_point                 |
|----|-----------------|-----------------------------|------------------------------|------------------------------|
| 1  | 机电楼            | {"x":115.798058,"y":28.661722} | {"x":115.801668,"y":28.662683} | {"x":115.799863,"y":28.6622025} |
| 2  | 建工楼            | {"x":115.798964,"y":28.662739} | {"x":115.802559,"y":28.66329}  | {"x":115.8007615,"y":28.6630145}|
| 3  | 一食堂            | {"x":115.803739,"y":28.66482}  | {"x":115.804527,"y":28.664344} | {"x":115.804133,"y":28.664582} |
| 4  | 润溪湖（1）        | {"x":115.805922,"y":28.66449}  | {"x":115.804366,"y":28.665149} | {"x":115.805144,"y":28.6648195}|
| 5  | 润溪湖（3）        | {"x":115.80722,"y":28.662584}  | {"x":115.810235,"y":28.661492} | {"x":115.8087275,"y":28.662038}|
| 6  | 艺术楼            | {"x":115.807285,"y":28.661303} | {"x":115.808915,"y":28.659976} | {"x":115.8081,"y":28.660639500000002}|
| 7  | 外经楼            | {"x":115.805396,"y":28.661501} | {"x":115.80677,"y":28.66265}  | {"x":115.806083,"y":28.6620755}|
| 8  | 文法楼            | {"x":115.803744,"y":28.661209} | {"x":115.805332,"y":28.660023} | {"x":115.80453800000001,"y":28.660615999999997}|
| 9  | 正气广场          | {"x":115.805461,"y":28.659948} | {"x":115.807285,"y":28.658969} | {"x":115.806373,"y":28.6594585}|
| 10 | 白帆              | {"x":115.810444,"y":28.663709} | {"x":115.813459,"y":28.65967}  | {"x":115.81195149999999,"y":28.6616895}|
| 11 | 体育馆            | {"x":115.810755,"y":28.659557} | {"x":115.812515,"y":28.658682} | {"x":115.811635,"y":28.6591195}|
| 12 | 游泳馆            | {"x":115.810187,"y":28.663709} | {"x":115.812043,"y":28.664829} | {"x":115.811115,"y":28.664269}|
| 13 | 休闲运动场        | {"x":115.809919,"y":28.664886} | {"x":115.811689,"y":28.667747} | {"x":115.81080399999999,"y":28.6663165}|
| 14 | 休闲广场          | {"x":115.808256,"y":28.665747} | {"x":115.806823,"y":28.665968} | {"x":115.80753949999999,"y":28.6658575}|
| 15 | 休闲13栋         | {"x":115.806453,"y":28.665992} | {"x":115.807451,"y":28.666476} | {"x":115.806952,"y":28.666234}|
| 16 | 树人广场          | {"x":115.802779,"y":28.657034} | {"x":115.803723,"y":28.656224} | {"x":115.803251,"y":28.656629000000002}|
| 17 | 龙腾湖（1）        | {"x":115.800552,"y":28.656078} | {"x":115.802151,"y":28.654516} | {"x":115.8013515,"y":28.655297}|
| 18 | 龙腾湖（2）        | {"x":115.802655,"y":28.656088} | {"x":115.803996,"y":28.655382} | {"x":115.8033255,"y":28.655735}|
| 19 | 龙腾湖（3）        | {"x":115.805434,"y":28.657453} | {"x":115.806378,"y":28.657058} | {"x":115.805906,"y":28.657255499999998}|
| 20 | 龙腾湖（4）        | {"x":115.806217,"y":28.65863}  | {"x":115.809768,"y":28.657952} | {"x":115.80799250000001,"y":28.658291}|
| 21 | 天健运动场        | {"x":115.793815,"y":28.653009} | {"x":115.796679,"y":28.654817} | {"x":115.79524699999999,"y":28.653913000000003}|
| 22 | 医学实验大楼      | {"x":115.796829,"y":28.656088} | {"x":115.798213,"y":28.653216} | {"x":115.797521,"y":28.654652}|
| 23 | 研究生            | {"x":115.793461,"y":28.65299}  | {"x":115.79625,"y":28.651258} | {"x":115.7948555,"y":28.652124}|
