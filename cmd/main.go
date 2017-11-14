package main

import (
	"fmt"

	"github.com/zdq0394/simpleci/simpleci/server"
)

func main() {
	fmt.Println("Hello World")
	server.Start("../conf/config.conf")
}
