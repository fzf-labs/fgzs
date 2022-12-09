package main

import (
	"database/sql"
	"fgzs/pkg/conv"
	"fgzs/pkg/util/fileutil"
	"fgzs/pkg/util/strutil"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"path"
	"regexp"
	"strings"
)

var FileName = flag.String("f", "", "the config file")

// SHOW CREATE TABLE `fgzs_shortbook`.`shortbook_comment_reply_like`;
// SELECT table_name FROM information_schema.tables WHERE table_schema='fgzs_shortbook';
func main() {
	flag.Parse()
	dsn := GetDsn(*FileName)
	database := regexp.MustCompile(`/(.*?)\?`).FindString(dsn)
	database = strings.TrimLeft(database, "/")
	database = strings.TrimRight(database, "?")
	//连接数据库
	db := ConnectDB(dsn)
	defer db.Close()
	var tables []string
	//查所有的table
	rows, err := db.Query(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema='%s'", database))
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return
		}
		tables = append(tables, table)
	}
	type Result struct {
		Table       string `json:"table"`
		CreateTable string `json:"create_table"`
	}
	str := ""
	for _, v := range tables {
		var res Result
		err := db.QueryRow(fmt.Sprintf("SHOW CREATE TABLE %s", v)).Scan(&res.Table, &res.CreateTable)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res.CreateTable != "" {
			str += res.CreateTable + ";\n"
		}
	}
	target, _ := strutil.SubstrTarget(*FileName, "/app", "left", false)
	p := target + "/storage/sql/" + database + ".sql"
	err = fileutil.WriteContentCover(p, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SQL CREATE TABLE 导出成功")
}

func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return db
}

func GetDsn(fileName string) string {
	fileDir := path.Dir(fileName)
	filePathBase := path.Base(fileName)
	fileExt := path.Ext(fileName)
	filePrefix := filePathBase[0 : len(filePathBase)-len(fileExt)]
	config := viper.New()
	config.AddConfigPath(fileDir)    //设置读取的文件路径
	config.SetConfigName(filePrefix) //设置读取的文件名
	config.SetConfigType("yaml")     //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("Mysql.DataSourceName"))
}
