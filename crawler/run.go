package crawler

import (
	"UrlScan/inflag"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client *http.Client

func load() {
	if inflag.U == "" {
		fmt.Println("input your target-url!")
		os.Exit(0)
	}
	u, ok := url.Parse(inflag.U)
	if inflag.U != "" && ok != nil {
		fmt.Println("url格式错误,请填写正确url")
		os.Exit(1)
	}
	inflag.U = u.String()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 30,
			KeepAlive: time.Second * 30,
		}).DialContext,
		MaxIdleConns:          inflag.TH / 2,
		MaxIdleConnsPerHost:   inflag.TH + 10,
		IdleConnTimeout:       time.Second * 90,
		TLSHandshakeTimeout:   time.Second * 90,
		ExpectContinueTimeout: time.Second * 10,
	}

	//加载proxy
	if inflag.PX != "" {
		tr.DisableKeepAlives = true
		proxyUrl, parseErr := url.Parse(inflag.PX)
		if parseErr != nil {
			fmt.Println("代理地址错误: \n" + parseErr.Error())
			os.Exit(1)
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	client = &http.Client{Timeout: time.Duration(inflag.TI) * time.Second,
		Transport: tr,
	}
}
func Run() {
	inflag.Init()
	load()
}
