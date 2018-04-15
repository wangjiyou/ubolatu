package db

import (
	"fmt"
	//"time"

	//"ubolatu/config"
	"ubolatu/pub"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
)

type UserInfo pub.UserInfoRequest

func IsExistOpenID(openId string) bool {
	_, err := orm.Table("user_infos").Where("open_id = ?", openId).Rows()
	if err != nil {
		fmt.Printf("Not error should happen, got %v", err)
		return false
	}
	return true
}

func UpdateSessionKey(openId string, sessionKey string) {
	err := orm.Model(UserInfo{}).Where(&UserInfo{OpenID: openId}).Update("NickName", "hahah").Error
	if err != nil {
		fmt.Println("Unxpected error on conditional update err:", err)
		return
	}
	fmt.Println("Update OK")
}

func CreateUserInfos() bool {
	ok := orm.HasTable(&UserInfo{})
	if ok {
		fmt.Println("Table exist")
		ISUDTest()
		return true
	}
	fmt.Println("Table should exist, but HasTable informs it does not")
	if err := orm.CreateTable(&UserInfo{}).Error; err != nil {
		fmt.Println("created err:", err)
		return false
	}
	fmt.Println("create userinfos OK")
	ISUDTest()
	return true
}

func ISUDTest() {
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

	err = orm.Model(UserInfo{}).Where(&UserInfo{NickName: "wangjiyou"}).Update("NickName", "hahah").Error
	if err != nil {
		fmt.Println("Unxpected error on conditional update err:", err)
		return
	}
	fmt.Println("Update OK")

	err = orm.Where(UserInfo{NickName: "hahah"}).Delete(&UserInfo{}).Error
	if err != nil {
		fmt.Println("Unexpected error on conditional delete")
		return
	}
	fmt.Println("Delete OK")

}
