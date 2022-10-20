package main

import (
	"fmt"
	"os"

	"github.com/hard-shutdown/xmezum"
)

func encryptFile(path string, key []byte) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		return
	}
	statdata, err := file.Stat()
	if statdata.Size() == 0 || statdata.IsDir() {
		return
	}

	buf := make([]byte, statdata.Size())
	_, err = file.Read(buf)

	err = file.Truncate(0)
	_, err = file.Seek(0, 0)

	enc, err := xmezum.Encrypt(key, buf)
	if err != nil {
		return
	}

	n, err := file.Write(enc)
	if err != nil {
		return
	}
	fmt.Println(n)

	file.Close()
}

func decryptFile(path string, key []byte) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		return
	}
	statdata, err := file.Stat()
	if statdata.Size() == 0 || statdata.IsDir() {
		return
	}

	buf := make([]byte, statdata.Size())
	_, err = file.Read(buf)

	err = file.Truncate(0)
	_, err = file.Seek(0, 0)

	enc, err := xmezum.Decrypt(key, buf)
	if err != nil {
		return
	}

	n, err := file.Write(enc)
	if err != nil {
		return
	}
	fmt.Println(n)

	file.Close()
}
