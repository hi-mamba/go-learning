# goland-learn
Standing on Shoulders of Giants


# GOROOT、GOPATH和Project目录说明

### [原文](https://blog.csdn.net/zwqjoy/article/details/78788918)

# 可查看当前go环境变量
> go env

##  GOROOT
  GOROOT就是go的安装路径


## GOBIN
go install编译存放路径。不允许设置多个路径。可以为空。
为空时则遵循“约定优于配置”原则，可执行文件放在各自GOPATH目录的bin文件夹中
（前提是：package main的main函数文件不能直接放到GOPATH的src下面。

  GOPATH目录结构
  goWorkSpace  // (goWorkSpace为GOPATH目录)

```
    -- bin  // golang编译可执行文件存放路径，可自动生成。
    -- pkg  // golang编译的.a中间文件存放路径，可自动生成。
    -- src  // 源码路径。按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）。
````


## go目录结构1:
```
project1 // (project1添加到GOPATH目录了)
  -- bin
  -- pkg
  -- src
     -- models       // package
     -- controllers  // package
     -- main.go      // package main［注意，本文所有main.go均指包main的入口函数main所在文件］

 project2 // (project2添加到GOPATH目录了)
      -- bin
      -- pkg
      -- src
         -- models       // package
         -- controllers  // package
         -- main.go      // package main
```
package main文件直接在GOPATH目录到src下。

使用go build可以在src文件夹下编译生成名为“src”的可执行文件。这是golang默认约定。
一般我个人不怎么用这个命令。因为它会生成可执行文件在src目录下面。