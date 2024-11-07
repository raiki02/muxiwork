package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func body(r *http.Response) {
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
	fmt.Println("===========================================")
}

func status(r *http.Response) {
	fmt.Println("r.Status = ", r.Status)
	fmt.Println("r.StatusCode = ", r.StatusCode)
	fmt.Println("===========================================")

}

func header(r *http.Response) {
	fmt.Println(r.Header["content-type"])
	fmt.Println(r.Header.Get("content-type"))
	fmt.Println("===========================================")

}

func encoding(r *http.Response) {
	bufReader := bufio.NewReader(r.Body)
	bytes, _ := bufReader.Peek(1024)
	e, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	fmt.Println(e)

	body2 := transform.NewReader(bufReader, e.NewDecoder())
	content, _ := ioutil.ReadAll(body2)
	fmt.Println(string(content))
	fmt.Println("===========================================")

}

func main() {
	r, _ := http.Get("https://www.bilibili.com/")
	defer func() { _ = r.Body.Close() }()

	body(r)
	status(r)
	header(r)
	encoding(r)
}
