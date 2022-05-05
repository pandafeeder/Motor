/*
Print out JSON representation of the Tree
*/

package main

import (
        "motor/core"
	"flag"
	"fmt"
        "encoding/json"
)

func main () {
	flag.Parse()
	root_dir := flag.Arg(0)
        graph_nodes, err := core.Build(root_dir)
        if err != nil {
                panic(err)
        }
        jbytes, err := json.MarshalIndent(graph_nodes, "", "    ")
        if err != nil {
                panic(err)
        }
        fmt.Println(string(jbytes))
}


