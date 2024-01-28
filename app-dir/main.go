package main

import (
	toolkit "toolkitmodule"
)

func main() {
	var tools toolkit.Tools
	tools.CreateDirIfNotExist("./test-dir")
}
