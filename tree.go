package main

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
