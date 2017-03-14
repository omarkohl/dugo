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

func humanizeSize(bytes int64) string {
    if bytes < 1024 {
        return strconv.FormatInt(bytes, 10) + " bytes"
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
    flag.Parse()
    root := flag.Arg(0)
    tree := buildTree(root)
    dirs, files := tree.getBiggestDirsAndFiles(11)
    // Ignore the first dir because it is the root dir. TODO take into account
    // the case of several parameters and file instead of dir as parameter
    fmt.Println("Top 10 directories are:")
    for _, d := range dirs[1:] {
        fmt.Printf("%s    (%s)\n", d.name, humanizeSize(d.cummulativeSize))
    }
    fmt.Println("Top 10 files are:")
    for _, f := range files {
        fmt.Printf("%s    (%s)\n", f.name, humanizeSize(f.cummulativeSize))
    }
}
