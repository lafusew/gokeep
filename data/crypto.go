package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
)

func Encrypt(str string, key string) (string, error) {
	text := []byte(str)
	
	k, err:= keyTo32byteArr(key)
	
	if err != nil {
		log.Fatalln(err.Error())
	}

	c, err := aes.NewCipher(k)

	if err != nil {
		log.Fatalln(err.Error())
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatalln(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Print(err);
	}

	e := gcm.Seal(nonce, nonce, text, nil)

	return string(e), err
}

func Decrypt(str []byte, key string) (string, error){
	k, err := keyTo32byteArr(key)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c, err := aes.NewCipher(k)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatalln(err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(str) < nonceSize {
		log.Fatalln(err.Error())
	}

	nonce, encrypted := str[:nonceSize], str[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(plaintext), err
}

func keyTo32byteArr(key string) ([]byte, error) {
	var err error = nil
	if len(key) > 32 {
		err = errors.New("key can't be more than 32 characters")
	}

	k := []byte(key)

	for len(k) < 32 {
		b := []byte{0}
		k = append(k, b...)
	}

	return k, err
}
