package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postForm() {
	u := make(url.Values)
	u.Add("Name", "Liuyuan")
	u.Add("Gf", "broke up")
	upload := u.Encode()

	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(upload),
	)

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func postJson() {

	u := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Yuanliu",
		Age:  81,
	}
	upload, _ := json.Marshal(u)

	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/json",
		bytes.NewReader(upload),
	)

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func main() {
	//postForm()
	postJson()
}
