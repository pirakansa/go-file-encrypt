// Package encryptiondata ...
// Copyright (c) 2018 pirakansa
package encryptiondata

import (
	"encoding/json"
	"os"
)

// Encriptkey encrypt.json
type Encriptkey struct {
	Encrypt struct {
		Key           string `json:"Key"`
		InitialVector string `json:"InitialVector"`
	} `json:"Encrypt"`
}

// InitJSON initialze Encriptkey
func (key *Encriptkey) InitJSON() {
	json.Unmarshal([]byte(defaultEncriptData), &key)
}

const defaultEncriptData = `
{
    "Encrypt": {
        "Key": "=============256bit=============",
        "InitialVector": "=====128bit====="
    }
}
`

// GetEncriptKey Encriptkey
func GetEncriptKey(path string) *Encriptkey {
	key := new(Encriptkey)
	key.InitJSON()
	kfile, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return key
	}
	defer kfile.Close()
	data := make([]byte, 4096)
	n, err := kfile.Read(data)
	if err != nil {
		return key
	}
	_ = json.Unmarshal([]byte(data[:n]), &key)

	return key
}
