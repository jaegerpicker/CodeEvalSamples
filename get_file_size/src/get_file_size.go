package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat(os.Args[1])
	if err == nil {
		fmt.Println(fi.Size())
	}
}
