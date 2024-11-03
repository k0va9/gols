package main

import (
	"flag"
	"os"
	"strings"
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

type FileInfo struct {
	name string
}

func walk(target string, allFlag bool) []FileInfo {
	var result []FileInfo

	files, _ := os.ReadDir(target)

	for _, f := range files {
		ent, _ := f.Info()
		if !allFlag && strings.HasPrefix(ent.Name(), ".") {
			continue
		}

		info := FileInfo{
			name: ent.Name(),
		}

		result = append(result, info)
	}

	return result
}

func main() {
	flag.Parse()
}
