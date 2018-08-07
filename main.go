package main

import (
	"fmt"
	"log"
)

func main() {
	output, outdated, err := goDep()
	if err != nil {
		log.Fatalf("%s%s", output, err)
	}
	fmt.Println("There are", outdated, "outdated deps")
	if outdated > 0 {
		fmt.Println(string(output))
	}
}
