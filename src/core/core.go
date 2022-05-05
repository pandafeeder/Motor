package core

import . "motor/tree"


func MainLoop(graph_nodes []*Node) {
        pipe := GetPipe()
        scheduler := Scheduler{Nodes: graph_nodes, Pipe: pipe}
        scheduler.Run()
}
