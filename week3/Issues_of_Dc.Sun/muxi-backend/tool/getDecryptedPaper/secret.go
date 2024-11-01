package getDecryptedPaper

import (
	"encoding/base64"
)

func xorEncryptDecrypt(input string, key string) string {
	keyLen := len(key)
	output := make([]byte, len(input))

	for i := range input {
		output[i] = input[i] ^ key[i%keyLen]
	}

	return string(output)
}
func GetDecryptedPaper(encodedPaper string, key string) string {
	data, _ := base64.StdEncoding.DecodeString(encodedPaper)
	return xorEncryptDecrypt(string(data), key)
}
