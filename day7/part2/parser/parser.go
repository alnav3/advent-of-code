package parser

import (
	"directory"
	"strconv"
	"strings"
)

func ParseCommand(line string, dir *directory.Directory, actualDir *directory.Directory) (*directory.Directory, *directory.Directory) {
	parts := strings.Split(line, " ")
	if parts[0] == "$" {
		return executeCommand(parts, dir, actualDir)
	} else {
		return createStructure(parts, dir, actualDir)
	}
}

func executeCommand(parts []string, dir *directory.Directory, actualDir *directory.Directory) (*directory.Directory, *directory.Directory) {
	command := parts[1]
	if command == "cd" {
		Name := parts[2]
		switch Name {
		case "..":
			actualDir = actualDir.Parent
		case "/":
			actualDir = dir
		default:
			actualDir = actualDir.GoToDir(Name)
		}
	}
	return dir, actualDir
}

func createStructure(parts []string, dir *directory.Directory, actualDir *directory.Directory) (*directory.Directory, *directory.Directory) {
	if parts[0] == "dir" {
		actualDir.Dir = append(actualDir.Dir, &directory.Directory{Name: parts[1], Parent: actualDir, Dir: make([]*directory.Directory, 0),
			Files: make([]*directory.File, 0)})
	} else {
		size, _ := strconv.Atoi(parts[0])
		actualDir.Files = append(actualDir.Files, &directory.File{Name: parts[1], Size: size})
	}
	return dir, actualDir
}
