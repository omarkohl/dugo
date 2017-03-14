package main

import (
    "testing"
    "fmt"
    "path/filepath"
    "os"
    "time"
)

type fileInfoMock struct {
    name string
    size int64
    isDir bool
}

func (fs fileInfoMock) Size() int64        { return fs.size }
func (fs fileInfoMock) Mode() os.FileMode  { return 777 }
func (fs fileInfoMock) ModTime() time.Time { return time.Now() }
func (fs fileInfoMock) Sys() interface{}   { return nil }
func (fs fileInfoMock) IsDir() bool        { return fs.isDir }
func (fs fileInfoMock) Name() string       { return fs.name }

// Verify the buildTree function returns a directory tree
func TestBuildTree(t *testing.T) {
    // Mock the filepath.Walk function
    oldWalk := Walk
    defer func() { Walk = oldWalk }()
    Walk = func (root string, walkFn filepath.WalkFunc) error {
        walkFn(
            "example_dir",
            fileInfoMock{name: "example_dir", size: 4096, isDir: true},
            nil,
        )
        walkFn(
            "example_dir/subdir",
            fileInfoMock{name: "subdir", size: 4096, isDir: true},
            nil,
        )
        walkFn(
            "example_dir/subdir/file.txt",
            fileInfoMock{name: "file.txt", size: 308, isDir: false},
            nil,
        )
        return nil
    }
    tree := buildTree("example_dir")
    if tree == nil {
        t.Fatal("Expected a tree and got nil")
    }
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


func assertEqualStr(t *testing.T, result, expected string) {
    if result != expected {
        t.Error(fmt.Sprintf("Expected '%s' and got '%s'", expected, result))
    }
}

func TestHumanizeSize(t *testing.T) {
    result := humanizeSize(500)
    expected := "500 bytes"
    assertEqualStr(t, result, expected)

    result = humanizeSize(1024)
    expected = "1.0 kB"
    assertEqualStr(t, result, expected)

    result = humanizeSize(1050)
    expected = "1.0 kB"
    assertEqualStr(t, result, expected)

    result = humanizeSize(21058)
    expected = "20.6 kB"
    assertEqualStr(t, result, expected)

    result = humanizeSize(10000068)
    expected = "9.5 MB"
    assertEqualStr(t, result, expected)

    result = humanizeSize(10002000068)
    expected = "9.3 GB"
    assertEqualStr(t, result, expected)

    result = humanizeSize(10002000068000)
    expected = "9.1 TB"
    assertEqualStr(t, result, expected)
}
