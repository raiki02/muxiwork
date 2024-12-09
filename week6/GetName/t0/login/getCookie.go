package login

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func GetCookie() (cookies []string) {
	//fmt.Println(" username & password")
	//var username, password string
	//fmt.Scan(&username, &password)
	firstUrl := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"
	req, _ := http.NewRequest("GET", firstUrl, nil)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Do(req)
	//c =  ASP.NET_SessionId=ctbwrs45x3e3wj55jshezsi1; path=/; HttpOnly
	temp1 := res.Header.Get("Set-Cookie")
	//ASP.NET_SessionId=ctbwrs45x3e3wj55jshezsi1
	c1, _, _ := strings.Cut(temp1, ";")

	req, _ = http.NewRequest(
		"GET",
		"https://account.ccnu.edu.cn/cas/login",
		nil,
	)
	client = &http.Client{}
	res, _ = client.Do(req)

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	//LT-143710-5GIqNLVcRpAQf9h2epn4toINXuUmHg-account.ccnu.edu.cn
	lt_value := doc.Find("input[name=lt]").AttrOr("value", "")
	//e1s1
	execution_value := doc.Find("input[name=execution]").AttrOr("value", "")
	//JSESSIONID=D7CBDC0E46658E1C38FCD5C2FFD907D4vTeH7M; Path=/cas; HttpOnly
	temp2 := res.Header.Get("Set-Cookie")
	//JSESSIONID=D7CBDC0E46658E1C38FCD5C2FFD907D4vTeH7M
	c2, _, _ := strings.Cut(temp2, ";")

	data := url.Values{}
	data.Set("username", Myusername)
	data.Set("password", Mypassword)
	data.Set("lt", lt_value)
	data.Set("execution", execution_value)
	data.Set("_eventId", "submit")
	data.Set("submit", "登录")

	req, _ = http.NewRequest(
		"POST",
		"https://account.ccnu.edu.cn/cas/login;"+c2,
		bytes.NewBufferString(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", c2)

	res, _ = client.Do(req)
	defer res.Body.Close()
	//CASTGC=TGT-16158-Dq2gFfAi9AVB0CBOYhlcFdh9vLiDnLSeWdbqGwQ05NMP0OxyXx-account.ccnu.edu.cn; Path=/cas/
	temp3 := res.Cookies()[1]
	//CASTGC=TGT-16158-Dq2gFfAi9AVB0CBOYhlcFdh9vLiDnLSeWdbqGwQ05NMP0OxyXx-account.ccnu.edu.cn
	c3, _, _ := strings.Cut(temp3.String(), ";")

	cookies = []string{c1, c2, c3}

	return cookies
}
