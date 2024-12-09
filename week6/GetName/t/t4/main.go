package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	//map[0:LT-417780-4XTv3WiT332prkFiXLjBYtf9n7qw2r-account.ccnu.edu.cn 1:e1s1]
	kv := make(map[int]string, 4)

	req, _ := http.NewRequest(
		"GET",
		"http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx",
		nil)
	client := http.Client{}
	res, _ := client.Do(req)
	bd, _ := io.ReadAll(res.Body)
	html := string(bd)

	sectionRegex := regexp.MustCompile(`<section class="row btn-row">([\s\S]*?)</section>`)
	sectionMatch := sectionRegex.FindStringSubmatch(html)

	inputRegex := regexp.MustCompile(`<input[^>]*value="([^"]*)"[^>]*>`)
	inputMatches := inputRegex.FindAllStringSubmatch(sectionMatch[1], 4)

	for i, inputMatch := range inputMatches {
		kv[i] = inputMatch[1]
	}
	//JSESSIONID=780B19501E9FB1C05881F121089D13FBgBrRio; Path=/cas; HttpOnly
	cookies := res.Header.Get("Set-Cookie")
	//JSESSIONID=780B19501E9FB1C05881F121089D13FBgBrRio
	cookie_l, _, _ := strings.Cut(cookies, ";")
	cookie := res.Cookies()[0]

	turl := "\nhttps://account.ccnu.edu.cn/cas/login;jsessionid=" + cookie.Value

	formData := url.Values{
		"username":  {"2024214274"},
		"password":  {"wasd%2Biop4444"},
		"lt":        {kv[0]},
		"execution": {kv[1]},
		"_eventId":  {kv[2]},
		"submit":    {"%E7%99%BB%E5%BD%95"},
	}

	body := bytes.NewBufferString(formData.Encode())
	req, _ = http.NewRequest(
		"POST",
		turl,
		body,
	)
	req.Header.Add("Cookie", cookie_l)
	res, _ = client.Do(req)
	cooki1 := res.Header.Get("Set-Cookie")
	cookie2 := res.Header.Get("Cookie")
	fmt.Println(cooki1)
	fmt.Println(cookie2)
}
