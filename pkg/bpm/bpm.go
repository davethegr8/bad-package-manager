package bpm

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
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
    fmt.Println(file)
    // basedir of file

    deps := parse(file)
    fmt.Println(deps)
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
