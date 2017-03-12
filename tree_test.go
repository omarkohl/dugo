package main

import (
    "fmt"
    "testing"
    "reflect"
)

// Test that a single FileTreeNode is constructed correctly
func TestSingleTreeNode(t *testing.T) {
    node := NewFileTreeNode("asdf", 100, false, nil)
    expected := FileTreeNode{
        name: "asdf",
        size: 100,
        cummulativeSize: 100,
        isDir: false,
        children: make(map[string]*FileTreeNode),
    }
    if ! reflect.DeepEqual(*node, expected) {
        expectedMsg := fmt.Sprintf("Expected %v and got %v", expected, *node)
        t.Error(expectedMsg)
    }
}

// Test that a tree is constructed correctly and sizes are added up as expected
func TestTree(t *testing.T) {
    node1 := NewFileTreeNode("asdf", 100, false, nil)
    node2 := NewFileTreeNode("asdf2", 200, false, nil)
    node3 := NewFileTreeNode("asdf3", 200, false, nil)
    children := make(map[string]*FileTreeNode, 3)
    children[node1.name] = node1
    children[node2.name] = node2
    children[node3.name] = node3
    nodeParent := NewFileTreeNode("test", 4096, true, children)
    expected := FileTreeNode{
        name: "test",
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
