### 静态语言


### 高可扩展性



### hi i am golang home-page
![image](./assets/homepage.png)


build with docker 
expose port with  8080(mac) -> 80(docker)



### info
➜  app go version
go version go1.11.4 linux/amd64


### 多个文件如何运行
```go
go run app.go  key.go
```


### gopath
###### 学个Golang真不容易,终于弄懂了$GOPATH的用途,, 他可以设置多个路径啊,第一路径自然是全局路径,`$HOEM/golang/src/github`,第二路径是项目路径,用`;`隔开,,当设置好全局路径之后,目录下会生成三个目录, `src`源码source,`pkg`打包后的package,`bin`可执行命令
```
/home/user/go/
    src/
        crash/
            bang/              (go code in package bang)
                b.go
        foo/                   (go code in package foo)
            f.go
            bar/               (go code in package bar)
                x.go
            vendor/
                crash/
                    bang/      (go code in package bang)
                        b.go
                baz/           (go code in package baz)
                    z.go
            quux/              (go code in package main)
                y.go
```

> #### 用法
> cd $GOPATH => 你定义的环境变量
> ```
> import (
>   github.com/package   ->  which you git pull from github.com
> )
> ```
> go build github.com/<which path you install in gopath>


### goroot
就是go的安装路径