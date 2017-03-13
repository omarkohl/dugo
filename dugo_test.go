package main

import (
    "testing"
    "fmt"
    "path/filepath"
    "os"
    "time"
)

type fileInfoMock struct {
    size int64
    isDir bool
}

func (fs fileInfoMock) Size() int64        { return fs.size }
func (fs fileInfoMock) Mode() os.FileMode  { return 777 }
func (fs fileInfoMock) ModTime() time.Time { return time.Now() }
func (fs fileInfoMock) Sys() interface{}   { return nil }
func (fs fileInfoMock) IsDir() bool        { return fs.isDir }
func (fs fileInfoMock) Name() string       { return "mock" }

// Verify the buildTree function returns a directory tree
func TestBuildTree(t *testing.T) {
    // Mock the filepath.Walk function
    oldWalk := Walk
    defer func() { Walk = oldWalk }()
    Walk = func (root string, walkFn filepath.WalkFunc) error {
        walkFn(
            "example_dir",
            fileInfoMock{size: 4096, isDir: true},
            nil,
        )
        walkFn(
            "example_dir/subdir",
            fileInfoMock{size: 4096, isDir: true},
            nil,
        )
        walkFn(
            "example_dir/subdir/file.txt",
            fileInfoMock{size: 308, isDir: false},
            nil,
        )
        return nil
    }
    tree := buildTree("example_dir")
    if tree == nil {
        t.Fatal("Expected a tree and got nil")
    }
    fmt.Println(tree)
    if tree.name != "example_dir" {
        t.Fatal(fmt.Sprintf("Expected 'example_dir' and got %v", tree.name))
    }
    if len(tree.children) != 1 {
        t.Fatal(fmt.Sprintf("Expected 1 child and got %v", len(tree.children)))
    }
    if len(tree.children["subdir"].children) != 1 {
        t.Fatal(fmt.Sprintf("Expected 1 child and got %v", len(tree.children["subdir"].children)))
    }
    if tree.children["subdir"].children["file.txt"].name != "file.txt" {
        t.Fatal(fmt.Sprintf("Expected 'file.txt' and got %v", tree.children["subdir"].children["file.txt"].name))
    }
    if tree.cummulativeSize != 8500 {
        t.Fatal(fmt.Sprintf("Expected 8500 and got %v", tree.cummulativeSize))
    }
}
