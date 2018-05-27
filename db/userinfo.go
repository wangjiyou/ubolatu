package db

import (
	"fmt"

	"ubolatu/pub"
)

type UserInfo pub.UserInfoRequest

func DeleteUserInfo(openId string) error {
	err := orm.Where(UserInfo{OpenID: openId}).Delete(&UserInfo{}).Error
	if err != nil {
		fmt.Println("delete openId:", openId, " err:", err)
		return err
	}
	return nil
}

func SetUserInfo(_info pub.UserInfoRequest) {
	info := UserInfo(_info)
	orm.Save(&info)
}

func IsExistOpenID(openId string) bool {
	_, err := orm.Table("user_infos").Where("open_id = ?", openId).Rows()
	if err != nil {
		fmt.Printf("Not error should happen, got %v", err)
		return false
	}
	return true
}

func SetSessionKey(openId string, sessionKey string) {
	err := orm.Model(UserInfo{}).Where(&UserInfo{OpenID: openId}).Update("NickName", "hahah").Error
	if err != nil {
		fmt.Println("Unxpected error on conditional update err:", err)
		return
	}
	fmt.Println("Update OK")
}

func GetSessionKey(openId string) string {
	//err := orm.Model(UserInfo{}).Where(&UserInfo{OpenID: openId}).Update("NickName", "hahah").Error
	rows, err := orm.Table("user_infos").Where("open_id = '%s'", openId).Rows()
	if err != nil {
		fmt.Printf("GetSessionKey got %v", err)
		return ""
	}

	var results []UserInfo
	for rows.Next() {
		var result UserInfo
		if err := orm.ScanRows(rows, &result); err != nil {
			fmt.Printf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}
	return results[0].SessionKey
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
	userInfo.NickName = "nickname_test"
	orm.Save(&userInfo)

	rows, err := orm.Table("user_infos").Where("nick_name = 'nickname_test'").Rows()
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

	err = orm.Model(UserInfo{}).Where(&UserInfo{NickName: "nickname_test"}).Update("NickName", "hahah").Error
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
