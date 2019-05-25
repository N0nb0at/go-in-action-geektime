package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World !")
	//os.Exit(-1)
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello ", os.Args[1])
	}
}
