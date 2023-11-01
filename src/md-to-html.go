package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filepath := os.Args[1]

	contents, err := os.ReadFile(filepath)
	check(err)
	fmt.Println(string(contents))

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
