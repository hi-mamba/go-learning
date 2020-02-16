
## [参考](https://www.cnblogs.com/yjf512/p/10482604.html)

## [参考](https://github.com/fenggolang/golang/blob/master/%E8%A7%A3%E5%86%B3go%20get%E6%97%A0%E6%B3%95%E4%B8%8B%E8%BD%BD%E8%A2%AB%E5%A2%99%E7%9A%84%E5%8C%85.md)

# go get golang.org被墙问题解决

## 1、使用gopm(Go Package Manager)代替go下载,是go上的包管理工具，十分好用

1. 下载安装gopm

推荐使用 官方出品：<https://github.com/golang/dep>

> go get -u github.com/gpmgo/gopm

2. 使用gopm安装被墙的包

> gopm get github.com/Shopify/sarama

## 2、golang 在 github 上建立了一个镜像库

如 https://github.com/golang/net 即是 https://golang.org/x/net 的镜像库,
获取 golang.org/x/net 包（其他包类似），其实只需要以下步骤：

```shell script
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
```