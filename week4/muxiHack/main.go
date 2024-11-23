package main

import (

	//	"encoding/json"

	"fmt"
	"net/http"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
	//	"strings"
	//	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	//	"github.com/spf13/viper"
)

func t(url string) {
	request, _ := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://gtainmuxi.muxixyz.com/api/v1/"+url,
		"",
		httptool.DEFAULT,
	)
	client := http.Client{}
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://gtainmuxi.muxixyz.com/api/v1/organization/code",
		nil,
	)

	req.Header.Add("code", "114")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	pp := res.Header.Get("passport")

	request.AddHeader("passport", pp)
	response, _ := request.SendRequest()
	fmt.Println(url, "\n")
	mp, _ := response.GetHeader("map-fragments")
	fmt.Println("mp = ", mp)
}

func main() {
	var base = "http://gtainmuxi.muxixyz.com/api/v1/"
	//var ctt []byte
	client := http.Client{}
	req, _ := http.NewRequest(
		http.MethodGet,
		base+"organization/code",
		nil,
	)

	req.Header.Add("code", "114")
	res, _ := client.Do(req)

	defer res.Body.Close()
	pp := res.Header.Get("passport")
	fmt.Println("pp = ", pp)
	//pp = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMjAwIiwiaWF0IjoxNzMxMzAzNTExLCJuYmYiOjE3MzEzMDM1MTF9.5x7bQLzIrmrnv_G4Bv0q-v51dOXyiKTD5G4siZSahtM

	// req, err = http.NewRequest(http.MethodGet, url11, nil)
	// req.Header.Add("passport", pp)
	// if err != nil {
	// 	panic(err)
	// }
	// /* res body
	// {
	// "code":0,
	// "message":"OK",
	// "data":{
	// 	"text":"恭喜你拿到了 passport，现在你可以着手准备骇入银行。\r\n银行的第一道门是代码安全门，我们计划将错误代码写入此门来破解它。\r\n但是这个门具有识别明文代码的功能，所以我们还需要一个密钥加密我们的错误代码，再写入至此门。\r\n不需要担心，两者我们都为你提供了，你只需要解析 response 中的密文（在 ExtraInfo 中）来得到它们。\r\n你现在的任务：\r\n解析密文，获取 error_code 和 secret_key\r\n使用 secret_key 加密 error_code\r\n然后将加密过的 error_code 放入 请求body 并以 \"正确的请求方法\" 发送至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate , 同时注意response的信息。",
	// 	"extra_info":"c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	// 	}
	// }*/
	// res, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// _, _ = ioutil.ReadAll(res.Body)
	// //fmt.Printf("ctt: %s \n", ctt)

	// str := `{
	// "code":0,
	// "message":"OK",
	// "data":{
	// 	"text":"恭喜你拿到了 passport，现在你可以着手准备骇入银行。\r\n银行的第一道门是代码安全门，我们计划将错误代码写入此门来破解它。\r\n但是这个门具有识别明文代码的功能，所以我们还需要一个密钥加密我们的错误代码，再写入至此门。\r\n不需要担心，两者我们都为你提供了，你只需要解析 response 中的密文（在 ExtraInfo 中）来得到它们。\r\n你现在的任务：\r\n解析密文，获取 error_code 和 secret_key\r\n使用 secret_key 加密 error_code\r\n然后将加密过的 error_code 放入 请求body 并以 \"正确的请求方法\" 发送至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate , 同时注意response的信息。",
	// 	"extra_info":"c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	// 	}
	// }`

	// viper.SetConfigType("json")

	// if err := viper.ReadConfig(strings.NewReader(str)); err != nil {
	// 	fmt.Println(err)
	// }

	// // c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==
	// _ = viper.Get("data.extra_info")

	//fmt.Println(encrypt.Base64Decode(extra_info.(string)))

	//fmt.Println("===")

	// secret_key := "MuxiStudio203304"
	// error_code := "for {go func(){time.Sleep(1*time.Hour)}()}"
	// /*[106 119 65 97 85 113 73 74 122 121 114 87 99 51 101 65 111 108 52 85 56 104 76 67 52 103 113 108 55 70 102 100 66 78 68 120 110 54 110 66 118 79 76 111 121 120 84 116 108 84 74 70 115 57 73 122 118 56 105 73 98 74 87 116]`*/
	// encrypted_code, _ := encrypt.AESEncryptOutInBase64([]byte(error_code), []byte(secret_key))
	// // fmt.Println(string(encrypted_code))
	// data := map[string]interface{}{
	// 	"content":    string(encrypted_code),
	// 	"extra_info": "",
	// }
	// dataStr, _ := json.Marshal(data)
	// req, _ = http.NewRequest(
	// 	http.MethodPut,
	// 	base+"bank/gate",
	// 	strings.NewReader(string(dataStr)),
	// )
	// req.Header.Add("passport", pp)
	// req.Header.Add("code", "200")
	// req.Header.Add("Content-Type", "application/json")
	// res, _ = http.DefaultClient.Do(req)
	// ctt, _ := ioutil.ReadAll(res.Body)
	// fmt.Printf("ctt: %s \n", ctt)
	// mp := res.Header.Get("map-fragments")
	// fmt.Println("mp = ", mp)

	//4.1
	//-> jpg
	// req, _ = http.NewRequest(
	// 	http.MethodGet,
	// 	base+"organization/iris_sample",
	// 	nil,
	// )
	// req.Header.Add("passport", pp)
	// req.Header.Add("code", "200")
	// res, _ = client.Do(req)

	// //jpg ->

	// buf := new(bytes.Buffer)
	// writer := multipart.NewWriter(buf)
	// contenttype := writer.FormDataContentType()
	// formFile, _ := writer.CreateFormFile("file", "file")
	// _, _ = io.Copy(formFile, res.Body)
	// _ = writer.Close()

	// req, _ = http.NewRequest(
	// 	http.MethodPost,
	// 	base+"bank/iris_recognition_gate",
	// 	buf,
	// )

	// req.Header.Set("content-type", contenttype)
	// req.Header.Add("passport", pp)
	// req.Header.Add("code", "200")
	// res, _ = client.Do(req)

	// ctt, _ = ioutil.ReadAll(res.Body)
	// fmt.Println(string(ctt))
	4.2
	request, _ := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://gtainmuxi.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT,
	)
	request.AddHeader("passport", pp)
	response, _ := request.SendRequest()
	_ = response.Save("iris_sample.jpg")

	request, _ = httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://gtainmuxi.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"iris_sample.jpg",
		httptool.FILE,
	)
	request.AddHeader("passport", pp)
	response, _ = request.SendRequest()
	response.ShowBody()
	mp, _ := response.GetHeader("map-fragments")
	fmt.Println("mp = ", mp)

	//5
	//t("organization/code")
	//t("organization/secret_key")
	//t("bank/gate")
	//t("bank/iris_recognition_gate")
	//t("organization/iris_sample")

	//rt := "muxi/backend/computer/examination"
	request, _ := httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"permute.go",
		httptool.FILE,
	)
	request.AddHeader("passport", pp)
	response, _ := request.SendRequest()
	response.ShowBody()
}
