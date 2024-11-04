package main

import (
	"flag"
	"fmt"
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
	permission string
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
			permission: ent.Mode().String(),
		}

		result = append(result, info)
	}

	return result
}

func printEntry(ent FileInfo) {
	fmt.Printf("%s %s \n",ent.permission,ent.name)
}

func main() {
	flag.Parse()
	target := "."
	if len(flag.Arg(0)) > 0 {
		target = flag.Arg(0)
	}
	for _, item := range walk(target, allFlag) {
		printEntry(item)
	}
	print("\n")
}
