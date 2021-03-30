package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	iv := []byte("1234567890123456")

	dataOrigen, err := ioutil.ReadFile("image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	dataEncriptada, err := encript(dataOrigen, key, iv, aes.BlockSize)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("image.enc", dataEncriptada, 0644)

	if err != nil {
		log.Fatal(err)
	}

	//dataEncriptada, err := ioutil.ReadFile("image.jpg.enc")

	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//dataEncriptada, err = base64.StdEncoding.DecodeString(string(dataEncriptada))

	//if err != nil {
	//	log.Fatal(err)
	//}

	dataDesencriptada, err := decrypt(key, iv, dataEncriptada)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("decrypt.jpg", dataDesencriptada, 0644)

	if err != nil {
		log.Fatal(err)
	}

}

func encript(plaintext, key, iv []byte, blockSize int) (ciphertext []byte, err error) {

	bPlaintext := PKCS5Padding(plaintext, blockSize, len(plaintext))

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	ciphertext = make([]byte, len(bPlaintext))

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext, bPlaintext)

	return
}

func decrypt(key, iv, encText []byte) (deciphertext []byte, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	decrypted := make([]byte, len(encText))

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decrypted, encText)

	deciphertext = PKCS5UnPadding(decrypted)

	return
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
