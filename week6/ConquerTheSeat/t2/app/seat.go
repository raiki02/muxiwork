package app

import (
	"TakeSeat/model"
	"TakeSeat/tools"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Start(cookies []string) {
	var (
		date      string
		startTime string
		endTime   string
		seatId    string
		info      = model.Infos{}
		base      = "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?"
		client    = &http.Client{}
		resp      = model.TResponse{}
	)

	params := url.Values{}
	params.Add("dialogid", "")
	params.Add("up_file", "")
	params.Add("memo", "")
	params.Add("act", "set_resv")
	params.Add("lab_id", "")
	params.Add("kind_id", "")
	params.Add("room_id", "")
	params.Add("type", "dev")
	params.Add("prop", "")
	params.Add("test_id", "")
	params.Add("term", "")
	params.Add("Vnumber", "")
	params.Add("classkind", "")
	params.Add("test_name", "")

	for {
		fmt.Println("yyyy-mm-dd, hhmm, hhmm, seatId")
		fmt.Scan(&date, &startTime, &endTime, &seatId)
		info.InputInfo(date, startTime, endTime, seatId)
		fmt.Println("info ", info)

		params.Add("dev_id", info.SeatId)
		params.Add("start", info.Date+"+"+info.BeginHour+":"+info.BeginMinute)
		params.Add("end", info.Date+"+"+info.EndHour+":"+info.EndMinute)
		params.Add("start_time", info.Start)
		params.Add("end_time", info.End)
		params.Add("_", string(time.Now().UnixMilli()))

		targetUrl := base + params.Encode()

		req, err := http.NewRequest(
			"GET",
			targetUrl,
			nil,
		)
		if err != nil {
			fmt.Println("req made err : ", err)
			return
		}

		for _, cookie := range cookies {
			req.Header.Add("Cookie", cookie)
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client exec err : ", err)
			return
		}
		defer res.Body.Close()
		fmt.Println("conquering the seat...")

		bd, _ := io.ReadAll(res.Body)
		fmt.Println(string(bd))

		tools.Convert(bd, &resp)

		ret := resp.Ret
		if ret == 1 {
			fmt.Println("seat conquered")
			break
		} else {
			fmt.Println("seat not conquered")
			continue
		}
	}

}
