package main

//timeStamp, netSession
/*
roomId ,bgSeat,roomName:
101699179,N1173,南湖分馆一楼开敞座位区
101699187,N1001,南湖分馆一楼中庭开敞座位区
101699189,N2001,南湖分馆二楼开敞座位区
101699191,K2001,南湖分馆二楼卡座区
*/

//mk resv
/*
http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?
dialogid=&
dev_id=101699888& //seat_id
lab_id=&
kind_id=&
room_id=&
type=dev&
prop=&
test_id=&
term=&
Vnumber=&
classkind=&
test_name=&
start=2024-11-29+20%3A45& //settime
end=2024-11-29+22%3A00&  //settime
start_time=2045&		//settime
end_time=2200&		//settime
up_file=&
memo=&
act=set_resv&
_=1732611395713 //timestamp
*/

//mv resv
/*
http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?
act=del_resv&
id=	159555958&		//resv_id
_=1732611395714 //timestamp

*/
type Choice struct {
	roomId string
	bgSeat string
}

var choices = []Choice{
	{"101699179", "N1173"},
	{"101699187", "N1001"},
	{"101699189", "N2001"},
	{"101699191", "K2001"},
}

const (
	Typea = "南湖分馆一楼开敞座位区"
	Typeb = "南湖分馆一楼中庭开敞座位区"
	Typec = "南湖分馆二楼开敞座位区"
	Typed = "南湖分馆二楼卡座区"
)

func getSeat() {

}

func main() {

}
