package userinfo

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/xlstudio/wxbizdatacrypt"

	"ubolatu/client"
	"ubolatu/pub"
)

// StringService provides operations on strings.
type StringService interface {
	SetUserInfo(string) (string, error)
	Count(string) int
	OnLogin(pub.LoginRequest) (string, error)
}

type UstringService struct{}

func (UstringService) SetUserInfo(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	fmt.Println("SetUserInfo NikeName:", s)
	return strings.ToUpper(s), nil
}

func (UstringService) OnLogin(request pub.LoginRequest) (string, error) {
	fmt.Println("onlogin:", request.Code)
	//v := url.Values{}
	//get openid and serectkey
	/*
		if IsExist(openid)
		    UpdateDBSecretKey()
		else
			GetUserInfo()
			AddDBUserInfo()
	*/
	return request.Code, nil
}

func GetSession(code string) (string, string) {
	//LoginWeiXinServerUrl = "https://api.weixin.qq.com/sns/jscode2session?
	//appid=APPID&secret=APPSECRET&js_code=JSCODE&grant_type=authorization_code"
	value := url.Values{}
	value.Set("appid", TagAppId)
	value.Set("secret", TagAppSecret)
	value.Set("js_code", code)
	value.Set("grant_type", TagAuthCodeFlag)
	loginUrl := fmt.Sprintf("%s?%s", TagLoginWeiXinServerUrl, value.Encode())
	client.HttpDo("GET", loginUrl, []byte(""))
	var openId string
	var sessionKey string
	return openId, sessionKey
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
