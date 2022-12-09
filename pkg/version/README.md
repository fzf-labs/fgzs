## go编译文件记录git版本信息

git版本信息编译到go二进制文件中

```shell
go build -v -ldflags ${ldflags}
```