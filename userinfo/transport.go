package userinfo

import (
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
	TagAppId                = "123"
	TagAppSecret            = "qwe"
)

func SetUserInfoEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pub.UserInfoRequest)
		fmt.Println("SetUserInfoEndpoint:", req)
		v, err := svc.SetUserInfo(req)
		if err != nil {
			return pub.UserResponse{v, 500}, nil
		}
		return pub.UserResponse{v, 200}, nil
	}
}

func UdecodeUserInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pub.UserInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func OnLoginEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pub.LoginRequest)
		_, code := svc.OnLogin(req)
		return pub.UserResponse{"", code}, nil
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
