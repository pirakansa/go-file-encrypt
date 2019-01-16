// Copyright (c) 2018 pirakansa
package main

import (
	"crypto/cipher"
	"flag"
	"io"
	"os"

	"github.com/pirakansa/go-file-encrypt/encryptiondata"
)

// xor
func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}

// main
func main() {
	var ifpath, ofpath, keypath string
	var enc, dec bool

	flag.BoolVar(&enc, "encode", false, "do encode")
	flag.BoolVar(&enc, "e", false, "(short) --encode")
	flag.BoolVar(&dec, "decode", false, "do decode")
	flag.BoolVar(&dec, "d", false, "(short) --decode")
	flag.StringVar(&ifpath, "if", "", "input file")
	flag.StringVar(&ifpath, "i", "", "(short) --if")
	flag.StringVar(&ofpath, "of", "", "output file")
	flag.StringVar(&ofpath, "o", "", "(short) --of")
	flag.StringVar(&keypath, "kf", "", "key file")
	flag.StringVar(&keypath, "k", "", "(short) --kf")
	flag.Parse()

	if !xor(enc, dec) {
		os.Exit(1)
	}

	infile, err := os.OpenFile(ifpath, os.O_RDONLY, 0755)
	if err != nil {
		os.Exit(1)
	}
	defer infile.Close()
	outfile, err := os.OpenFile(ofpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0755)
	if err != nil {
		os.Exit(1)
	}
	defer outfile.Close()

	var stream cipher.Stream
	key := encryptiondata.GetEncriptKey(keypath)
	if enc {
		stream = encryptiondata.GetCipherObj(key)
	} else {
		stream = encryptiondata.GetDecipherObj(key)
	}

	writer := cipher.StreamWriter{S: stream, W: outfile}
	io.Copy(writer, infile)

}
