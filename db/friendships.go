package db

import (
	"fmt"

	"ubolatu/pub"
)

type FriendShip pub.FriendShipRequest

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

	err = ormTiDB.Where(FriendShip{FriendName: "hahah"}).Delete(&FriendShip{}).Error
	if err != nil {
		fmt.Println("Unexpected error on conditional delete")
		return
	}
	fmt.Println("Delete OK")

}
