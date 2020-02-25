

# build command-line-arguments: cannot find module for path xxx

应该改成 import 模块名/包名

[例子](/go-learning-example/go-basic-example/002、package-variable-function/07、variables.go)

## 参考

[一次报错的 彻底明白 go的GOROOT ,GOPATH 以及go mod （cannot find module providing package bufio）](https://blog.csdn.net/sinat_23156865/article/details/100655475)

> 在go module下 你源码中 import …/ 这样的引入形式不支持了， 应该改成 import 模块名/包名  这样就ok了

[build command-line-arguments: cannot load XXX: cannot find module providing package Goland包导入问题解决](https://blog.csdn.net/slphahaha/article/details/99637047?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task)