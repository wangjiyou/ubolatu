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
	OnLogin(string) (string, error)
}

type UstringService struct{}

func (UstringService) SetUserInfo(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	fmt.Println("SetUserInfo NikeName:", s)
	return strings.ToUpper(s), nil
}

func (UstringService) OnLogin(s string) (string, error) {
	fmt.Println("onlogin:", s)
	return s, nil
}

func (UstringService) Count(s string) int {
	fmt.Println("userinfo:", s)
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
