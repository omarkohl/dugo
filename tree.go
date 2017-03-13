package main

import (
    "os"
    "strings"
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

func (n *FileTreeNode) findDescendant(pathList []string) *FileTreeNode {
    if pathList[0] != n.name {
        // TODO Error
        return nil
    }
    if len(pathList) == 1 {
        return n
    }
    return n.children[pathList[1]].findDescendant(pathList[1:])
}

func (n *FileTreeNode) findDescendantStrPath(path string) *FileTreeNode {
    pathList := strings.Split(path, string(os.PathSeparator))
    return n.findDescendant(pathList)
}
