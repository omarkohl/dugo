package main

import (
    "fmt"
    "os"
    "strings"
    "errors"
)

type FileTreeNode struct {
    name            string
    size            int64
    cummulativeSize int64
    isDir           bool
    children        map[string]*FileTreeNode
}

// Constructor function for FileTreeNode. Use this instead of instantiating
// FileTreeNode directly to ensure the cummulativeSize is calculated correctly.
func NewFileTreeNode(
    name string,
    size int64,
    isDir bool,
    children map[string]*FileTreeNode,
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
        children: children,
    }
}

func (n *FileTreeNode) findDescendant(pathList []string) (*FileTreeNode, error) {
    if pathList[0] != n.name {
        return nil, errors.New("First element in pathList doesn't match this node")
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

func (n *FileTreeNode) getBiggestDirsAndFiles(count int) ([]*FileTreeNode, []*FileTreeNode){
    topDirs := make([]*FileTreeNode, count)
    topFiles := make([]*FileTreeNode, count)
    return topDirs, topFiles
}
