package tools

import (
	"TakeSeat/model"
	"encoding/json"
	"fmt"
)

func Convert(bd []byte, res *model.TResponse) {
	err := json.Unmarshal(bd, res)
	if err != nil {
		fmt.Println("json unmarshal err : ", err)
		return
	}
}
