package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printContent(resp *http.Response) {
	defer func() { _ = resp.Body.Close() }()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}
func requestByParams() {
	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)

	param := make(url.Values)
	param.Add("name", "Liuyuan")
	param.Add("age", "38")
	param.Add("girlfriends", "null")
	param.Add("wives", "null")
	param.Add("children", "2")

	request.URL.RawQuery = param.Encode()
	resp, _ := http.DefaultClient.Do(request)
	printContent(resp)
	fmt.Println("======================================================")
}

func requestByHead() {
	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	request.Header.Add("user-agent", "Liuyuan666")
	resp, _ := http.DefaultClient.Do(request)
	printContent(resp)
}

func main() {
	requestByParams()
	requestByHead()
}
