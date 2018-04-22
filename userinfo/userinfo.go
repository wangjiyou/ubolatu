package userinfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
	SetUserInfo(pub.UserInfoRequest) (string, error)
	Count(string) int
	OnLogin(pub.LoginRequest) (error, int)
}

type UstringService struct{}

func (UstringService) SetUserInfo(request pub.UserInfoRequest) (string, error) {
	fmt.Println("SetUserInfo NickName:", request.NickName)
	db.SetUserInfo(request)
	return strings.ToUpper(request.NickName), nil
}

func (UstringService) OnLogin(request pub.LoginRequest) (error, int) {
	fmt.Println("onlogin:", request.Code)
	//get openid and serectkey

	err, session := GetSession(request.Code)
	if err != nil {
		return err, http.StatusBadGateway
	}
	if db.IsExistOpenID(session.Openid) {
		db.UpdateSessionKey(session.Openid, session.SessionKey)
		return nil, http.StatusOK
	}

	return nil, http.StatusNoContent
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

func EncryptedTest() {
	appID := "wx4f4bc4dec97d474b"
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv := "r7BXXKkLb8qrSNn05n0qiA=="

	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appID, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
