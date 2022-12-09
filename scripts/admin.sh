#!/bin/bash

SERVER="$1"
COMMAND="$2"
# 命令行参数
ARGS=""
BASE_DIR=$PWD
INTERVAL=2


FileName(){
   echo ${SERVER##*/}
}
Dir(){
   echo ${SERVER%/*}
}

Pid() {
 ps axo pid,command | awk "{if (\$3==\"$(FileName)\") print \$1}"
#  pidof "$SERVER"
}

function start() {
  if [ "$(Pid)" != "" ]; then
    echo "$SERVER already running!!!"
    exit 1
  fi
  echo "$SERVER begin start..."
  cd "$(Dir)"
  nohup ./"$(FileName)" "$(FileName)" >"$(FileName)"".nohup" 2>&1 &
  echo "sleeping..." && sleep $INTERVAL
  # check status
  if [ "$(Pid)" == "" ]; then
    echo "$SERVER start failed!!!"
    exit 1
  fi
  echo "$SERVER is running..."
}

function status() {
  if [ "$(Pid)" != "" ]; then
    echo "$SERVER" is running
  else
    echo "$SERVER" is not running
  fi
}

function stop() {
  if [ "$(Pid)" != "" ]; then
    kill -9 "$(Pid)"
  fi

  echo "sleeping..." && sleep $INTERVAL

  if [ "$(Pid)" != "" ]; then
    echo "$SERVER stop failed"
    exit 1
  fi
}

function main() {
  if [ ! -f "$SERVER" ]; then
    echo "文件不存在"
    exit 1
  fi

  case "$COMMAND" in
  'start')
    start
    ;;
  'stop')
    stop
    ;;
  'status')
    status
    ;;
  'restart')
    stop && start
    ;;
  *)
    echo "参数错误:启动方式必传 usage: $0 {file} {start|stop|restart|status}"
    exit 1
    ;;
  esac
}

main
