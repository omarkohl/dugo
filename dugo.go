package main

import (
    "fmt"
    "path/filepath"
    "os"
    "flag"
    "strings"
)

var Walk = filepath.Walk

func buildTree(root string) *FileTreeNode {
    var tree *FileTreeNode
    visit := func (path string, f os.FileInfo, err error) error {
        if path == root {
            tree = NewFileTreeNode(f.Name(), f.Size(), f.IsDir(), nil)
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

func printTree(tree *FileTreeNode, indent int) {
    fmt.Printf("%s%s\n", strings.Repeat(" ", indent), tree.name)
    for _, child := range tree.children {
        printTree(child, indent + 2)
    }
}

func main() {
    fmt.Println("This is dugo (Disk Usage with Go)!")
    flag.Parse()
    root := flag.Arg(0)
    tree := buildTree(root)
    printTree(tree, 0)
}
