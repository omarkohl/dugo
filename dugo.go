package main

import (
    "fmt"
    "path/filepath"
    "os"
    "flag"
    "strings"
    "strconv"
    "math"
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
            node := NewFileTreeNode(f.Name(), f.Size(), f.IsDir(), nil)
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

func humanizeSize(bytes int) string {
    if bytes < 1024 {
        return strconv.Itoa(bytes) + " bytes"
    }
    units := []string{"kB", "MB", "GB", "TB"}
    for i, u := range units {
        res := float64(bytes) / math.Pow(1024.0, float64(i + 1))
        if res < 1024.0 || u == "TB" {
            return strconv.FormatFloat(res, 'f', 1, 64) + " " + u
        }
    }
    return "ERROR!"  // Should never happen
}

func main() {
    fmt.Println("This is dugo (Disk Usage with Go)!")
    flag.Parse()
    root := flag.Arg(0)
    tree := buildTree(root)
    dirs, files := tree.getBiggestDirsAndFiles(11)
    // Ignore the first dir because it is the root dir. TODO take into account
    // the case of several parameters and file instead of dir as parameter
    fmt.Println("Top 10 directories are:")
    for _, d := range dirs[1:] {
        fmt.Println(d.name)
    }
    fmt.Println("Top 10 files are:")
    for _, f := range files {
        fmt.Println(f.name)
    }
}
