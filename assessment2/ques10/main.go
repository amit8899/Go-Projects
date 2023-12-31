package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type matchResult struct {
	base  string
	index int
	ext   string
}

// var re = regexp.MustCompile("^(.*?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
// var replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "Whether or not this should be a real or dry run")
	flag.Parse()

	WalkDir := "..\\assessment2\\ques10\\sample"
	toRename := make(map[string][]string)
	filepath.Walk(WalkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)

		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, newPath, err.Error())
				}
			}
		}
	}
}

func match(filename string) (*matchResult, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern", filename)
	}
	return &matchResult{strings.Title(name), number, ext}, nil
}
