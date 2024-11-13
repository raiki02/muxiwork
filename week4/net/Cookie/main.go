package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func rrCookie() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	firstRequest, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies/set?name=liuyuan&password=123456",
		nil,
	)

	firstResponse, _ := client.Do(firstRequest)
	defer firstResponse.Body.Close()

	secondRequest, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies",
		nil,
	)
	for _, cookie := range firstResponse.Cookies() {
		secondRequest.AddCookie(cookie)
	}
	secondresponse, _ := client.Do(secondRequest)
	defer secondresponse.Body.Close()

	content, _ := ioutil.ReadAll(secondresponse.Body)

	fmt.Printf("%s \n", content)
}

func jarCookie() {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}

	r, _ := client.Get("http://httpbin.org/cookies/set?name=hello&password=world")
	defer r.Body.Close()

	_, _ = io.Copy(os.Stdout, r.Body)
}

func main() {
	//rrCookie()
	jarCookie()
}
