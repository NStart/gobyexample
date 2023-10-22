package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println(1)

	os.Exit(3)
}
