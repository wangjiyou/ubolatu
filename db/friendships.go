package db

import (
	"encoding/json"
	"fmt"

	"ubolatu/pub"
)

type FriendShip pub.FriendShipRequest

func FindFriendShip(OwnerOpenId string, AddType string) ([]byte, error) {
	//rows, err := DB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("name, age").Rows()
	rows, err := ormTiDB.Table("friend_ships").Where("owner_id = ? and add_type = ?", OwnerOpenId, AddType).Select("friend_id, friend_name").Rows()
	if err != nil {
		fmt.Printf("Not error should happen, got %v", err)
		return []byte{}, err
	}

	var results []FriendShip
	for rows.Next() {
		var result FriendShip
		if err := ormTiDB.ScanRows(rows, &result); err != nil {
			fmt.Printf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}
	return json.Marshal(results)
}

func ModifyFriendShipAddType(OwnerOpenId string, FriendOpenID string, AddType string) error {
	err := ormTiDB.Model(FriendShip{}).Where(&FriendShip{OwnerID: OwnerOpenId, FriendID: FriendOpenID}).Update("AddType", AddType).Error
	if err != nil {
		fmt.Println("Unxpected error on conditional update err:", err)
		return err
	}

	return nil
}

func DeleteFriendShip(OwnerOpenId string, FriendOpenID string) error {
	err := ormTiDB.Where(FriendShip{OwnerID: OwnerOpenId, FriendID: FriendOpenID}).Delete(&FriendShip{}).Error
	if err != nil {
		fmt.Println("delete owneropenId:", OwnerOpenId, "friendID:", FriendOpenID, " err:", err)
		return err
	}
	return nil
}

func SetFriendShip(_info pub.FriendShipRequest) {
	info := FriendShip(_info)
	ormTiDB.Save(&info)
}

func CreateFriendShips() bool {
	ok := ormTiDB.HasTable(&FriendShip{})
	if ok {
		fmt.Println("FriendShips Table exist")
		TidbTest()
		return true
	}
	fmt.Println("Table should exist, but HasTable friendships it does not")
	if err := ormTiDB.CreateTable(&FriendShip{}).Error; err != nil {
		fmt.Println("created err:", err)
		return false
	}
	fmt.Println("create friendships OK")
	TidbTest()
	return true
}

func TidbTest() {
	friendShip := FriendShip{}
	friendShip.FriendName = "nickname_test"
	ormTiDB.Save(&friendShip)

	rows, err := ormTiDB.Table("friend_ships").Where("friend_name = 'nickname_test'").Rows()
	if err != nil {
		fmt.Printf("Not error should happen, got %v", err)
		return
	}

	var results []FriendShip
	for rows.Next() {
		var result FriendShip
		if err := ormTiDB.ScanRows(rows, &result); err != nil {
			fmt.Printf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}
	fmt.Println("friendships result:", results)

	err = ormTiDB.Model(FriendShip{}).Where(&FriendShip{FriendName: "nickname_test"}).Update("NickName", "hahah").Error
	if err != nil {
		fmt.Println("Unxpected error on conditional update err:", err)
		return
	}
	fmt.Println("Update OK")

	err = ormTiDB.Where(FriendShip{FriendName: "nickname_test"}).Delete(&FriendShip{}).Error
	if err != nil {
		fmt.Println("Unexpected error on conditional delete")
		return
	}
	fmt.Println("Delete OK")

}
