package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/http"
)

func get() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	fmt.Println("=============================================")
}

func post() {
	resp, _ := http.Post("http://httpbin.org/post", "", nil)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	fmt.Println("=============================================")
}

func put() {
	request, _ := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	r, _ := http.DefaultClient.Do(request)
	defer func() { _ = r.Body.Close() }()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	fmt.Println("=============================================")

}

func delete() {
	request, _ := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	r, _ := http.DefaultClient.Do(request)
	defer func() { _ = r.Body.Close() }()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	fmt.Println("=============================================")
}

func main() {
	get()
	post()
	put()
	delete()
}
