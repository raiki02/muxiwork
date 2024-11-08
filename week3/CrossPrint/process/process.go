package process

//
import (
	"sync"
)

var (
	Oc     = make([]string, 0)
	index1 = 0
	index2 = 0
	Wg     sync.WaitGroup
	m      sync.Mutex
)

func p1(ss1 []string) {
	Wg.Add(1)
	defer Wg.Done()
	m.Lock()
	//time.Sleep(time.Second)
	temp1 := ss1[index1 : index1+2]
	Oc = append(Oc, temp1...)
	index1 += 2
	m.Unlock()
}
func p2(ss2 []string) {
	Wg.Add(1)
	defer Wg.Done()
	m.Lock()
	//time.Sleep(time.Second)
	temp2 := ss2[index2 : index2+2]
	Oc = append(Oc, temp2...)
	index2 += 2
	m.Unlock()
}

func Crossprint(ss1, ss2 []string) {
	for i := 0; i < 13; i++ {
		p1(ss1)

		p2(ss2)
		Wg.Wait()
	}
}
