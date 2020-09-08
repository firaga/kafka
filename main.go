package main

import (
	"flag"
	"fmt"
)

const (
	ModP = "p"
	ModC = "c"
	ModM = "m"
)

func main() {
	var mod string
	flag.StringVar(&mod, "mod", ModM, "intut mod")
	flag.Parse()
	fmt.Println("mod :", mod)
	//os.Exit(0)
	switch mod {
	case ModP:
		producerTest()
	case ModC:
		consumerTest()
	default:
		metadataTest()
	}
}
