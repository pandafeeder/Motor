package main

import (
        "path/filepath"
        "fmt"
        "strings"
)

func main() {
        pipes, err := filepath.Glob(".motor_ipc_pipe*.ipc")
        if err != nil {
                panic(err)
        }
        if pipes == nil {
                fmt.Println("Info: no pipe found, quit")
        }
        time_stamps := make([]string, 0)
        for _, f := range pipes {
                strings.Split(f, "_")[
        }
}
