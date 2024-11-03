package main

import (
	"flag"
)

var allFlag bool
var helpFlag bool
var longFlag bool

func init() {
	flag.BoolVar(&allFlag, "all", false, "")
	flag.BoolVar(&allFlag, "a", false, "")
	flag.BoolVar(&longFlag, "l", false, "")
	flag.BoolVar(&helpFlag, "help", false, "")
}

func main() {
	flag.Parse()
}
