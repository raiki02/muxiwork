package main

//
import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

var wg sync.WaitGroup

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time
	msgs := make(chan MSG)
	go func() {
		for i := 0; ; i++ {
			//wg.Add(1)
			//defer wg.Done()
			msgs <- MSG{ms: message{Topic: "消费主题", Partition: 0, Offset: 0}, feedEvent: FeedEventDM{Type: "grade", UserID: i, Title: "成绩提醒", Content: "您的成绩是xxx"}}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for msg := range msgs {
		consumeMSG = append(consumeMSG, msg)
		if len(consumeMSG) >= ConsumeNum {
			go func() {
				m := consumeMSG[:ConsumeNum]
				fn(m)
				lastConsumeTime = time.Now()
				consumeMSG = consumeMSG[ConsumeNum:]
			}()
			//lastConsumeTime = time.Now()
			//	consumeMSG = consumeMSG[ConsumeNum:]
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			if len(consumeMSG) > 0 {
				go func() {
					m := consumeMSG[:ConsumeNum]
					fn(m)
					lastConsumeTime = time.Now()
					consumeMSG = consumeMSG[ConsumeNum:]
				}()
				//lastConsumeTime = time.Now()
				//consumeMSG = consumeMSG[ConsumeNum:]
			}
		}
	}
	//wg.Wait()
}
