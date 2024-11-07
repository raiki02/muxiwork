package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadfile(url string, filename string) {
	r, _ := http.Get(url)
	defer func() { _ = r.Body.Close() }()

	space, _ := os.Create(filename)
	defer func() { _ = space.Close() }()
	n, _ := io.Copy(space, r.Body)
	fmt.Println(n)
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf(" \r进度： %.2f%%", float32(r.Current*10000/r.Total)/100)

	return
}

func downloadProcess(url, filename string) {
	r, _ := http.Get(url)
	defer func() { _ = r.Body.Close() }()
	space, _ := os.Create(filename)
	defer func() { _ = space.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	_, _ = io.Copy(space, reader)

}

func main() {
	var url, filename string
	//downloadProcess(url, filename)
	fmt.Print("下载网址, 存储文件名:")
	fmt.Scanln(&url, &filename)

	downloadProcess(url, filename)
}
