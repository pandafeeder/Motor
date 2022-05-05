/*
*/

package main

import (
        "motor/core"
	"flag"
)

func main () {
	flag.Parse()
	root_dir := flag.Arg(0)
        graph_nodes, err := core.Build(root_dir)
        if err != nil {
                panic(err)
        }
        core.MainLoop(graph_nodes)
}
