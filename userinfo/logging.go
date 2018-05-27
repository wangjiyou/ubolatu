package userinfo

import (
	"time"

	"ubolatu/pub"

	"github.com/go-kit/kit/log"
)

type UloggingMiddleware struct {
	Logger log.Logger
	Next   StringService
}

func (mw UloggingMiddleware) AddFriend(s pub.FriendShipRequest) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "AddFriend",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.AddFriend(s)
	return
}

func (mw UloggingMiddleware) SetUserInfo(s pub.FullUserInfo) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "SetUserInfo",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.SetUserInfo(s)
	return
}

func (mw UloggingMiddleware) OnLogin(s pub.LoginRequest) (openId string, code int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "OnLogin",
			"input", s,
			"output", code,
			"err", openId,
			"took", time.Since(begin),
		)
	}(time.Now())

	openId, code = mw.Next.OnLogin(s)
	return
}

func (mw UloggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
