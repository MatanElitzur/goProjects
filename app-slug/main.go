package main

import (
	"log"
	toolkit "toolkitmodule"
)

func main() {
	toSlug := "now is it the time 123"
	var tools toolkit.Tools
	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugified)
}
