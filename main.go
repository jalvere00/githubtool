package main

import (
	"flag"
	"fmt"
	"os"
)

var nameFlag = flag.String("name", "John", "Test Flag for names.")

func main() {
	flag.Parse()
	argsWithProg := os.Args

	fmt.Println("Args: ", argsWithProg)
	fmt.Println("Flag: ", *nameFlag)
	fmt.Println(getRepoRelease("lencx", "ChatGPT"))
}
