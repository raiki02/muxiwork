package model

type TResponse struct {
	Ret  int         `json:"ret"`
	Act  string      `json:"act"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ext  interface{} `json:"ext"`
}
