package main

type FileTreeNode struct {
    name            string
    size            int64
    cummulativeSize int64
    isDir           bool
    children        []FileTreeNode
}

func NewFileTreeNode(
    name string,
    size int64,
    isDir bool,
    children []FileTreeNode,
) *FileTreeNode {
    return &FileTreeNode{
        name: name,
        size: size,
        cummulativeSize: size,
        isDir: isDir,
        children: children,
    }
}
