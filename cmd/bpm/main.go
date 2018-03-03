package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "github.com/davethegr8/bad-package-manager/pkg/bpm"
)

var (
    File      string
    // V, VV      bool
)

type flagArray []string

func init() {
    // -f file
    flag.StringVar(&File, "f", "dependencies.json", "dependencies file")
    // flag.BoolVar(&V, "v", false, "verbose")
    // flag.BoolVar(&VV, "vv", false, "very verbose")
}

func errorMessage(message string) bool {
    fmt.Fprintln(os.Stderr, message)
    return false
}

func main() {
    fmt.Println("hello world")

    flag.Parse()

    ok := true

    File := resolvePath(File)

    if _, err := os.Stat(File); os.IsNotExist(err) {
        ok = errorMessage("ERROR: dependencies file does not exist")
    }

    if !ok {
        fmt.Println("Usage: bpm [-f filename]")
        flag.PrintDefaults()
        os.Exit(1)
    }

    bpm.Process(File)
}

func resolvePath(file string) (absfile string) {
    fmt.Println(file)

    if filepath.IsAbs(file) {
        return file
    }

    absfile, err := filepath.Abs(file)
    if err != nil {
        panic(err)
    }

    return absfile
}
