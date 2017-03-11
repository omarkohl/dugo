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
    return &FileTreeNode{}
}
