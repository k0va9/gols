package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
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
	owner string
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
			owner: getOwner(ent),
		}

		result = append(result, info)
	}

	return result
}

func getOwner(file os.FileInfo) string {
	if s, ok := file.Sys().(*syscall.Stat_t); ok {
		uid := strconv.Itoa(int(s.Uid))
		userInfo, err := user.LookupId(uid)
		if err == nil {
			return userInfo.Username
		}

	}
	return ""
}

var ownerWith = 0

func printEntry(ent FileInfo) {
	if ownerWith < len(ent.owner) {
		ownerWith = len(ent.owner)
	}
	fmt.Printf("%s %-*s %s\n",ent.permission,ownerWith,ent.owner,ent.name)
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
