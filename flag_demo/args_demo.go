package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args是一[]string
	fmt.Println(os.Args)
	if len(os.Args) > 0 {
		for index, value := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, value)
		}
	}
}
