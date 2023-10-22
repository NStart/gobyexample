package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPrt := flag.Int("numb", 42, "an int")
	forkPrt := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "s string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPrt)
	fmt.Println("fork:", *forkPrt)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
