package cookie

import (
	"fmt"
	"net/http"
	"strings"
)

func GetCookie() string {
	url := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"
	req, _ := http.NewRequest(
		"GET",
		url,
		nil)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Do(req)
	entire := res.Header.Get("Set-Cookie")
	target, _, ok := strings.Cut(entire, ";")
	if !ok {
		fmt.Println("Target cookie not found")
		return ""
	}
	return target

}
