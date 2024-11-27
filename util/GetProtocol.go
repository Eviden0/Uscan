package util

import (
	"net/http"
	"strings"
)

// GetProtocol 判断http协议
func GetProtocol(domain string) string {
	if strings.HasPrefix(domain, "http") {
		return domain
	}

	response, err := http.Get("https://" + domain)
	if err == nil {
		return "https://" + domain
	}
	response, err = http.Get("http://" + domain)
	if err == nil {
		return "http://" + domain
	}
	defer response.Body.Close()
	if response.TLS == nil {
		return "http://" + domain
	}
	return ""
}
