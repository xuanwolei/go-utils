package goutils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	//纳秒转换为秒
	HttpTimeOut = 60 * time.Second
)

type HttpRequest struct {
	Cookies []*http.Cookie
	Headers map[string]string
	Request *http.Request
	Client  *http.Client
	Resp    *http.Response
}

func HttpGet(uri string) ([]byte, error) {

	req, err := NewHttpRequest(uri, "GET", "")
	if err != nil {
		return nil, err
	}
	return req.Call()
}

func HttpPost(uri string, param interface{}) ([]byte, error) {
	req, err := NewHttpRequest(uri, "POST", param)
	if err != nil {
		return nil, err
	}
	return req.Call()
}

func NewHttpRequest(uri string, method string, param interface{}) (*HttpRequest, error) {
	var data string = ""
	var header = make(map[string]string)
	//设置默认header
	header["Content-Type"] = "application/x-www-form-urlencoded"
	//	urlVal, _ := url.Parse(uri)
	switch v := param.(type) {
	case map[string]string:
		for key, value := range v {
			data += "&" + key + "=" + value
		}
		break
	case string:
		data = v
		break
	}
	req, err := http.NewRequest(method, uri, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	return &HttpRequest{Request: req, Headers: header, Client: &http.Client{Timeout: time.Duration(HttpTimeOut)}}, nil
}

//调用
func (this *HttpRequest) Call() ([]byte, error) {
	var err error
	//设置header头
	for k, v := range this.Headers {
		this.Request.Header.Add(k, v)
	}
	//遍历cookie
	if len(this.Cookies) > 0 {
		for _, cookie := range this.Cookies {
			this.Request.AddCookie(cookie)
		}
	}

	this.Resp, err = this.Client.Do(this.Request)

	if err != nil {
		return nil, err
	}
	defer this.Resp.Body.Close()
	return ioutil.ReadAll(this.Resp.Body)
}

//将get请求的参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}
