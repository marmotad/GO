package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	fmt.Println(os.Executable())
}
