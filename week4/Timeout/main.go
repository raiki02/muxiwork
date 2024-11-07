package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func timeout() {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
			ResponseHeaderTimeout: 5 * time.Second,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   2 * time.Second,
		},
	}
	r, _ := client.Get("http://httpbin.org/delay/10")
	defer r.Body.Close()
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func main() {
	timeout()
}
