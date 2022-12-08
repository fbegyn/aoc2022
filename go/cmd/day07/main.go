package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fbegyn/aoc2022/go/helpers"
)

type File struct {
	name string
	size int
}

func (f *File) Size() int {
	return f.size
}

type Dir struct {
	name   string
	parent *Dir
	child  []*Dir
	files  []*File
}

func (d *Dir) AddFile(f *File) bool {
	for _, c := range d.files {
		if c.name == f.name {
			return false
		}
	}
	d.files = append(d.files, f)
	return true
}

func (d *Dir) AddNewChild(n string) *Dir {
	for _, c := range d.child {
		if c.name == n {
			return c
		}
	}
	d.child = append(d.child, &Dir{
		name:   n,
		parent: d,
		child:  []*Dir{},
		files:  []*File{},
	})
	return d.child[len(d.child)-1]
}

func (d *Dir) Size() (size int) {
	for _, c := range d.child {
		size += c.Size()
	}
	for _, c := range d.files {
		size += c.Size()
	}
	return size
}

func (d *Dir) FindChild(n string) *Dir {
	for _, c := range d.child {
		if c.name == n {
			return c
		}
	}
	return nil
}

func (d *Dir) CheckSmaller(thresh int) (t []*Dir) {
	if d.Size() <= thresh {
		t = append(t, d)
	}
	for _, c := range d.child {
		t = append(t, c.CheckSmaller(thresh)...)
	}
	return
}

func (d *Dir) CheckLarger(thresh int) (t []*Dir) {
	if d.Size() >= thresh {
		t = append(t, d)
	}
	for _, c := range d.child {
		t = append(t, c.CheckLarger(thresh)...)
	}
	return
}

func main() {
	inputFile := os.Args[1]
	input := helpers.InputToLines(inputFile)

	totalSize := 70000000
	spaceNeeded := 30000000

	root := Dir{
		name:   "/",
		parent: nil,
		child:  []*Dir{},
		files:  []*File{},
	}
	wd := &root

	for _, l := range input {
		fields := strings.Fields(l)
		switch fields[0] {
		case "$":
			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "..":
					wd = wd.parent
				case "/":
					wd = &root
				default:
					wd = wd.FindChild(fields[2])
				}
			}
		default:
			switch fields[0] {
			case "dir":
				wd.AddNewChild(fields[1])
			default:
				file := File{
					name: fields[1],
					size: helpers.Atoi(fields[0]),
				}
				wd.AddFile(&file)
			}
		}
	}

	spaceLeft := totalSize - root.Size()
	toDelete := spaceNeeded - spaceLeft
	smallest := &root
	sum := 0
	for _, d := range root.CheckSmaller(100000) {
		sum += d.Size()
	}

	for _, d := range root.CheckLarger(toDelete) {
		if d.Size() < smallest.Size() {
			smallest = d
		}
	}

	fmt.Printf("Sum of the directories below threshold: %d\n", sum)
	fmt.Printf("Smallest directory to delete has size : %d\n", smallest.Size())
}
