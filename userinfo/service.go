package userinfo

import (
	"errors"
	"fmt"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	SetUserInfo(string) (string, error)
	Count(string) int
	OnLogin(LoginRequest) (string, error)
}

type UstringService struct{}

func (UstringService) SetUserInfo(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	fmt.Println("SetUserInfo NikeName:", s)
	return strings.ToUpper(s), nil
}

func (UstringService) OnLogin(request LoginRequest) (string, error) {
	fmt.Println("onlogin:", request.Code)
	//v := url.Values{}
	return request.Code, nil
}

func (UstringService) Count(s string) int {
	fmt.Println("userinfo:", s)
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
