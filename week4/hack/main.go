package main

import (
	"fmt"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	// req, _ := httptool.NewRequest(
	// 	httptool.GETMETHOD, // 请求方法
	// 	"http://gtainmuxi.muxixyz.com/api/v1/organization/code", // 请求 URL
	// 	"",               // 请求体
	// 	httptool.DEFAULT, // 请求体类型
	// )
	// res, _ := req.SendRequest()
	// fmt.Println("Response:", res)
	// res.ShowHeader()
	pp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiNDkwNzEiLCJpYXQiOjE3MzE1MDQ2OTksIm5iZiI6MTczMTUwNDY5OX0.ZtG_ySlNOjHnzLQrdSBLWiQ_QTgsEtcm70bx4sViPLc"
	//encryptedData := "andBYVVxSUp6eXJXYzNlQW9sNFU4aExDNGdxbDdGZmRCTkR4bjZuQnZPTG95eFR0bFRKRnM5SXp2OGlJYkpXdA=="
	//secret_key := "MuxiStudio203304"
	//error_code := "for {go func(){time.Sleep(1*time.Hour)}()}"

	//得出加密过的error_code

	//encryptedData, _ := encrypt.AESEncryptOutInBase64([]byte(error_code), []byte(secret_key)) //用secret_key加密error_code
	//fmt.Println(string(encryptedData))
	encryptedData := "jwAaUqIJzyrWc3eAol4U8hLC4gql7FfdBNDxn6nBvOLoyxTtlTJFs9Izv8iIbJWt"
	request, _ := httptool.NewRequest(
		httptool.PUTMETHOD, // 请求方法
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", // 请求 URL
		//"http://gtainmuxi.muxixyz.com/api/v1/bank/gate", // 请求 URL
		encryptedData,
		httptool.DEFAULT,
	)

	// 打印请求信息（用于调试）
	fmt.Println("Request:", request)
	//pp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiTGVub3ZvIiwiaWF0IjoxNzMxNDIzNTkyLCJuYmYiOjE3MzE0MjM1OTJ9.4iwfhAqVrXDzqVYAXofil9IpTLfX5jfHMH6igvEcIvk"
	// 添加 passport 请求头
	request.AddHeader("passport", pp)

	// 发送请求
	resp, _ := request.SendRequest()

	// 打印响应
	fmt.Println("Response:", resp)

}
