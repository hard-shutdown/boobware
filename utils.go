package main

import (
	"os"

	"github.com/hard-shutdown/xmezum"
)

func encryptFile(path string, key []byte) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		return
	}
	statdata, err := file.Stat()
	buf := make([]byte, statdata.Size())
	_, err = file.Read(buf)
	xmezum.Encrypt(buf, key)
}
