package main

import (
    "fmt"
    "testing"
    "reflect"
)

func TestConstructor(t *testing.T) {
    tree := NewFileTreeNode("/test/asdf", 100, false, []FileTreeNode{})
    fmt.Println(tree)
    expected := FileTreeNode{
        name: "/test/asdf",
        size: 100,
        cummulativeSize: 100,
        isDir: false,
        children: []FileTreeNode{},
    }
    if ! reflect.DeepEqual(*tree, expected) {
        expectedMsg := fmt.Sprintf("Expected %v and got %v", expected, tree)
        t.Error(expectedMsg)
    }
}
