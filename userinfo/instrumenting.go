package userinfo

import (
	"fmt"
	"time"

	"ubolatu/pub"

	"github.com/go-kit/kit/metrics"
)

type UinstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           StringService
}

func (mw UinstrumentingMiddleware) ModiFriend(s pub.FriendShipRequest) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "modifriend", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	fmt.Println("instrument modiFriend")
	output, err = mw.Next.ModiFriend(s)
	return
}

func (mw UinstrumentingMiddleware) DelFriend(s pub.FriendShipRequest) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delfriend", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	fmt.Println("instrument DelFriend")
	output, err = mw.Next.DelFriend(s)
	return
}

func (mw UinstrumentingMiddleware) AddFriend(s pub.FriendShipRequest) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "addfriend", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	fmt.Println("instrument AddFriend")
	output, err = mw.Next.AddFriend(s)
	return
}

func (mw UinstrumentingMiddleware) SetUserInfo(s pub.FullUserInfo) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "set userinfos", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	fmt.Println("instrument SetUserInfo")
	output, err = mw.Next.SetUserInfo(s)
	return
}

func (mw UinstrumentingMiddleware) OnLogin(s pub.LoginRequest) (openId string, code int) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "on login", "error", fmt.Sprint(err != nil)}
		lvs := []string{"method", "OnLogin", "error", openId}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	openId, code = mw.Next.OnLogin(s)
	return
}

func (mw UinstrumentingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.CountResult.Observe(float64(n))
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
