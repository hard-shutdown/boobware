package main

import (
	"fmt"

	"github.com/karrick/godirwalk"
)

func main() {
	key := []byte{160, 185, 192, 168, 229, 105, 196, 195, 171, 99, 112, 156, 149, 205, 170, 229}

	dirname := "/home/gitpod/.cargo"
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			decryptFile(osPathname, key)
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err == nil {
		return
	} else {
		fmt.Println(err)
	}
}
