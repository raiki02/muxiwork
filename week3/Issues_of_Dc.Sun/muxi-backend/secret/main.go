package main

import (
	"io/ioutil"
	"net/http"

	afero "github.com/spf13/afero"
)

func main() {
	url := "http://121.43.151.190:8000/"
	resp1, _ := http.Get(url + "paper")

	resp2, _ := http.Get(url + "secret")

	Paper, _ := ioutil.ReadAll(resp1.Body)

	Key, _ := ioutil.ReadAll(resp2.Body)

	Encrypted := afero.xorEncryptDecrypt(string(Paper), string(Key))
	Decrypted := afero.GetDecryptedPaper(Encrypted, string(Key))
	afero.SavePaper("../paper/Dc.Sun's_papers.txt", Decrypted)
	defer resp1.Body.Close()
	defer resp2.Body.Close()
}
