package model

import "strconv"

func (i *Infos) InputInfo(year, startTime, endTime, seatId string) {
	i.Date = year
	i.BeginHour = startTime[:2]
	i.BeginMinute = startTime[2:]
	i.EndHour = endTime[:2]
	i.EndMinute = endTime[2:]
	i.SeatId = seatId
	i.Start = startTime
	i.End = endTime
}

/*
101699179,N1173,南湖分馆一楼开敞座位区
101699187,N1001,南湖分馆一楼中庭开敞座位区
101699189,N2001,南湖分馆二楼开敞座位区
101699191,K2001,南湖分馆二楼卡座区
*/

func ProcessnNUm(nNum int) int {
	if nNum >= 1000 && nNum <= 1999 {
		return 1
	} else if nNum >= 2000 && nNum <= 2999 {
		return 2
	} else if nNum >= 3000 && nNum <= 3999 {
		return 3
	}
	return 0

}

// 转换为dev_id
func processFinal() {

}

func (i *Infos) ProcessSeatId() {
	ab := i.SeatId[:1]
	sNum := i.SeatId[1:]
	nNum, _ := strconv.Atoi(sNum)
	choice := ProcessnNUm(nNum)
	switch ab {
	case "N":
		switch choice {
		case 1:

		case 2:

		case 3:

		default:
		}

	case "K":

	}
}
