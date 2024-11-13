package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
	// "github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	encryptedData := "andBYVVxSUp6eXJXYzNlQW9sNFU4aExDNGdxbDdGZmRCTkR4bjZuQnZPTG95eFR0bFRKRnM5SXp2OGlJYkpXdA=="

	request, err := httptool.NewRequest(
		httptool.PUTMETHOD, // 请求方法
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", // 请求 URL
		encryptedData,
		httptool.DEFAULT,
	)

	// 打印请求信息（用于调试）
	fmt.Println("Request:", request)

	// 添加 passport 请求头
	request.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiTGVub3ZvIiwiaWF0IjoxNzMxNDIzNTkyLCJuYmYiOjE3MzE0MjM1OTJ9.4iwfhAqVrXDzqVYAXofil9IpTLfX5jfHMH6igvEcIvk")

	// 发送请求
	resp, err := request.SendRequest()
	if err != nil {
		fmt.Println("failed to send request due to the error:", err)
		return
	}

	// 打印响应
	fmt.Println("Response:", resp)
}
