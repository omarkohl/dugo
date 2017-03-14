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
var Abs = filepath.Abs

func buildTree(root string) (*FileTreeNode, error) {
    var err error
    root, err = Abs(root)
    if err != nil {
        return nil, err
    }
    basePath := filepath.Dir(root) + string(os.PathSeparator)
    var tree *FileTreeNode
    visit := func (path string, f os.FileInfo, err error) error {
        if path == root {
            tree = NewFileTreeNode(f.Name(), f.Size(), f.IsDir(), nil)
        } else {
            path = strings.TrimPrefix(path, basePath)
            dir := filepath.Dir(path)
            descendant, err := tree.findDescendantStrPath(dir)
            if err != nil {
                return err
            }
            node := NewFileTreeNode(f.Name(), f.Size(), f.IsDir(), nil)
            descendant.addChild(node)
        }
        return nil
    }
    err = Walk(root, visit)
    if err != nil {
        return nil, err
    }
    tree.recalculateCummulativeSize()
    return tree, nil
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
    tree, err := buildTree(root)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    dirs, files := tree.getBiggestDirsAndFiles(11)
    // Ignore the first dir because it is the root dir. TODO take into account
    // the case of several parameters and file instead of dir as parameter
    fmt.Println("\nTop 10 directories are:\n")
    for _, d := range dirs[1:] {
        fmt.Printf("%-60s%20s\n", d.fullPath(), "(" + humanizeSize(d.cummulativeSize) + ")")
    }
    fmt.Println("\n\nTop 10 files are:\n")
    for _, f := range files {
        fmt.Printf("%-60s%20s\n", f.fullPath(), "(" + humanizeSize(f.cummulativeSize) + ")")
    }
    criticalPath := dirs[0].criticalPath()
    strCp := ""
    for _, v := range criticalPath {
        strCp += v.name
        if v.isDir {
            strCp += "/"
        }
    }
    lastElement := criticalPath[len(criticalPath) - 1]
    fmt.Printf(
        "\n\nThe critical path is:\n%s (%s)\n",
        strCp,
        humanizeSize(lastElement.cummulativeSize),
    )
}
