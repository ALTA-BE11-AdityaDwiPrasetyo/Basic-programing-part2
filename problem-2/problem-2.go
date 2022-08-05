package main

import (
	"fmt"
)
func main() {
func list(a *gami.Asterisk) {
	fmt.Println("Enter conference name: ")
	var conf string
	fmt.Scanf("%s\n", &conf)

	ml, err := a.GetConfbridgeList(conf)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ml)
}
}
