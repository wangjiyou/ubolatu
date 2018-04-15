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

func (mw UloggingMiddleware) SetUserInfo(s pub.UserInfoRequest) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.SetUserInfo(s)
	return
}

func (mw UloggingMiddleware) OnLogin(s pub.LoginRequest) (err error, code int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "OnLogin",
			"input", s,
			"output", code,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err, code = mw.Next.OnLogin(s)
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
