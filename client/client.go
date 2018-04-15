package client

import (
	"bytes"
	//"crypto/hmac"
	//"crypto/sha1"
	//"encoding/hex"
	//"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	//"net/url"
	"time"

	//"ubolatu/config"
)

//var conf = config.GlobalConfig()

func HttpDo(_method string, _url string, _data []byte) (error, string) {
	var result string
	var err error
	err, result = _HttpDo(_method, _url, _data)
	if err != nil {
		err, result = _HttpDo(_method, _url, _data)
		if err != nil {
			err, result = _HttpDo(_method, _url, _data)
		}
	}
	return err, result
}

func _HttpDo(_method string, _url string, _data []byte) (error, string) {
	var body []byte
	var response *http.Response
	var request *http.Request
	var err error
	method := _method
	url := _url
	data := _data
	if 0 == len(string(data)) {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	}

	timeout := time.Duration(time.Duration(10) * time.Second)
	if err == nil {
		/*
			for k, v := range GenerateHeader(method, url) {
				request.Header.Set(k, v)
			}
		*/
		debug(httputil.DumpRequestOut(request, true))
		client := http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   2 * time.Second,
					Deadline:  time.Now().Add(3 * time.Second),
					KeepAlive: 2 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 2 * time.Second,
			},
			Timeout: timeout * time.Second,
		}
		response, err = client.Do(request)
		//response, err = (&http.Client{Timeout: timeout}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		debug(httputil.DumpResponse(response, true))
		body, err = ioutil.ReadAll(response.Body)
	}

	if err == nil {
		return nil, string(body)
	} else {
		return err, ""
	}
}

func debug(data []byte, err error) {
	if err == nil {
		//logger.Debug("%s\n\n", string(data))
	} else {
		//logger.Debug("err:", err)
	}
}

/*
func GetSign(method string, _url *url.URL, strtime string) string {
	mac := hmac.New(sha1.New, []byte(conf.SecretKey))
	mac.Write([]byte(method + "\n"))
	mac.Write([]byte(_url.Path + "\n"))
	mac.Write([]byte(strtime + "\n"))
	return hex.EncodeToString(mac.Sum(nil))
}

func GenerateHeader(method string, rawurl string) map[string]string {
	ret := make(map[string]string)
	url, _ := url.Parse(rawurl)
	timestr := time.Now().Format(time.UnixDate)
	signStr := GetSign(method, url, timestr)
	ret["Authorization"] = fmt.Sprintf("%s,%s", conf.AccessKey, signStr)
	ret["Content-Type"] = "application/json; charset=utf-8"
	ret["Date"] = timestr
	//ret["Connection"] = "keep-alive"
	ret["Connection"] = "close"
	return ret
}
*/
