package main

import (
	"errors"
	"fmt"
	"net/http"
)

func redirectLimitTimes() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 10 {
				return errors.New("redirect fking too many times")
			}
			return nil
		},
	}

	request, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/redirect/20",
		nil,
	)

	_, err := client.Do(request)
	if err != nil {
		panic(err)
	}
}

func redirectForbidden() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies.set?name=LIUYUAN",
		nil,
	)
	//fmt.Println("http.DefaultClient.Do(request)")
	//http.DefaultClient.Do(request)

	r, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(r.Request.URL)
}

func main() {
	//redirectLimitTimes()
	redirectForbidden()
}
