package main

import (
	"bufio"
	"directory"
	"fmt"
	"os"
	"parser"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	dir := &directory.Directory{}
	dir.Dir = make([]*directory.Directory, 0)
	dir.Name = "/"
	dir.Files = make([]*directory.File, 0)
	actualDir := dir
	for scanner.Scan() {
		line := scanner.Text()
		dir, actualDir = parser.ParseCommand(line, dir, actualDir)

	}
	totalSize := dir.Size()
	neededSize := totalSize - 70000000 + 30000000
	dirs := getSizeDirectoriesGreaterThan(dir, neededSize)
	directory.SortDirs(dirs)
	fmt.Println(dirs[0].Size())
}

func getSizeDirectoriesGreaterThan(dir *directory.Directory, size int) []*directory.Directory {
	result := []*directory.Directory{}
	for _, directory := range dir.Dir {
		if directory.Size() > size {
			result = append(result, directory)
		}
		if len(directory.Dir) > 0 {
			result = append(result, getSizeDirectoriesGreaterThan(directory, size)...)
		}
	}
	return result

}
func totalSize(dirs []directory.Directory) int {
	total := 0
	for _, dir := range dirs {
		total += dir.Size()
	}
	return total
}
