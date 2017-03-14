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
    dirTest := NewFileTreeNode("test", 4096, true, nil)
    dirTest.addChild(node1)
    dirTest.addChild(node2)
    dirTest.addChild(node3)
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


// Verify recalculateCummulativeSize computes the correct size
func TestRecalculateCummulativeSize(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    dir2 := NewFileTreeNode("dir2", 4096, true, nil)
    file1 := NewFileTreeNode("file1", 800, false, nil)
    dir1.addChild(dir2)
    dir2.addChild(file1)
    // Structure: dir1/dir2/file1
    if dir1.cummulativeSize != 4096 {
        t.Error(fmt.Sprintf("Expected 4096 but got %v", dir1.cummulativeSize))
    }
    dir1.recalculateCummulativeSize()
    if dir1.cummulativeSize != 8992 {
        t.Error(fmt.Sprintf("Expected 8992 but got %v", dir1.cummulativeSize))
    }
    if dir2.cummulativeSize != 4896 {
        t.Error(fmt.Sprintf("Expected 4896 but got %v", dir2.cummulativeSize))
    }
}


func TestGetBiggestDirsAndFiles(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    dir2 := NewFileTreeNode("dir2", 4096, true, nil)
    dir3 := NewFileTreeNode("dir3", 4096, true, nil)
    dir4 := NewFileTreeNode("dir4", 4096, true, nil)
    file1 := NewFileTreeNode("file1", 8000, false, nil)
    file2 := NewFileTreeNode("file2", 10000, false, nil)
    dir1.addChild(dir2)
    dir1.addChild(dir3)
    dir1.addChild(dir4)
    dir2.addChild(file1)
    dir3.addChild(file2)
    dir1.recalculateCummulativeSize()
    // Structure:
    // dir1
    //   dir2
    //     file1
    //   dir3
    //     file2
    //   dir4
    dirs, files := dir1.getBiggestDirsAndFiles(5)
    expectedDirs := []*FileTreeNode{dir1, dir3, dir2, dir4}
    if ! reflect.DeepEqual(dirs, expectedDirs) {
        t.Error(fmt.Sprintf("Expected %v but got %v", expectedDirs, dirs))
    }
    expectedFiles := []*FileTreeNode{file2, file1}
    if ! reflect.DeepEqual(files, expectedFiles) {
        t.Error(fmt.Sprintf("Expected %v but got %v", expectedFiles, files))
    }
}


// Verify the expected criticalPath is returned
func TestCriticalPath(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    dir2 := NewFileTreeNode("dir2", 4096, true, nil)
    dir3 := NewFileTreeNode("dir3", 4096, true, nil)
    dir4 := NewFileTreeNode("dir4", 4096, true, nil)
    file1 := NewFileTreeNode("file1", 8000, false, nil)
    file2 := NewFileTreeNode("file2", 10000, false, nil)
    dir1.addChild(dir2)
    dir1.addChild(dir3)
    dir1.addChild(dir4)
    dir2.addChild(file1)
    dir3.addChild(file2)
    dir1.recalculateCummulativeSize()
    // Structure:
    // dir1
    //   dir2
    //     file1
    //   dir3
    //     file2
    //   dir4
    cp := dir1.criticalPath()
    expected := []*FileTreeNode{dir1, dir3, file2}
    if ! reflect.DeepEqual(cp, expected) {
        t.Error(fmt.Sprintf("Expected %v but got %v", expected, cp))
    }
}


// Verify only files/dirs with at least 30% of the parent's size are
// part of the critical path
func TestCriticalPath2(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    dir2 := NewFileTreeNode("dir2", 4096, true, nil)
    dir1.addChild(dir2)
    for i := 0; i < 10; i++ {
        file := NewFileTreeNode(
            fmt.Sprintf("%s%d", "file", i + 1),
            8000,
            false,
            nil,
        )
        dir2.addChild(file)
    }
    dir1.recalculateCummulativeSize()
    // Structure:
    // dir1
    //   dir2
    //     file1
    //     file2
    //     ...
    //     file10
    cp := dir1.criticalPath()
    expected := []*FileTreeNode{dir1, dir2}  // No files included
    if ! reflect.DeepEqual(cp, expected) {
        t.Error(fmt.Sprintf("Expected %v but got %v", expected, cp))
    }
}


// Verify the full path of a node is returned correctly
func TestFullPath(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    dir2 := NewFileTreeNode("dir2", 4096, true, nil)
    file1 := NewFileTreeNode("file1", 8000, false, nil)
    dir1.addChild(dir2)
    dir2.addChild(file1)
    // Structure:
    // dir1
    //   dir2
    //     file1
    path := file1.fullPath()
    expected := "dir1/dir2/file1"
    if path != expected {
        t.Fatal(fmt.Sprintf("Expected %s and got %s", expected, path))
    }
}


// Verify addChild adds the node to children[] and sets the parent
// attribute of the node
func TestAddChild(t *testing.T) {
    dir1 := NewFileTreeNode("dir1", 4096, true, nil)
    file1 := NewFileTreeNode("file1", 8000, false, nil)
    dir1.addChild(file1)
    expected := []*FileTreeNode{file1}
    if ! reflect.DeepEqual(dir1.children, expected) {
        t.Error(
            fmt.Sprintf(
                "Expected %v but got %v",
                expected,
                dir1.children,
            ),
        )
    }
    if file1.parent != dir1 {
        t.Error(
            fmt.Sprintf(
                "Expected %v but got %v",
                dir1,
                file1.parent,
            ),
        )
    }
}
