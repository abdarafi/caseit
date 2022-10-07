package main

import (
	"fmt"
	"os"
	"strcase"
)

func main() {
	input := strcase.ToKebab(os.Args[1])
	fmt.Printf("test %s\n", input)
}