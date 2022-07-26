package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int

	flag.IntVar(&port, "p", 8060, "端口号")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)
	ginMain(addr)

	return
}
