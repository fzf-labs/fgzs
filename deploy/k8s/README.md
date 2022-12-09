# k8s 部署相关

##1.k8s创建find-endpoints这个serviceAccount并绑定相应权限
```shell
kubectl apply -f auth.yaml 
```
##2.配置k8s拉取私有仓库镜像
k8s在默认情况下，只能拉取harbor镜像仓库的公有镜像，如果拉取私有仓库镜像，则是会报 `ErrImagePull` 和 `ImagePullBackOff` 的错误

1、先在jenkins发布机器登陆harbor

```shell
$ docker login 120.77.97.206:8077
$ Username: admin
$ Password:
Login Succeeded
```

2、在k8s中生成登陆harbor配置文件

```shell
#查看上一步登陆harbor生成的凭证
$ cat /root/.docker/config.json  
{
	"auths": {
		"120.77.97.206:8077": {
			"auth": "YWRtaW46YmxramFsaWxp"
		}
}
```

3、对秘钥文件进行base64加密

```shell
$ cat /root/.docker/config.json  | base64 -w 0

ewoJImF1dGhzIjogewoJCSIxMjAuNzcuOTcuMjA2OjgwNzciOiB7CgkJCSJhdXRoIjogIllXUnRhVzQ2WW14cmFtRnNhV3hwIgoJCX0KCX0KfQo=
```

4、创建docker-secret.yaml

```yml
apiVersion: v1
kind: Secret
metadata:
  name: docker-login
type: kubernetes.io/dockerconfigjson
data:
  .dockerconfigjson: ewoJImF1dGhzIjogewoJCSIxMjAuNzcuOTcuMjA2OjgwNzciOiB7CgkJCSJhdXRoIjogIllXUnRhVzQ2WW14cmFtRnNhV3hwIgoJCX0KCX0KfQ=

```

```shell
$ kubectl create -f docker-secret.yaml -n fgzs

secret "docker-login" created
```



