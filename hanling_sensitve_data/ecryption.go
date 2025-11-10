package hanling_sensitve_data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var Key = "thisis32bitlongpassphraseimusing"

func Encrypt(plaintext, key string) string {
	//generate cipher text block using algorithm and key
	//generate random vector key using cipher block
	//encrypt the plaintext
	//encode it using base64 encoding

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	vector := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, vector)
	if err != nil {
		fmt.Println(err)
	}
	stream := cipher.NewCFBEncrypter(block, vector)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext)

}

func Decrypt(ciphertext, key string) string {
	encrypted, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Println(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	vector := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, vector)
	stream.XORKeyStream(encrypted, encrypted)
	return string(encrypted)

}
