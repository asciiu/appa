package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CheckErr(str string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", str, err.Error())
		os.Exit(1)
	}
}

func GenerateNonce() ([]byte, string, error) {
	// Never use more than 2^32 random nonces with a given key because of
	// the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, "", err
	}

	return nonce, fmt.Sprintf("%x", nonce), nil

}

func ValidateKeyAndNonce(keyHexStr, nonceHexStr string) ([]byte, []byte, error) {
	key, err := hex.DecodeString(keyHexStr)
	if err != nil {
		return nil, nil, err
	}

	nonce, err := hex.DecodeString(nonceHexStr)
	if err != nil {
		return nil, nil, err
	}

	return key, nonce, nil
}

func Encrypt(key, nonce, body []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		return "", err
	}

	cipherText := aesgcm.Seal(nil, nonce, body, nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key []byte, nonce []byte, cipherHexStr string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHexStr)
	if err != nil {
		return "", err
	}

	plainText, err := aesgcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
