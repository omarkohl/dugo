package main

import (
    "fmt"
    "testing"
    "reflect"
)

// Test that a single FileTreeNode is constructed correctly
func TestSingleTreeNode(t *testing.T) {
    node := NewFileTreeNode("/test/asdf", 100, false, []FileTreeNode{})
    expected := FileTreeNode{
        name: "/test/asdf",
        size: 100,
        cummulativeSize: 100,
        isDir: false,
        children: []FileTreeNode{},
    }
    if ! reflect.DeepEqual(*node, expected) {
        expectedMsg := fmt.Sprintf("Expected %v and got %v", expected, node)
        t.Error(expectedMsg)
    }
}

// Test that a tree is constructed correctly and sizes are added up as expected
func TestTree(t *testing.T) {
    node1 := NewFileTreeNode("/test/asdf", 100, false, []FileTreeNode{})
    node2 := NewFileTreeNode("/test/asdf2", 200, false, []FileTreeNode{})
    node3 := NewFileTreeNode("/test/asdf3", 200, false, []FileTreeNode{})
    children := make([]FileTreeNode, 3)
    children[0] = *node1
    children[1] = *node2
    children[2] = *node3
    nodeParent := NewFileTreeNode("/test", 4096, true, children)
    expected := FileTreeNode{
        name: "/test",
        size: 4096,
        cummulativeSize: 4596,
        isDir: true,
        children: children,
    }
    if ! reflect.DeepEqual(*nodeParent, expected) {
        expectedMsg := fmt.Sprintf("Expected %v and got %v", expected, nodeParent)
        t.Error(expectedMsg)
    }
}
