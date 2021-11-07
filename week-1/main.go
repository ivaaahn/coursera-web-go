package main

import "os"

func main() {
	err := unique(os.Stdin, os.Stdout)

	if err != nil {
		panic(err.Error())
	}
}
