package userinfo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"ubolatu/pub"

	"github.com/go-kit/kit/endpoint"
)

const (
	TagLoginWeiXinServerUrl = "https://api.weixin.qq.com/sns/jscode2session"
	TagAuthCodeFlag         = "authorization_code"
	TagAppId                = "wx167df20a470f1429"
	TagAppSecret            = "df2a93e9e249ebc2a12f4841a6d503d7"
)

func AddFriendEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pub.FriendShipRequest)
		/*openid, code :=*/ svc.AddFriend(req)
		//fmt.Println("AddFriendEndpoint openid:", openid, " code:", code)
		return pub.UserResponse{"", 200}, nil
	}
}

func AddFriendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pub.FriendShipRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func SetUserInfoEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pub.FullUserInfo)
		//fmt.Println("set user info:", req)
		//fmt.Println("set user info req.rawData:", req.RawData)
		//fmt.Println("set user info req.UserInfo.NikeName:", req.UserInfo.NickName)

		//req := request.(pub.UserInfoRequest)

		fmt.Println("SetUserInfoEndpoint:", req)
		v, err := svc.SetUserInfo(req)
		if err != nil {
			return pub.UserResponse{v, 500}, nil
		}

		return pub.UserResponse{"OK", 200}, nil
	}
}

func UdecodeUserInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pub.FullUserInfo

	fmt.Println("------UdecodeUserInfoRequest------")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()

	err := json.Unmarshal([]byte(s), &request)
	if err != nil {
		fmt.Println("UdecodeUserInfoRequest err:", err, " body:", s)
		return nil, err
	}
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	fmt.Println("UdecodeUserInfoRequest err:", err)
	//	return nil, err
	//}
	//(appID string, sessionKey string, encryptedData string, iv string)
	//sessionKey := `4upDyvJumQRUKp6p9P\/\/Wg==`
	//EncryptedTest(TagAppId, string(sessionKey), request.EncryptedData, request.Iv)
	fmt.Println("UdecodeUserInfoRequest ok,request:", request)
	return request, nil
}

func OnLoginEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pub.LoginRequest)
		openid, code := svc.OnLogin(req)
		fmt.Println("OnLoginEndpoint openid:", openid, " code:", code)
		return pub.UserResponse{openid, code}, nil
	}
}

func OnLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pub.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func UencodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
