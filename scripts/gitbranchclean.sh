#!/bin/bash

#将git分支记录清空

#分支
GitBranch="$1"

#Git 代码分支检查
GitBranchCheck() {
  if [ "$GitBranch" == "" ]; then
    Log "请传入参数:代码分支 usage: $0 {branch}" "red"
    exit
  fi
  cd "$CodePath" || exit
  if [ "$(git rev-parse --verify "$GitBranch" 2>/dev/null)" == "" ]; then
    Log "代码分支不存在" "red"
    exit
  fi
}




Log() {
  local RED_COLOR='\033[31m'    #红
  local GREEN_COLOR='\033[32m'  #绿
  local YELLOW_COLOR='\033[33m' #黄
  local BLUE_COLOR='\033[34m'   #蓝
  local RES_COLOR='\033[0m'

  case "$2" in
  'red')
    echo -e "$RED_COLOR*** $1 ***$RES_COLOR"
    echo -e "$RED_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'green')
    echo -e "$GREEN_COLOR*** $1 ***$RES_COLOR"
    echo -e "$GREEN_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'yellow')
    echo -e "$YELLOW_COLOR*** $1 ***$RES_COLOR"
    echo -e "$YELLOW_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  'blue')
    echo -e "$BLUE_COLOR*** $1 ***$RES_COLOR"
    echo -e "$BLUE_COLOR*** $1 ***$RES_COLOR" >>"$LogFile"
    ;;
  *)
    echo -e "*** $1 ***" >>"$LogFile"
    ;;
  esac

}

main(){
   GitBranchCheck
   echo "Git分支:$GitBranch记录清空开始"
   git checkout --orphan latest_branch
   git add -A
   git commit -am "clean"
   git branch -D "$GitBranch"
   git branch -m "$GitBranch"
   git push -f origin "$GitBranch"
   echo "Git分支:$GitBranch记录清空结束"
}

main