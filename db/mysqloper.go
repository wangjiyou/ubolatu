package db

import (
	"fmt"
	"time"

	"ubolatu/config"
	"ubolatu/pub"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserInfo pub.UserInfoRequest

var conf = config.GlobalConfig()
var orm *gorm.DB

func InitMysql() {
	var err error
	/*
	           "UserName":"root",
	           "Password":"123456",
	           "IPAddress":"10.31.2.212",
	           "PortAddress":"3306",
	   		"DBName":"Center",
	*/
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
	userInfo := UserInfo{}
	userInfo.NickName = "wangjiyou"
	orm.Save(&userInfo)

	rows, err := orm.Table("user_infos").Where("nick_name = 'wangjiyou'").Rows()
	if err != nil {
		fmt.Printf("Not error should happen, got %v", err)
		return
	}

	var results []UserInfo
	for rows.Next() {
		var result UserInfo
		if err := orm.ScanRows(rows, &result); err != nil {
			fmt.Printf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}
	fmt.Println("result:", results)

}
