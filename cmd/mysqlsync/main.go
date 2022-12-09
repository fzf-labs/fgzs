package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type MysqlDsn struct {
	User string
	Pass string
	Host string
}

func main() {
	source := MysqlDsn{
		User: "root",
		Pass: "evHpuc61vrd1tWlS",
		Host: "120.24.89.142",
	}
	target := MysqlDsn{
		User: "root",
		Pass: "123456",
		Host: "127.0.0.1",
	}
	database := []string{
		"fgzs_behavior",
		"fgzs_common",
		"fgzs_file",
		"fgzs_game",
		"fgzs_home",
		"fgzs_member",
		"fgzs_message",
		"fgzs_shortbook",
		"fgzs_system",
		"fgzs_wallet",
	}
	//mysqldump fgzs_test -uroot -pevHpuc61vrd1tWlS -h120.24.89.142  --add-drop-table | mysql fgzs_test -uroot -p123456 -h127.0.0.1
	if len(database) > 0 {
		for _, v := range database {
			do(source, target, v)
		}
	}

}

func do(source MysqlDsn, target MysqlDsn, database string) {
	s := fmt.Sprintf("mysqldump %s -u%s -p%s -h%s  --add-drop-table | mysql %s -u%s -p%s -h%s", database, source.User, source.Pass, source.Host, database, target.User, target.Pass, target.Host)
	cmd := exec.Command("bash", "-c", s)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
