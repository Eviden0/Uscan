package test

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"testing"
)

// 验证是否有 Set-Cookie 响应头 仅仅提取返回包里面的cookie
// SimpleCookieJar 是一个简单的 CookieJar 实现
type SimpleCookieJar struct {
	mu      sync.Mutex
	storage map[string][]*http.Cookie // 用于存储每个 URL 的 Cookie
}

// NewSimpleCookieJar 创建一个新的 SimpleCookieJar 实例
func NewSimpleCookieJar() *SimpleCookieJar {
	return &SimpleCookieJar{
		storage: make(map[string][]*http.Cookie),
	}
}

// SetCookies 实现 CookieJar 的 SetCookies 方法
func (jar *SimpleCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.mu.Lock()
	defer jar.mu.Unlock()

	host := u.Host
	// 合并新 Cookie 和已有 Cookie，去重
	existingCookies := jar.storage[host]
	for _, newCookie := range cookies {
		found := false
		for i, existingCookie := range existingCookies {
			if existingCookie.Name == newCookie.Name {
				existingCookies[i] = newCookie // 替换旧的 Cookie
				found = true
				break
			}
		}
		if !found {
			existingCookies = append(existingCookies, newCookie)
		}
	}
	jar.storage[host] = existingCookies
}

// Cookies 实现 CookieJar 的 Cookies 方法
func (jar *SimpleCookieJar) Cookies(u *url.URL) []*http.Cookie {
	jar.mu.Lock()
	defer jar.mu.Unlock()

	return jar.storage[u.Host]
}
func TestGetHtml(t *testing.T) {
	// 创建自定义的 CookieJar
	jar := NewSimpleCookieJar()

	// 创建自定义 HTTP 客户端
	client := &http.Client{Jar: jar}

	// 发起一个请求
	resp, err := client.Get("https://portal2020.xtu.edu.cn/cas/login")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	resp.Body.Close()

	// 检查存储的 Cookie
	url, _ := url.Parse("https://portal2020.xtu.edu.cn/cas/login")
	cookies := jar.Cookies(url)
	fmt.Println("Cookies stored in the jar:")
	for _, cookie := range cookies {
		fmt.Printf("%s = %s\n", cookie.Name, cookie.Value)
	}
}
func TestGetHtml2(t *testing.T) {
	// 创建一个标准库的 CookieJar
	jar, _ := cookiejar.New(nil)

	// 创建一个 HTTP 客户端并关联 CookieJar
	client := &http.Client{Jar: jar}

	// 发起一个请求，设置 Cookie
	resp, err := client.Get("https://www.op.gg/")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 检索并打印指定 URL 的 Cookies
	targetURL, _ := url.Parse("https://www.op.gg/")
	cookies := jar.Cookies(targetURL)
	fmt.Println("Cookies for URL:", targetURL)
	for _, cookie := range cookies {
		fmt.Printf("%s = %s\n", cookie.Name, cookie.Value)
	}
}
