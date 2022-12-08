package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name     string
	subdirs  map[string]Dir
	files    map[string]File
	fullPath string
}

var MemoDirSize map[string]int = map[string]int{}

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	path := []string{"/"}
	filestruct := mkdir("/", path)

	currentDir := getDirFromPath(filestruct, path)

	for _, line := range inputSplit {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "$") {
			// This is a command
			if line == "$ cd /" {
				path := []string{"/"}
				currentDir = getDirFromPath(filestruct, path)
				continue
			}

			if line == "$ cd .." {
				path = path[:len(path)-1]
				currentDir = getDirFromPath(filestruct, path)
				continue
			}

			if line == "$ ls" {
				// well we are expecting this.. no?
				continue
			}

			// Else we have like "$ cd dirname"
			dirname := strings.Split(line, " ")[2]
			path = append(path, dirname)
			currentDir = getDirFromPath(filestruct, path)

			continue
		}

		// Handle Dirs
		if strings.HasPrefix(line, "dir") {
			dirName := strings.Split(line, "dir ")[1]
			currentDir.subdirs[dirName] = mkdir(dirName, path)
			continue
		}

		// Handle Files
		file := strings.Split(line, " ")
		sizeStr, filename := file[0], file[1]
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			log.Panic("bad parse", line)
		}
		currentDir.files[filename] = newFile(filename, size)
	}

	getTotalSizeOfDirectory(filestruct)

	ans := 0

	for _, size := range MemoDirSize {
		if size <= 100000 {
			ans += size
		}
	}
	fmt.Println(ans)

	// Part 2

	freeSpace := 70_000_000 - (MemoDirSize["///"]) // whatever
	need := 30_000_000 - freeSpace
	fmt.Println(need)
	smallest := math.MaxInt
	var deleteThisDir string
	for dir, size := range MemoDirSize {
		if size < need {
			continue
		}
		if size < smallest {
			deleteThisDir = dir
			smallest = size
		}
	}

	fmt.Println(deleteThisDir, smallest)
}

func mkdir(name string, path []string) Dir {
	fullPath := ""
	for _, dir := range path {
		fullPath += dir + "/"
	}
	return Dir{
		name:     name,
		subdirs:  map[string]Dir{},
		files:    map[string]File{},
		fullPath: fullPath + name,
	}
}

func getDirFromPath(root Dir, path []string) Dir {
	directory := root
	for _, subdir := range path[1:] {
		directory = directory.subdirs[subdir]
	}
	return directory
}

func newFile(name string, size int) File {
	return File{
		name: name,
		size: size,
	}
}

func getTotalSizeOfDirectory(dir Dir) int {
	size := 0
	memodSize, ok := MemoDirSize[dir.fullPath]
	if ok {
		return memodSize
	}
	for _, file := range dir.files {
		size += file.size
	}
	for _, subdir := range dir.subdirs {
		size += getTotalSizeOfDirectory(subdir)
	}
	MemoDirSize[dir.fullPath] = size
	return size
}
