package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	afero "github.com/spf13/afero"
)

func main() {
	url := "http://121.43.151.190:8000/"
	resp1, err := http.Get(url + "paper")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
		return
	}
	resp2, err := http.Get(url + "secret")
	if err != nil {
		log.Fatal("failed to send request : %v", err)
		return
	}
	Paper, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("read from paper err : ", err)
		return
	}
	Key, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("read from paper err : ", err)
		return
	}
	Encrypted := afero.xorEncryptDecrypt(string(Paper), string(Key))
	Decrypted := afero.GetDecryptedPaper(Encrypted, string(Key))
	afero.SavePaper("../paper/Dc.Sun's_papers.txt", Decrypted)
	defer resp1.Body.Close()
	defer resp2.Body.Close()
}
