package userinfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/xlstudio/wxbizdatacrypt"

	"ubolatu/client"
	"ubolatu/db"
	"ubolatu/pub"
)

type LoginSession struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

// StringService provides operations on strings.
type StringService interface {
	SetUserInfo(pub.FullUserInfo) (string, error)
	Count(string) int
	OnLogin(pub.LoginRequest) (string, int)
	AddFriend(pub.FriendShipRequest) (string, error)
	DelFriend(pub.FriendShipRequest) (string, error)
	ModiFriend(pub.FriendShipRequest) (string, error)
	FindFriend(pub.FriendShipRequest) (string, error)
	FindFans(pub.FriendShipRequest) (string, error)
}

type UstringService struct{}

func (UstringService) FindFans(request pub.FriendShipRequest) (string, error) {
	request.CreateAt = time.Now().String()
	fmt.Println("FindFans request:", request)
	content, err := db.FindFansShip(request.FriendID, request.AddType)
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}

func (UstringService) FindFriend(request pub.FriendShipRequest) (string, error) {
	request.CreateAt = time.Now().String()
	fmt.Println("FindFriend request:", request)
	content, err := db.FindFriendShip(request.OwnerID, request.AddType)
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}

func (UstringService) ModiFriend(request pub.FriendShipRequest) (string, error) {
	request.CreateAt = time.Now().String()
	fmt.Println("ModiFriend request:", request)
	err := db.ModifyFriendShipAddType(request.OwnerID, request.FriendID, request.AddType)
	return "", err
}

func (UstringService) DelFriend(request pub.FriendShipRequest) (string, error) {
	/*
		err, session := GetSession(request.Code)
		if err != nil {
			return err.Error(), http.StatusBadGateway
		}
		return session.Openid, http.StatusOK
	*/
	request.CreateAt = time.Now().String()
	fmt.Println("DelFriend request:", request)
	err := db.DeleteFriendShip(request.OwnerID, request.FriendID)
	return "", err
}

func (UstringService) AddFriend(request pub.FriendShipRequest) (string, error) {
	/*
		err, session := GetSession(request.Code)
		if err != nil {
			return err.Error(), http.StatusBadGateway
		}
		return session.Openid, http.StatusOK
	*/
	request.CreateAt = time.Now().String()
	fmt.Println("AddFriend request:", request)
	db.SetFriendShip(request)
	return "", nil
}

func (UstringService) SetUserInfo(request pub.FullUserInfo) (string, error) {
	fmt.Println("SetUserInfo full userinfo:", request)
	if request.OpenID == "" {
		fmt.Println("OpenID is nill")
		return "OpenID is nill", errors.New("OpenID is nill")
	}
	userInfo := pub.UserInfoRequest{}
	userInfo.OpenID = request.OpenID
	userInfo.NickName = request.UserInfo.NickName
	userInfo.Gender = fmt.Sprintf("%v", request.UserInfo.Gender)
	userInfo.AvatarURL = request.UserInfo.AvatarUrl
	userInfo.City = request.UserInfo.City
	userInfo.Province = request.UserInfo.Province
	userInfo.Country = request.UserInfo.Country
	userInfo.UnionID = ""
	userInfo.PhoneNumber = ""
	userInfo.Timestamp = time.Now().String()
	if db.IsExistOpenID(request.OpenID) {
		db.DeleteUserInfo(request.OpenID)
	}

	db.SetUserInfo(userInfo)
	//return strings.ToUpper(request.NickName), nil
	return "", nil
}

func (UstringService) OnLogin(request pub.LoginRequest) (string, int) {
	//fmt.Println("onlogin:", request.Code)
	//get openid and serectkey

	err, session := GetSession(request.Code)
	if err != nil {
		return err.Error(), http.StatusBadGateway
	}
	//	if db.IsExistOpenID(session.Openid) {
	//		db.SetSessionKey(session.Openid, session.SessionKey)
	//		return session.Openid, http.StatusOK
	//	}

	return session.Openid, http.StatusOK //http.StatusNoContent
}

func GetSession(code string) (error, LoginSession) {
	//LoginWeiXinServerUrl = "https://api.weixin.qq.com/sns/jscode2session?
	//appid=APPID&secret=APPSECRET&js_code=JSCODE&grant_type=authorization_code"
	session := LoginSession{}
	value := url.Values{}
	value.Set("appid", TagAppId)
	value.Set("secret", TagAppSecret)
	value.Set("js_code", code)
	value.Set("grant_type", TagAuthCodeFlag)
	loginUrl := fmt.Sprintf("%s?%s", TagLoginWeiXinServerUrl, value.Encode())
	err, body := client.HttpDo("GET", loginUrl, []byte(""))
	if err != nil {
		fmt.Println("get url:", loginUrl, " err:", err)
		return err, session
	}
	fmt.Println("get url:", loginUrl, "session:", string(body))
	err = json.Unmarshal(body, &session)
	if err != nil {
		fmt.Println("unmarshal body ", string(body), " err:", err)
		return err, session
	}
	fmt.Println("session:", session)
	return nil, session
}

func (UstringService) Count(s string) int {
	fmt.Println("userinfo:", s)
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

func EncryptedTest(appID string, sessionKey string, encryptedData string, iv string) {
	//appID := "wx4f4bc4dec97d474b"
	//sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	//encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	//iv := "r7BXXKkLb8qrSNn05n0qiA=="
	fmt.Println(sessionKey)
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appID, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		fmt.Println("EncryptedTest err:", err)
	} else {
		fmt.Println("EncryptedTest", result)
	}
}
