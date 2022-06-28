package encdata

import (
	"bytes"
	"crypto/cipher"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFile_Encript(t *testing.T) {
	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()
	os.Stdout = nil

	testFile, _ := ioutil.TempFile("", "tmpfile")
	defer func() {
		// testFile.Close()
		os.Remove(testFile.Name())
	}()

	decfilepath := "../../test/internal/encdata/decriptfile"
	infile, _ := os.OpenFile(decfilepath, os.O_RDONLY, 0755)
	defer infile.Close()

	var stream cipher.Stream
	key := GetEncriptKey("dummy_passwd")
	stream = GetCipherObj(key)

	writer := cipher.StreamWriter{S: stream, W: testFile}
	io.Copy(writer, infile)

	data, _ := ioutil.ReadFile(testFile.Name())
	encfilepath := "../../test/internal/encdata/encriptfile"
	expdata, _ := ioutil.ReadFile(encfilepath)
	if bytes.Compare(data, expdata) != 0 {
		t.Errorf("not equal")
	}

}

func TestFile_Decript(t *testing.T) {
	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()
	os.Stdout = nil

	testFile, _ := ioutil.TempFile("", "tmpfile")
	defer func() {
		// testFile.Close()
		os.Remove(testFile.Name())
	}()

	encfilepath := "../../test/internal/encdata/encriptfile"
	infile, _ := os.OpenFile(encfilepath, os.O_RDONLY, 0755)
	defer infile.Close()

	var stream cipher.Stream
	key := GetEncriptKey("dummy_passwd")
	stream = GetDecipherObj(key)

	writer := cipher.StreamWriter{S: stream, W: testFile}
	io.Copy(writer, infile)

	data, _ := ioutil.ReadFile(testFile.Name())
	decfilepath := "../../test/internal/encdata/decriptfile"
	expdata, _ := ioutil.ReadFile(decfilepath)
	if bytes.Compare(data, expdata) != 0 {
		t.Errorf("not equal")
	}

}
