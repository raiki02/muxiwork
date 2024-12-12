package tools

import (
	"crypto/md5"
	"log"
	"strconv"
)

func Encrypt(pw string) (e_pw string) {
	h := md5.New()
	tmp, err := h.Write([]byte(pw))
	if err != nil {
		log.Fatal(err)
		return ""
	}
	e_pw = strconv.FormatInt(int64(tmp), 16)
	return e_pw
}
