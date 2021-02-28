package main

import (
	"fmt"
	"flag"
)

func main() {

	rootDir := flag.String("d", ".", "Hosted directory (Default: '.' )")

	fmt.Println("We go host some stuff at " + *rootDir)
}
