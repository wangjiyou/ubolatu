package main

import (
	"flag"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"

	"ubolatu/config"
	"ubolatu/db"
	"ubolatu/test"
	"ubolatu/userinfo"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var configPtr = flag.String("c", "config/config.json", "Config file path, if ignored will be load from ./config/config.json")
	err := config.LoadConfigFile(*configPtr)
	if err != nil {
		logger.Log("config.LoadConfigFile err:", err)
		return
	}

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	//test start
	var svc test.StringService
	svc = test.UstringService{}
	svc = test.UloggingMiddleware{logger, svc}
	svc = test.UinstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	uppercaseHandler := httptransport.NewServer(
		test.UmakeUppercaseEndpoint(svc),
		test.UdecodeUppercaseRequest,
		test.UencodeResponse,
	)

	countHandler := httptransport.NewServer(
		test.UmakeCountEndpoint(svc),
		test.UdecodeCountRequest,
		test.UencodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	//test end

	db.InitMysql()
	db.InitTiDB()

	//----userinfo start
	var u_svc userinfo.StringService
	u_svc = userinfo.UstringService{}
	u_svc = userinfo.UloggingMiddleware{logger, u_svc}
	u_svc = userinfo.UinstrumentingMiddleware{requestCount, requestLatency, countResult, u_svc}

	u_setuserinfoHandler := httptransport.NewServer(
		userinfo.SetUserInfoEndpoint(u_svc),
		userinfo.UdecodeUserInfoRequest,
		userinfo.UencodeResponse,
	)

	u_loginHandler := httptransport.NewServer(
		userinfo.OnLoginEndpoint(u_svc),
		userinfo.OnLoginRequest,
		userinfo.UencodeResponse,
	)

	u_addFriendHandler := httptransport.NewServer(
		userinfo.AddFriendEndpoint(u_svc),
		userinfo.AddFriendRequest,
		userinfo.UencodeResponse,
	)

	u_delFriendHandler := httptransport.NewServer(
		userinfo.DelFriendEndpoint(u_svc),
		userinfo.DelFriendRequest,
		userinfo.UencodeResponse,
	)

	u_modiFriendHandler := httptransport.NewServer(
		userinfo.ModiFriendEndpoint(u_svc),
		userinfo.ModiFriendRequest,
		userinfo.UencodeResponse,
	)

	u_findFriendHandler := httptransport.NewServer(
		userinfo.FindFriendEndpoint(u_svc),
		userinfo.FindFriendRequest,
		userinfo.UencodeResponse,
	)

	http.Handle("/onLogin", u_loginHandler)
	http.Handle("/setUserInfo", u_setuserinfoHandler)
	http.Handle("/addFriend", u_addFriendHandler)
	http.Handle("/delFriend", u_delFriendHandler)
	http.Handle("/modiFriendType", u_modiFriendHandler)
	http.Handle("/findFriend", u_findFriendHandler)
	http.Handle("/findFocus", u_modiFriendHandler)
	http.Handle("/findFans", u_modiFriendHandler)

	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
