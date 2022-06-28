package main

import (
	"crypto/cipher"
	"flag"
	"fmt"
	"io"
	"os"

	"fenc/internal/encdata"
	"fenc/internal/input"
)

// xor
func xor(a bool, b bool) bool {
	return a != b
}

// main
func main() {
	var ifpath, ofpath string
	var enc, dec bool

	flag.BoolVar(&enc, "encode", false, "do encode")
	flag.BoolVar(&enc, "e", false, "(short) --encode")
	flag.BoolVar(&dec, "decode", false, "do decode")
	flag.BoolVar(&dec, "d", false, "(short) --decode")
	flag.StringVar(&ifpath, "if", "", "input file")
	flag.StringVar(&ifpath, "i", "", "(short) --if")
	flag.StringVar(&ofpath, "of", "", "output file")
	flag.StringVar(&ofpath, "o", "", "(short) --of")
	flag.Parse()

	if !xor(enc, dec) {
		fmt.Println("need to option ( encode or decode )")
		os.Exit(1)
	}

	infile, err := os.OpenFile(ifpath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		os.Exit(1)
	}
	defer infile.Close()
	outfile, err := os.OpenFile(ofpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0755)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		os.Exit(1)
	}
	defer outfile.Close()

	pass, err := input.ReadPassword()
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
		os.Exit(1)
	}

	var stream cipher.Stream
	key := encdata.GetEncriptKey(pass)
	if enc {
		stream = encdata.GetCipherObj(key)
	} else {
		stream = encdata.GetDecipherObj(key)
	}

	writer := cipher.StreamWriter{S: stream, W: outfile}
	io.Copy(writer, infile)

}
