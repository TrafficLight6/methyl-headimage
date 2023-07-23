# Methyl-headimage

__By TrafficLight6__

![LOGO](./logo.png) 

一个api服务程序，供给头像服务，可以修改为图床

## 下载与部署
- 下载
    用git clone或wget下载，再不行可以有scp协议上传目标服务器
- 部署
    - 使用源代码部署
    - 安装go与mysSQL
    - ~~安装依赖beego~~ 貌似不需要
    - 安装go第三方库
        ```
        go get methyl-headimage
        ```
    - 初始化MySQL
        打开controllers/initsql.go，在13行：
        ```go
        database, err := sqlx.Open("mysql", "headimg:headimg@tcp(127.0.0.1:3306)/headimg")
        //                                   用户名  数据库密码    主机及其端口    数据库名称
        ```
        进入MySQL，把headimg.sql导入
    - 修改文件上传大小上限（可选）
        打开controllers/web.go，在45行
        ```go
        if head.Size > 5000000 {    //看到这个5000000，则就是大小上限（单位为byte）
		c.Ctx.WriteString("{'code':400,'massage':'file is too large'}")
	    }
        ```
    - 运行main.go

## api请求
__所有请求返回格式均为json__
- /upload
    - 请求方法：POST
    - 请求说明：上传图片
    - 请求参数：
        - userid，数字，上传的图片所属用户
        - uploadfile，文件，上传的文件
    - 返回实例：{'code':200,'massage':'successfully'}
----------
- /getpath/all
    - 请求方法：GET
    - 请求说明：获取某个用户的图片
    - 请求参数：
        - userid，数字，用户id
    - 返回实例：{'code':200,'massage':["path1","path2"……]}
----------
- /getpath/inuse
    - 请求方法：GET
    - 请求说明：获取某个用户的正在使用图片（仅作用于头像应用）
    - 请求参数：
        - userid，数字，用户id
    - 返回实例：{'code':200,'massage':["path1"]}
----------
- /useimg
    - 请求方法：PUT
    - 请求说明：通过图片路径索引，使用某个用户的图片（仅作用于头像应用）
    - 请求参数：
        - userid，数字，用户id
        - imgpath，字符串，图片路径
    - 返回实例：{'code':200,'massage':'successfully'}

## 更新日志
    - beta 0.1.0，2023年7月23日
        第一个版本

更新于2023年7月23日