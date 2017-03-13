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

// Test the findDescendant method
func TestFindDescendant(t *testing.T) {
    node1 := NewFileTreeNode("asdf", 100, false, nil)
    node2 := NewFileTreeNode("asdf2", 200, false, nil)
    node3 := NewFileTreeNode("asdf3", 200, false, nil)
    children := make(map[string]*FileTreeNode, 3)
    children[node1.name] = node1
    children[node2.name] = node2
    children[node3.name] = node3
    dirTest := NewFileTreeNode("test", 4096, true, children)
    dirTopLevel := NewFileTreeNode(
        "top-level",
        4096,
        true,
        map[string]*FileTreeNode{dirTest.name: dirTest},
    )
    result, err := dirTopLevel.findDescendantStrPath("top-level/test/asdf2")
    if node2 != result {
        expectedMsg := fmt.Sprintf("Expected %v and got %v", node2, result)
        t.Error(expectedMsg)
    }
    if err != nil {
        expectedMsg := fmt.Sprintf("Expected no error and got %v", err)
        t.Error(expectedMsg)
    }
}


// Verify findDescendant returns an error if top level path doesn't match
func TestFindDescendantError(t *testing.T) {
    node1 := NewFileTreeNode("asdf", 100, false, nil)
    result, err := node1.findDescendantStrPath("something-else")
    if err == nil {
        t.Error("Expected an error and got nil")
    }
    if result != nil {
        t.Error(fmt.Sprintf("Expected nil and instead got %v", result))
    }
}

// Verify findDescendant returns an error if a key doesn't exist
func TestFindDescendantError2(t *testing.T) {
    node1 := NewFileTreeNode("asdf", 100, false, nil)
    parent := NewFileTreeNode(
        "parent",
        4096,
        true,
        map[string]*FileTreeNode{node1.name: node1},
    )
    // Correct structure is parent/asdf . parent/wrong doesn't exist
    result, err := parent.findDescendantStrPath("parent/wrong")
    if err == nil {
        t.Error("Expected an error and got nil")
    }
    if result != nil {
        t.Error(fmt.Sprintf("Expected nil and instead got %v", result))
    }
}
