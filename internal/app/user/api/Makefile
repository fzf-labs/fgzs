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

#GIT 分支
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
#GIT commit_id 短
GIT_COMMIT_ID = $(shell git rev-parse HEAD)
#服务端口
JOB_PORT = 30001
#镜像地址
HARBOR_REPO = "120.77.97.206:8077"
HARBOR_IMAGE = ${JOB_NAME}-${JOB_TYPE}:${GIT_COMMIT_ID}

.PHONY: go-api
go-api: swagger
	@goctl api go -api ${JOB_NAME}.api -dir ./ -home ../../../../deploy/goctl

.PHONY: swagger
swagger:
	@goctl api plugin -plugin goctl-swagger="swagger -filename ${JOB_NAME}.json" -api ${JOB_NAME}.api -dir ../../../../storage/swagger

.PHONY: go-run
# go run
go-run:
	@go run *.go -f ${BASEDIR}/etc/${JOB_NAME}.yaml

.PHONY: dockerfile
#dockerfile 配置文件生成
dockerfile:
	@rm -f Dockerfile
	@goctl docker -base alpine -go *.go

.PHONY: k8s-yaml
# k8s 配置文件生成
k8s-yaml:
	@rm -f k8s.yaml
	@goctl kube deploy -secret docker-login -replicas 2 -nodePort ${JOB_PORT} -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name ${JOB_NAME}-${JOB_TYPE} -namespace ${PROJECT_NAME} -image ${PROJECT_NAME}/${HARBOR_REPO}/${HARBOR_IMAGE} -o k8s.yaml -port ${JOB_PORT} -serviceAccount find-endpoints