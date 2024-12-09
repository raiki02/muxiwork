package main

import (
	"encoding/json"
	"fmt"
	"getname/db"
	"getname/login"
	"io"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Id          string `json:"id"`
	Pid         string `json:"Pid"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	SzLogonName string `json:"szLogonName"`
	SzHandPhone string `json:"szHandPhone"`
	SzTel       string `json:"szTel"`
	SzEmail     string `json:"szEmail"`
}

var (
	wg        sync.WaitGroup
	mu        sync.Mutex
	workers   = make(chan struct{}, 20)
	studentId = 2024000000
)

const EndId = 2024999999

// workers
func work() {
	defer wg.Done()
	for i := studentId; i < studentId+200; i++ {

		if i > EndId {
			fmt.Println("overend")
			break
		}

		time.Sleep(100 * time.Millisecond)
		mu.Lock()
		defer mu.Unlock()
		time := time.Now().UnixMilli()
		url := fmt.Sprintf("http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx?type=logonname&ReservaApply=ReservaApply&term=%d&_=%s", i, time)
		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{}
		//need ASP.NET_SessionId
		asp := login.GetCookie()[1]
		req.Header.Set("Cookie", asp)
		resp, _ := client.Do(req)
		bd, _ := io.ReadAll(resp.Body)
		var users []User
		json.Unmarshal(bd, &users)
		for _, user := range users {
			db.InsertUser(user.Id, user.Name, user.Pid, user.Label)
			//fmt.Println(user.Id, user.Name, user.Pid, user.Label)
		}
	}

}

func Begin() {
	total := 999999 / 200
	total += 1
	for i := 0; i < total; i++ {
		defer func() { <-workers }()
		workers <- struct{}{}
		wg.Add(1)
		go work()
	}

}

func check() {
	for {
		if len(workers) == 0 {
			break
		}
	}
	fmt.Println("done")
}

func main() {
	Begin()
	go check()
}
