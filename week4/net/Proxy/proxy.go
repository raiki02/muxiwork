//error

// package main

// import (
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"os"
// )

// func main() {
// 	//proxyUrl , _ := url.Parse("socks5://127.0.0.1:1080")

// 	proxyUrl, _ := url.Parse("http://127.0.0.1:8087")
// 	t := &http.Transport{
// 		Proxy: http.ProxyURL(proxyUrl),
// 	}

// 	client := &http.Client{
// 		Transport: t,
// 	}

// 	r, _ := client.Get("https://google.com")
// 	defer r.Body.Close()

// 	_, _ = io.Copy(os.Stdout, r.Body)
// }

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// 使用代理
	proxyUrl, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	t := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	client := &http.Client{
		Transport: t,
	}

	// 发送 GET 请求
	r, err := client.Get("http://8.148.22.219")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer r.Body.Close()

	// 检查响应状态
	if r.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response status:", r.Status)
		return
	}

	// 复制响应内容到标准输出
	_, err = io.Copy(os.Stdout, r.Body)
	if err != nil {
		fmt.Println("Error copying response body:", err)
		return
	}
}
