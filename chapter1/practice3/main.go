package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		s, sep := "", ""
		for _, a := range os.Args {
			s += s + sep + a
			sep = " "
		}
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs\n", secs)

	start = time.Now()
	for i := 0; i < 1000; i++ {
		strings.Join(os.Args, " ")
	}
	secs = time.Since(start).Seconds()
	fmt.Printf("%.2fs\n", secs)
}
