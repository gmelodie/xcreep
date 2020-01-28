package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: ./xcreap <INPUT FILE>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	content, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	err = clipboard.WriteAll(string(content))

	if err != nil {
		panic(err)
	}
}
