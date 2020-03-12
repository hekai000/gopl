package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println(os.Args)

	for index, cmd := range os.Args[1:] {
		fmt.Println(index, cmd)
	}
}
