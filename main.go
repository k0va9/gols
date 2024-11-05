package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var allFlag bool
var helpFlag bool
var longFlag bool
var ownerWith int
var groupWidth int
var blocksizeWidth int

func init() {
	flag.BoolVar(&allFlag, "all", false, "")
	flag.BoolVar(&allFlag, "a", false, "")
	flag.BoolVar(&longFlag, "l", false, "")
	flag.BoolVar(&helpFlag, "help", false, "")
}

type FileInfo struct {
	name       string
	permission string
	owner      string
	group      string
	size       string
	date	   time.Time
	isDir      bool
	nlink      uint64
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
			name:       ent.Name(),
			permission: ent.Mode().String(),
			owner:      getOwner(ent),
			group:      getGroup(ent),
			size:       strconv.FormatInt(ent.Size(), 10),
			date: 	    ent.ModTime(),
			isDir:      ent.IsDir(),
			nlink:      getNlink(ent),
		}

		// caluclate padding width
		if ownerWith < len(info.owner) {
			ownerWith = len(info.owner)
		}
		if groupWidth < len(info.group) {
			groupWidth = len(info.group)
		}
		if blocksizeWidth < len(info.size) {
			blocksizeWidth = len(info.size)
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

func getGroup(file os.FileInfo) string {
	if s, ok := file.Sys().(*syscall.Stat_t); ok {
		gid := strconv.Itoa(int(s.Gid))
		groupInfo, err := user.LookupGroupId(gid)
		if err == nil {
			return groupInfo.Name
		}

	}
	return ""
}

func printEntName(ent FileInfo) {
	if ent.isDir {
		fmt.Printf("\x1b[34m%s\x1b[0m  ", ent.name)
	} else {
		fmt.Printf("%s ", ent.name)
	}
}

func printOwnerInfo(ent FileInfo) {
	fmt.Printf("%-*s %-*s ", ownerWith, ent.owner, groupWidth, ent.group)
}

func printPermission(ent FileInfo) {
	fmt.Print(ent.permission + " ")
}

func printNlink(ent FileInfo) {
	fmt.Printf("%d ", ent.nlink)
}
func printSize(ent FileInfo) {
	fmt.Printf("%*s ", blocksizeWidth, ent.size)
}
func printModDate(ent FileInfo) {
	fmt.Printf("%s %2d %02d:%02d ", ent.date.Month().String()[:3], ent.date.Day(), ent.date.Hour(), ent.date.Minute())
}

func main() {
	flag.Parse()
	target := "."
	if len(flag.Arg(0)) > 0 {
		target = flag.Arg(0)
	}
	for _, item := range walk(target, allFlag) {
		if longFlag {
			printPermission(item)
			printNlink(item)
			printOwnerInfo(item)
			printSize(item)
			printModDate(item)
			printEntName(item)
			print("\n")
		} else {
			printEntName(item)
			print("  ")
		}
	}
	print("\n")
}
