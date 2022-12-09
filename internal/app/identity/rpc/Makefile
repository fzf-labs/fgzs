SHELL := /bin/bash
#当前文件夹
BASEDIR = $(shell pwd)
#项目名称
PROJECT_NAME := "fgzs"
#服务类型
JOB_TYPE = $(shell basename "${BASEDIR}")
#服务名
JOB_NAME = $(shell echo "${BASEDIR}"  | rev |cut -d '/' -f 2 | rev)
# PROTO文件
JOB_PROTO = $(shell find . -name *.proto)
# PROTO文件夹
JOB_PROTO_DIR = $(basename ${JOB_PROTO})
#数据库Dsn
DataSourceName= $(shell (cat ${BASEDIR}/etc/${JOB_NAME}.yaml | grep "DataSourceName" | xargs echo) | cut -d ' ' -f 2 | cut -d '?' -f 1)
#GIT 分支
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
#GIT commit_id 短
GIT_COMMIT_ID = $(shell git rev-parse HEAD)
#服务端口
JOB_PORT = $(shell (cat ${BASEDIR}/etc/${JOB_NAME}.yaml | grep "ListenOn" | xargs echo) | cut -d ':' -f 3)
#镜像地址
HARBOR_REPO = "120.77.97.206:8077"
#镜像名称
HARBOR_IMAGE = ${JOB_NAME}-${JOB_TYPE}:${GIT_COMMIT_ID}

.PHONY: proto
# 生成 proto 文件
proto:
	@goctl rpc protoc ${JOB_PROTO}  --go_out=./ --go-grpc_out=./  --zrpc_out=./ -v -home ../../../../deploy/goctl -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
	@protoc ${JOB_PROTO}  --go_out=./ --go-grpc_out=./ --validate_out=lang=go:./ -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
.PHONY: omitempty
#  proto文件中的 omitempty 删除
omitempty:
	@sed -i "" 's/,omitempty//g' *.pb.go
.PHONY: go-run
# go run
go-run:
	@go run *.go -f ${BASEDIR}/etc/${JOB_NAME}.yaml

.PHONY: model
model:
	@goctl model mysql datasource -url="${DataSourceName}" -table="*" -c -dir ../model -home ../../../../deploy/goctl
	@go run ../../../../cmd/sqldump/main.go -f ${BASEDIR}/etc/${JOB_NAME}.yaml

.PHONY: dockerfile
#dockerfile 配置文件生成
dockerfile:
	@rm -f Dockerfile
	@goctl docker -home ../../../../deploy/goctl -base alpine -go *.go

.PHONY: k8s-yaml
# k8s 配置文件生成
k8s-yaml:
	@rm -f k8s.yaml
	@goctl kube deploy -home ../../../../deploy/goctl -secret docker-login -replicas 2 -nodePort ${JOB_PORT} -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name ${JOB_NAME}-${JOB_TYPE} -namespace ${PROJECT_NAME} -image ${PROJECT_NAME}/${HARBOR_REPO}/${HARBOR_IMAGE} -o k8s.yaml -port ${JOB_PORT} -serviceAccount find-endpoints
