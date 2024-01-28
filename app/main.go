package main

import (
	"fmt"
	toolkit "toolkitmodule"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
