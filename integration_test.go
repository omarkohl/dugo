package main

import (
    "os/exec"
    "testing"
    "fmt"
)

func TestExample1(t *testing.T) {
    if testing.Short() {
        t.Skip()
    }
    out, err := exec.Command("dugo", "example1").Output()
    if err != nil {
        t.Error("Expected nil, got", err)
    }
    outStr := string(out)
    expectedOut := `
Top 10 directories are:

example1/dir2                                                          (12.4 kB)
example1/dir2/dir21                                                     (8.4 kB)
example1/dir1                                                           (8.0 kB)
example1/dir2/dir21/dir211                                              (4.4 kB)
example1/dir3                                                           (4.1 kB)
example1/dir1/dir11                                                     (4.0 kB)


Top 10 files are:

example1/dir2/dir21/dir211/fileB.txt                                 (379 bytes)
example1/dir3/fileC.txt                                               (52 bytes)
example1/dir1/dir11/fileA.txt                                         (17 bytes)


The critical path is:
example1/dir2/dir21/dir211/ (4.4 kB)
`
    if outStr != expectedOut {
        msg := fmt.Sprintf("Expected '%s' and got '%s'", expectedOut, outStr)
        t.Error(msg)
    }
}
