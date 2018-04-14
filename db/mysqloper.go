package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
type UserInfo struct {
	OpenID          string `json:"openId"`
	NickName        string `json:"nickName"`
	Gender          string `json:"gender"`
	City            string `json:"city"`
	Province        string `json:"province"`
	Country         string `json:"country"`
	AvatarURL       string `json:"avatarUrl"`
	UnionID         string `json:"unionId"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Timestamp       string `json:"timestamp"`
}
*/
var orm *gorm.DB

func InitMysql() {
	/*
		var err error
		orm, err = gorm.Open("mysql", "root:ying@tcp(127.0.0.1:3306)/center?charset=utf8")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("connect OK")
	*/

	//var orm *gorm.DB
	var err error

	dbHost := "127.0.0.1"
	dbName := "center"
	dbUser := "root"
	dbPasswd := "ying"
	dbPort := "3306"
	dbType := "mysql"

	connectString := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	//开启sql调试模式
	//GDB.LogMode(true)

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

func Find() {

	/*
		user1 := User{Name: "ScanRowsUser1", Age: 1, Birthday: parseTime("2000-1-1")}
		user2 := User{Name: "ScanRowsUser2", Age: 10, Birthday: parseTime("2010-1-1")}
		user3 := User{Name: "ScanRowsUser3", Age: 20, Birthday: parseTime("2020-1-1")}
		DB.Save(&user1).Save(&user2).Save(&user3)

		rows, err := DB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("name, age").Rows()
		if err != nil {
			t.Errorf("Not error should happen, got %v", err)
		}

		type Result struct {
			Name string
			Age  int
		}

		var results []Result
		for rows.Next() {
			var result Result
			if err := DB.ScanRows(rows, &result); err != nil {
				t.Errorf("should get no error, but got %v", err)
			}
			results = append(results, result)
		}

		if !reflect.DeepEqual(results, []Result{{Name: "ScanRowsUser2", Age: 10}, {Name: "ScanRowsUser3", Age: 20}}) {
			t.Errorf("Should find expected results")
		}
	*/
}
