package main

import (
	"TakeSeat/app"
	"TakeSeat/cookie"
)

/*
1模拟登录
2获取空座信息

	2.1橙色1/4圆为已被预约，覆盖在空座上
	2.2解析html获得未覆盖满的元素，再从里面提取座位号，展示给用户

3输入预约信息

	3.1输入的座位号转换成url中的dev_id

4取消
*/
func main() {
	cookie := cookie.GetCookie()
	app.Start(cookie)
}
