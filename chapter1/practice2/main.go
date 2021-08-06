package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		fmt.Printf("%d\t%v\n", i, a)
	}
}
