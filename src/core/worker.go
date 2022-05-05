package core

import . "motor/tree"
import (
        "os/exec"
        "fmt"
        "log"
        "time"
        "runtime"
)

type Worker struct {
        node *Node
        cmd  *exec.Cmd
        script_chan chan error
}

func (w *Worker) MainLoop(inst_chan chan Instruct, res_chan chan Respond) {
        for {
                select {
                case msg := <-inst_chan:
                                switch msg {
                                case Run:
                                        go w.Run(res_chan)
                                case Invalid:
                                        res_chan <- Respond("Invalid "+w.node.Name)
                                case Kill:
                                        res_chan <- Respond("Killing "+w.node.Name)
                                case Skip:
                                        res_chan <- Respond("Skipping "+w.node.Name)
                                case ForceValid:
                                        res_chan <- Respond("ForceValid "+w.node.Name)
                                default:
                                        res_chan <- Respond("Unsupported "+msg)
                                }
                case <- time.After(time.Second*10):
                        select {
                        case err := <-w.script_chan:
                                if err == nil {
                                        res_chan <- Respond(w.node.Name+" finished without error")
                                } else {
                                        res_chan <- Respond(fmt.Sprintf("%v", err))
                                }
                        }
                }
        }
}

func (w *Worker) Run(res_chan chan Respond) {
        if w.node.Status == Ready {
                res_chan <- Respond("Running "+w.node.Name)
                script := w.node.Sourcefile
                cmd := exec.Command(script)
                w.cmd = cmd
                err := cmd.Start()
                if err != nil {
                        w.script_chan <- err
                        runtime.Goexit()
                }
                err = cmd.Wait()
                if err != nil {
                        log.Printf("%s finished with error: %v", err)
                }
                w.script_chan <- err
        }
}


