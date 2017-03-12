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
