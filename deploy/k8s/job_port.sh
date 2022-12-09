#!/bin/sh
case $1 in
"admin-api") echo 30000
;;
"identity-api") echo 30001
;;
"identity-rpc") echo 31001
;;
"common-rpc") echo 31002
;;
"mq-rpc") echo 31003
;;
"member-api") echo 30004
;;
"member-rpc") echo 31004
;;
"game-api") echo 30005
;;
"game-rpc") echo 31005
;;
"shortbook-api") echo 30006
;;
"shortbook-rpc") echo 31006
;;
"system-rpc") echo 31007
;;
"wallet-api") echo 30008
;;
"wallet-rpc") echo 31008
;;
"behavior-api") echo 30009
;;
"behavior-rpc") echo 31009
;;
"message-api") echo 30010
;;
"message-rpc") echo 31010
;;
"home-api") echo 30011
;;
"home-rpc") echo 31011
;;
esac