package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		fmt.Println("env:", e)
		fmt.Println()

		pair := strings.SplitN(e, "=", 3)
		fmt.Println("pair:", pair)
		fmt.Println()

		fmt.Println("pair[0]:", pair[0])
		fmt.Println()
	}
}
