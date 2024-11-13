package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// define a handler
type countHandler struct {
	mu sync.Mutex
	n  int
}

// define method -> interface
func (h *countHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	h.mu.Lock()
	defer h.mu.Unlock()
	h.n -= 1
	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "this is countHandler no.%d \n", h.n)
	fmt.Fprintf(w, "count is %d \n", h.n)

}

func main() {
	http.Handle("/count", new(countHandler))
	http.ListenAndServe(":8080", nil)
}
