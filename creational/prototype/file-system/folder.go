package main

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f Folder) clone() Inode {
	cloneFolder := Folder{
		name: f.name,
	}
	var tempChildren []Inode
	for _, child := range f.children {
		childCopy := child.clone()
		tempChildren = append(tempChildren, childCopy)
	}

	cloneFolder.children = tempChildren

	return cloneFolder
}
