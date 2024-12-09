package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type User struct {
	StudentName string `json:"studentname" db:"studentname"`
	StudentID   string `json:"studentid" db:"studentid"`
	Grade       string `json:"grade" db:"grade"`
}

func login() string {
	req, _ := http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
		nil,
	)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Do(req)

	cookie := res.Header.Get("Set-Cookie")
	cookie, _, _ = strings.Cut(cookie, ";")
	return cookie
}

func get(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)
	cookie := login()
	fmt.Println(cookie)
	req.Header.Add("Cookie", cookie)
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

func print() {
	//for i := 2024000000; i <= 20249999999; i++ {
	time := time.Now().UnixMilli()
	fmt.Println(time)
	url := fmt.Sprintf("http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=%d&_=%d", 2024214274, time)
	fmt.Println(url)
	get(url)
	//}

}

func main() {
	//time := time.Now().UnixMilli()
	req, _ := http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
		nil,
	)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Do(req)

	cookie := res.Header.Get("Set-Cookie")
	cookie, _, _ = strings.Cut(cookie, ";")

	url := fmt.Sprintf("http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=%d&_=%d", 2024214274, 1732633798445)

	client = &http.Client{}
	req, _ = http.NewRequest(
		"GET",
		url,
		nil,
	)

	req.Header.Add("Cookie", cookie)
	res, _ = client.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}
