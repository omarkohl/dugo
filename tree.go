package main

import (
    "fmt"
    "os"
    "strings"
    "errors"
    "sort"
)

type FileTreeNode struct {
    name            string
    size            int64
    cummulativeSize int64
    isDir           bool
    children        map[string]*FileTreeNode
    parent          *FileTreeNode
}

// Constructor function for FileTreeNode. Use this instead of instantiating
// FileTreeNode directly to ensure the cummulativeSize is calculated correctly.
func NewFileTreeNode(
    name string,
    size int64,
    isDir bool,
    children map[string]*FileTreeNode,
    // TODO parameter parent
) *FileTreeNode {
    if children == nil {
        children = make(map[string]*FileTreeNode)
    }
    cummulativeSize := size
    for _, v := range children {
        cummulativeSize += v.size
    }
    return &FileTreeNode{
        name: name,
        size: size,
        cummulativeSize: cummulativeSize,
        isDir: isDir,
        children: children,  // TODO child.parent has to be set
        parent: nil,
    }
}

func (n *FileTreeNode) findDescendant(pathList []string) (*FileTreeNode, error) {
    if pathList[0] != n.name {
        errorMsg := fmt.Sprintf(
            "First element in pathList '%s' does not match %v",
            pathList[0],
            n,
        )
        return nil, errors.New(errorMsg)
    }
    if len(pathList) == 1 {
        return n, nil
    }
    child, ok := n.children[pathList[1]]
    if ! ok {
        return nil, errors.New(fmt.Sprintf("Node %v has no child %v", n, pathList[1]))
    }
    return child.findDescendant(pathList[1:])
}

func (n *FileTreeNode) findDescendantStrPath(path string) (*FileTreeNode, error) {
    pathList := strings.Split(path, string(os.PathSeparator))
    return n.findDescendant(pathList)
}

func (n *FileTreeNode) recalculateCummulativeSize() {
    total := n.size
    for _, v := range n.children {
        v.recalculateCummulativeSize()
        total += v.cummulativeSize
    }
    n.cummulativeSize = total
}

func min(a, b int) int {
    if (a < b) {
        return a
    }
    return b
}

// Type to implement sort.Interface
type BySize []*FileTreeNode
func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return a[i].cummulativeSize < a[j].cummulativeSize}

func (n *FileTreeNode) getBiggestDirsAndFiles(count int) ([]*FileTreeNode, []*FileTreeNode){
    var allDirs []*FileTreeNode
    var allFiles []*FileTreeNode
    var iterator func(node *FileTreeNode)
    iterator = func(node *FileTreeNode) {
        if node.isDir {
            allDirs = append(allDirs, node)
        } else {
            allFiles = append(allFiles, node)
        }
        for _, child := range node.children {
            iterator(child)
        }
    }
    iterator(n)
    sort.Sort(sort.Reverse(BySize(allDirs)))
    sort.Sort(sort.Reverse(BySize(allFiles)))
    return allDirs[:min(len(allDirs), count)], allFiles[:min(len(allFiles), count)]
}

func (n *FileTreeNode) criticalPath() []*FileTreeNode {
    res := []*FileTreeNode{n}
    var children []*FileTreeNode
    if len(n.children) == 0 {
        return res
    }
    for  _, value := range n.children {
        children = append(children, value)
    }
    sort.Sort(sort.Reverse(BySize(children)))
    if float64(children[0].cummulativeSize) / float64(n.cummulativeSize) >= 0.3 {
        return append(res, children[0].criticalPath()...)
    } else {
        return res
    }
}

func (n *FileTreeNode) addChild(child *FileTreeNode) {
    n.children[child.name] = child
    child.parent = n
}
