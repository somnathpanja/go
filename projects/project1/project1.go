package main

import (
	"fmt"

	"github.com/somnathpanja/learngo/modules/mod1"
)

func main() {
	reply, err := mod1.Hello("Somnath")
	fmt.Println(reply, err)
}
