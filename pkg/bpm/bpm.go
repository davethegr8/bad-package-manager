package bpm

import (
    "encoding/json"
    // "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// type Dependencies struct {
//     Require
// }

type Dependency struct {
    Dest string
    Src string
    Commitish string
}

func Process(file string) {
    // fmt.Println(file)

    deps := parse(file)

    fetch(file, deps)
}

func fetch(file string, deps []Dependency) {
    // fmt.Println("in fetch()")

    dir := filepath.Dir(file);
    // fmt.Println(dir)

    // mode of dir
    fi, err := os.Lstat(dir)
    if err != nil {
        panic(err)
    }
    mode := fi.Mode()

    for _, d := range deps {
        base := filepath.Base(d.Dest)
        parent := filepath.Dir(d.Dest)

        log.Println("working on", d.Src)

        // fmt.Println(parent, base)

        if _, err := os.Stat(parent); os.IsNotExist(err) {
            err := os.MkdirAll(parent, mode)

            if err != nil {
                panic(err)
            }
        }

        if _, err := os.Stat(parent + "/" + base); os.IsNotExist(err) {
            log.Println("cloning", d.Src)
            // clone repo into base
            cmd := exec.Command("git", "clone", d.Src, dir + "/" + parent + "/" + base, "--depth", "1")
            // fmt.Println(cmd)
            err := cmd.Start()
            if err != nil {
                panic(err)
            }
        } else if _, err := os.Stat(parent + "/" + base + "/.git"); os.IsNotExist(err) {
            panic("'" + parent + "/" + base + "' is not a git repo")
        } else {
            // , parent + "/" + base
            cmd := exec.Command("git", "remote", "-v")
            cmd.Dir = parent + "/" + base
            // fmt.Println(cmd)

            out, err := cmd.Output()
            // fmt.Println("output:", string(out))

            if err != nil {
                panic(err)
            }

            same := strings.Contains(string(out), d.Src)
            // fmt.Println("same?", d.Src, same)

            if !same {
                panic(parent + "/" + base + " is not " + d.Src)
                continue
            }

            log.Println("fetching latest")
            cmd = exec.Command("git", "fetch", "--all")
            cmd.Dir = parent + "/" + base
            // fmt.Println(cmd)
            cmd.Run()
        }

        cmd := exec.Command("git", "checkout", "-q", d.Commitish)
        cmd.Dir = parent + "/" + base
        // fmt.Println(cmd)
        cmd.Run()

        cmd = exec.Command("git", "pull")
        cmd.Dir = parent + "/" + base
        // fmt.Println(cmd)
        cmd.Run()

        log.Println("getting", d.Commitish)
    }
}

func parse(file string) ([]Dependency) {
    jsonBytes := read(file)
    var deps []Dependency

    var f interface{}
    err := json.Unmarshal(jsonBytes, &f)
    if err != nil {
        panic(err)
    }

    m := f.(map[string]interface{})

    reqs := m["require"].(map[string]interface{})

    var str, src, commitish = "", "", ""

    for k, v := range reqs {
        str, src, commitish = v.(string), "", "master"

        if strings.Contains(str, "#") {
            idx := strings.Index(str, "#")
            src, commitish = str[0:idx], str[idx+1:]
        } else {
            src = str
        }

        deps = append(deps, Dependency{k, src, commitish})
    }

    return deps
}

func read(file string) ([]byte) {
    dat, err := ioutil.ReadFile(file)
    check(err)

    return dat
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
