package main

import "fmt"

func main() {
	file1 := File{name: "file1"}
	file2 := File{name: "file2"}
	file3 := File{name: "file3"}

	folder1 := Folder{
		children: []Inode{file1},
		name:     "folder1",
	}

	folder2 := Folder{
		children: []Inode{folder1, file2, file3},
		name:     "folder2",
	}

	fmt.Println("\nPrint folder2 without clone:")
	folder2.print(" ")
	clonedFolder := folder2.clone()

	fmt.Println("\nPrint folder2 without clone:")
	clonedFolder.print(" ")
}
