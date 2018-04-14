package userinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type UserInfoRequest struct {
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

type UserResponse struct {
	Data       string `json:"data"`
	StatusCode int    `json:"statusCode"`
}

type LoginRequest struct {
	Code string `json:"code"`
}

const (
	//LoginWeiXinServerUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code"
	LoginWeiXinServerUrl = "https://api.weixin.qq.com/sns/jscode2session"
)

func SetUserInfoEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserInfoRequest)
		fmt.Println("SetUserInfoEndpoint:", req)
		v, err := svc.SetUserInfo(req.NickName)
		if err != nil {
			return UserResponse{v, 500}, nil
		}
		return UserResponse{v, 200}, nil
	}
}

func UdecodeUserInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func OnLoginEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		v, _ := svc.OnLogin(req)
		return UserResponse{v, 200}, nil
	}
}

func OnLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func UencodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

/*
type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
*/
