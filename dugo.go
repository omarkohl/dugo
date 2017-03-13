package main

import (
    "fmt"
    "path/filepath"
)

var Walk = filepath.Walk

func buildTree(root string) *FileTreeNode {
    return NewFileTreeNode(root, 111, true, nil)
}

func main() {
    fmt.Println("This is dugo (Disk Usage with Go)!")
}
