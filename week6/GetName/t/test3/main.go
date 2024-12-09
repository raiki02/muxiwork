package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	req, _ := http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
		nil,
	)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	client.Do(req)

	req, _ = http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=2024214274&_=1732634472991",
		nil,
	)
	req.Header.Add("Cookie", "ASP.NET_SessionId=aebiynrnwx00p1eqcd5we155")
	res, _ := client.Do(req)

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}
