package db

import (
	"fmt"
	"time"

	"ubolatu/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var conf = config.GlobalConfig()
var orm *gorm.DB

func InitMysql() {
	var err error
	dbHost := conf.IPAddress
	dbName := conf.DBName      //"center"
	dbUser := conf.UserName    //"root"
	dbPasswd := conf.Password  //"ying"
	dbPort := conf.PortAddress //"3306"
	dbType := "mysql"

	connectString := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	//开启sql调试模式
	//GDB.LogMode(true)
	fmt.Println("connectstring:", connectString)

	for orm, err = gorm.Open(dbType, connectString); err != nil; {
		fmt.Println("数据库连接异常! 5秒重试")
		fmt.Println("err:", err)
		time.Sleep(5 * time.Second)
		orm, err = gorm.Open(dbType, connectString)
	}
	fmt.Println("connect OK")
	CreateUserInfos()
}
