### 选择Golang的原因
- [ ] 静态语言
- [ ] 高可扩展性
- [ ] 它将`github.com`当做远程包扩展进去了,适应了现代化开发
- [ ] 时代在变化,网络基础设施也在变化,过去式语言/python/js+jquery+bootstrapt/也许已经不适用与现代开发,例如docker容器开发,而java作为一门失败的语言,过于复杂了.



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


### 字符串 - fmt
> #### 解释型字符串
> ```go
> var str = "string is \n me" // 打印成换行符
> ```

> #### 非解释型字符串
> ```go
> var str = `string is \n me` // 原样输出
> ```


### 打印
> #### `fmt.Println`
> Println formats using the default formats for its operands and writes to standard output.
> Spaces are always added between operands and a newline is appended.
> It returns the number of bytes written and any write error encountered.
> Println formats using the default formats for its operands and writes to standard output.
> `Println`格式使用其操作数的默认格式并写入标准输出。始终在操作数之间添加空格，并附加换行符。它返回写入的字节数和遇到的任何写入错误。Println格式使用其操作数的默认格式并写入标准输出。
> ```go
> func Println(a ...interface{}) (n int, err error) {
> 	return Fprintln(os.Stdout, a...)
> }
> ```
> #### 总结 
> 例如: 如果不关心格式 我们只想看值的时候,或者仅仅用于调试,用`Println`就比较合适



> #### `fmt.Printf`
> Printf formats according to a format specifier and writes to standard output.
> It returns the number of bytes written and any write error encountered.
> `Printf`  根据格式, 输出标准化的格式化输出, 它返回写入的数据大小, 并且返回任何错误
> ```go
> func Printf(format string, a ...interface{}) (n int, err error) {
>	return Fprintf(os.Stdout, format, a...)
> }
> ```
> #### 总结
> 格式化输出 printf("sad %f", 1.0) -> 常常这么用


### 字符类型
> #### byte 
> (实际上是uint8的别名), 代表`UTF-8`字符串的单个字节的值
> ```go
> str3 := "啊hello 世界"
> str3Arr := []byte(str3)
> str3Arr[0] = '1'
> fmt.Println(string(str3Arr))
> ```
> 输出乱码,为什么,因为byte是utf-8单个字符,一个byte = 8个字节, 他只能代表英文 ASCII 字符表示方式


> #### rune
> (实际上是Unicode的别名), 代表 单个Unicode字符
> ```go
> str3 := "啊hello 世界"
> str3Arr := []byte(str3)
> str3Arr[0] = '1'
> fmt.Println(string(str3Arr))
> ```
> 成功替换,因为rune是Unicode单个字符,一个byte = 8个字节, 他只能代表英文 ASCII 字符表示方式




###### int空值
```go
var a int
fmt.Println(a)   // a -> 0  空值 init 就是 0
```


### Golang 中的指针
```go
var foo int
var bar *int 
fmt.Printf("%v %T \n", foo, foo)
fmt.Printf("%v %T \n", *bar, bar)
```






### Golang 中的verb
| verb | 描述 |
|:-:|:-:|
| %d | 十进制整数 |
| %x %o %b | 16进制整数,8进制整数,2进制整数 |
| %f %g %e | 浮点数: 3.141593 - 3.141592623422 - 3.141593e+00 |
| %t | 布尔型: true/false |
| %c | 字符(Unicode码点) |
| %s | 字符串 |
| %q | 带"的字符串("abc")或者字符(如:'c') |
| %v | 内置格式的任何值 |
| %T | 任何值的类型 |
| %% | 百分号本身,无操作 |




  

### 吐槽点
- [ ] 感觉golang的包管理工具真的超级弱智,不适合我这种小白开发人员