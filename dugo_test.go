package main

import (
    "testing"
    "fmt"
    "path/filepath"
)

// Verify the buildTree function returns a directory tree
func TestBuildTree(t *testing.T) {
    // Mock the filepath.Walk function
    oldWalk := Walk
    defer func() { Walk = oldWalk }()
    Walk = func (root string, walkFn filepath.WalkFunc) error {
        fmt.Println(root)
        return nil
    }
    tree := buildTree("example_dir")
    if tree == nil {
        t.Error("Expected a tree and got nil")
    }
}
