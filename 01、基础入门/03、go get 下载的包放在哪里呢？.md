
## [原文](https://www.cnblogs.com/oxspirt/p/9340250.html)

# go get 下载的包放在哪里呢？

有些问题，我以前都是似懂非懂，没有去弄个究竟!!!!!

这个习惯非常不好，搞得有些东西看似懂了，又不能百分之百说自己懂了，可能下次就弄不出来了，
这样是不可取的。 不能有这种做事的风格。

```shell script
go env
```
```
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/kobe/Library/Caches/go-build"
GOENV="/Users/kobe/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/kobe/go"     //文件下载到这里,具体的是在 pkg/mod/golang.org 下【注意go的版本，我的是1.13.8】

省略..
```

就像写go的代码的时候， 要引用第三方库的时候， go get github.com/go-redis/redis 这样就ok了，
但是这个命令到底干了什么？怎么干的？下载的文件到底在哪里？

这些问题，好像是很模糊的。这个时候，就不能允许这种模糊一直持续下去。
否则的话，就是一种敷衍的态度，一种不思进取的态度。






