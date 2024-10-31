package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s2 := "0123456789101112131415161718192021222324252627282930"

	ss1 := make([]string, 0)
	ss2 := make([]string, 0)

	for _, v1 := range s1 {
		ss1 = append(ss1, string(v1))
	}
	for _, v2 := range s2 {
		ss2 = append(ss2, string(v2))
	}

	var (
		i1 = 0
		i2 = 0
		oc = make([]string, 0)
		wg sync.WaitGroup
		lc sync.RWMutex
	)

	go func() {
		for i := 0; i < len(ss2)/2; i++ {
			wg.Add(1)
			defer wg.Done()
			lc.RLock()
			t2 := ss2[i2 : i2+2]
			oc = append(oc, t2...)
			time.Sleep(time.Second)
			i2 += 2
			if i2 >= len(ss2) {
				break
			}
			lc.RUnlock()
		}

	}()
	func() {
		for i := 0; i < len(ss1)/2; i++ {
			wg.Add(1)
			defer wg.Done()
			lc.RLock()
			t1 := ss1[i1 : i1+2]
			oc = append(oc, t1...)
			time.Sleep(time.Second)
			i1 += 2
			lc.RUnlock()
		}

	}()

	wg.Wait()
	fmt.Println(oc)
}
