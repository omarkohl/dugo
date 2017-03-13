package main

import (
    "fmt"
    "path/filepath"
    "os"
)

var Walk = filepath.Walk

func buildTree(root string) *FileTreeNode {
    var tree *FileTreeNode
    visit := func (path string, f os.FileInfo, err error) error {
        fmt.Printf("Visiting %s with size %d\n", path, f.Size())
        if path == root {
            tree = NewFileTreeNode(path, f.Size(), f.IsDir(), nil)
        } else {
            dir := filepath.Dir(path)
            base := filepath.Base(path)
            descendant, err := tree.findDescendantStrPath(dir)
            if err != nil {
                fmt.Printf("Error! %v\n", err)
            }
            node := NewFileTreeNode(base, f.Size(), f.IsDir(), nil)
            descendant.children[base] = node
        }
        return nil
    }
    Walk(root, visit)
    tree.recalculateCummulativeSize()
    return tree
}

func main() {
    fmt.Println("This is dugo (Disk Usage with Go)!")
}
