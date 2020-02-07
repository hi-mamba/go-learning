## [参考](https://github.com/hi-mamba/docker-learning/blob/master/02%E3%80%81docker%20%E5%AE%89%E8%A3%85%E3%80%81%E5%85%A5%E9%97%A8/12%E3%80%81docker%20java%20%E4%BE%8B%E5%AD%90.md)

# docker执行 go

```shell script
 cp helloworld.go /tmp 
 
 cd /tmp 

 docker run -v "$PWD":/usr/src/hello-docker -w /usr/src/hello-docker golang:latest go run helloworld.go
```

执行结果
```
Hello World
```

docker run -v "$PWD":/usr/src/hello-docker -w /usr/src/hello-docker golang:latest go run helloworld.go

> /usr/src/hello-docker 这个文件夹路径你需要注意， 这里应该是你 docker配置的 file sharing 里面设置的

