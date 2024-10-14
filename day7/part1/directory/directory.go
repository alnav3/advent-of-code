package directory

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name   string
	Parent *Directory
	Dir    []*Directory
	Files  []*File
}

type File struct {
	Name string
	Size int
}

func (d *Directory) GoToDir(name string) *Directory {
	for _, dir := range d.Dir {
		if dir.Name == name {
			return dir
		}
	}
	return nil
}

func (d *Directory) String() string {
	dirStrs := []string{}
	for _, dir := range d.Dir {
		dirStrs = append(dirStrs, dir.String())
	}

	fileStrs := []string{}
	for _, file := range d.Files {
		fileStrs = append(fileStrs, file.String())
	}

	return fmt.Sprintf("Directory{Name: %s, Dir: [%s], Files: [%s]}", d.Name, strings.Join(dirStrs, ", "), strings.Join(fileStrs, ", "))
}

func (f *File) String() string {
	return fmt.Sprintf("File{Name: %s, Size: %d}", f.Name, f.Size)
}
func (d *Directory) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}
	for _, dir := range d.Dir {
		size += dir.Size()
	}
	return size
}
