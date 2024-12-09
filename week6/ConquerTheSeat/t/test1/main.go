package main

import (
	"fmt"
	"io"
	"net/http"
)

//模拟登录
//时间，位置，持续
//直到成功
//取消

var (
	RoomId    = 101699189
	Date      = "2024-11-26"
	StartTime = "17%3A30"
	EndTime   = "22%3A00"
	SessionID = 1732610169352
)

// 101699179,N1173|101699187,N1001|101699189,N2001|101699191,K2001
type A struct {
	DevId string
	DevNm string
}

func (aaa *A) ch(ta, tb string) {
	aaa.DevId = ta[:1]
	aaa.DevNm = tb[1:]
}
func main() {
	var aa = A{}

	var ta string
	var tb string

	fmt.Scan(&ta, &tb)
	fmt.Println("ta ", ta)
	fmt.Println("tb ", tb)
	aa.ch(ta, tb)
	fmt.Println("aa.DevNm", aa.DevNm)
	fmt.Println("aa.DevId", aa.DevId)
	fmt.Println("aa", aa)
}

func main1() {

	rq, _ := http.NewRequest("GET", "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=101699888&lab_id=&kind_id=&room_id=&type=dev&prop=&test_id=&term=&Vnumber=&classkind=&test_name=&start=2024-11-26+nu%3All&end=2024-11-26+22%3A00&end_time=2200&up_file=&memo=&act=set_resv&_=1732611395709", nil)
	rq.Header.Add("Cookie", "ASP.NET_SessionId=weydpce2el3jfj55z51dkn45")
	rs, _ := http.DefaultClient.Do(rq)
	bd, _ := io.ReadAll(rs.Body)
	fmt.Println(string(bd))

	return
	TargetUrl := fmt.Sprintf("http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/device.aspx?byType=devcls&classkind=8&display=fp&md=d&room_id=%d&purpose=&selectOpenAty=&cld_name=default&date=%s&fr_start=%s&fr_end=%s&act=get_rsv_sta&_=%d", RoomId, Date, StartTime, EndTime, SessionID)
	fmt.Println(TargetUrl)
	req, err := http.NewRequest(
		"POST",
		TargetUrl,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}

	req.Header.Add("Cookie", "ASP.NET_SessionId=weydpce2el3jfj55z51dkn45")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
