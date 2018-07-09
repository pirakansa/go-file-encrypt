// Package encryptiondata ...
// Copyright (c) 2018 pirakansa
package encryptiondata

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// GetCipherObj get cipher object
func GetCipherObj(key *Encriptkey) cipher.Stream {
	block, err := aes.NewCipher([]byte(key.Encrypt.Key))
	if err != nil {
		os.Exit(1)
	}
	return cipher.NewCFBEncrypter(block, []byte(key.Encrypt.InitialVector))
}

// GetDecipherObj get decipher object
func GetDecipherObj(key *Encriptkey) cipher.Stream {
	block, err := aes.NewCipher([]byte(key.Encrypt.Key))
	if err != nil {
		os.Exit(1)
	}
	return cipher.NewCFBDecrypter(block, []byte(key.Encrypt.InitialVector))
}
