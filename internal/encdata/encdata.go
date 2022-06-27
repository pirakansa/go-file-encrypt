// Package encdata ...
package encdata

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"os"
)

// Encriptkey
type Encriptkey struct {
	Key           [32]byte
	InitialVector [32]byte
}

// GetEncriptKey Encriptkey
func GetEncriptKey(pass string) Encriptkey {
	sh := sha256.Sum256([]byte(pass))
	md := md5.Sum([]byte(pass))
	sh_md := sha256.Sum256(md[:])
	return Encriptkey{Key: sh, InitialVector: sh_md}
}

// GetCipherObj get cipher object
func GetCipherObj(key Encriptkey) cipher.Stream {
	block, err := aes.NewCipher(key.Key[:])
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		os.Exit(2)
	}
	return cipher.NewCFBEncrypter(block, key.InitialVector[:block.BlockSize()])
}

// GetDecipherObj get decipher object
func GetDecipherObj(key Encriptkey) cipher.Stream {
	block, err := aes.NewCipher(key.Key[:])
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		os.Exit(2)
	}
	return cipher.NewCFBDecrypter(block, key.InitialVector[:block.BlockSize()])
}
