package main

import (
	"fmt"
	"io"
	"net/http"
)

//模拟登录
//遍历学号
//并发
//储存

func getMethod(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func postMethod(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {

}
