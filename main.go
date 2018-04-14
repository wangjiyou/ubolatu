package main

import (
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"

	"ubolatu/db"
	"ubolatu/test"
	"ubolatu/userinfo"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

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
	db.InitMysql()

	http.Handle("/u-setuserinfo", u_setuserinfoHandler)
	http.Handle("/onLogin", u_loginHandler)

	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
