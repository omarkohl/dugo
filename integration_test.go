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
    expectedOut := `Top 10 directories are:
dir2    (12.4 kB)
dir1    (8.0 kB)
dir3    (4.1 kB)

Top 10 files are:
fileB   (379 bytes)
fileC   (52 bytes)
fileA   (17 bytes)

The critical path is: dir2/dir21/dir211 (4.4 kB)
`
    if outStr != expectedOut {
        msg := fmt.Sprintf("Expected '%s' and got", expectedOut)
        t.Error(msg, outStr)
    }
}
