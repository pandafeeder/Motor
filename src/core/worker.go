package core

import . "motor/tree"
import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

type Worker struct {
	node        *Node
	cmd         *exec.Cmd
	script_chan chan error
}

func (w *Worker) MainLoop(inst_chan chan Instruct, res_chan chan Respond) {
	for {
                println(w.node.Name+" for loop")
		select {
		case msg := <-inst_chan:
                        println(w.node.Name+" getting from inst_chan")
			switch msg {
			case Run:
                                if w.node.Status == Ready {
                                        w.node.Status = Running
				        go w.Run(res_chan)
                                }
			case Invalid:
				res_chan <- Respond("Invalid " + w.node.Name)
			case Kill:
				res_chan <- Respond("Killing " + w.node.Name)
			case Skip:
				res_chan <- Respond("Skipping " + w.node.Name)
			case ForceValid:
				res_chan <- Respond("ForceValid " + w.node.Name)
			default:
				res_chan <- Respond("Unsupported " + msg)
			}
		case <-time.After(time.Second * 10):
                        println(w.node.Name+" getting from timeout")
			select {
			case err := <-w.script_chan:
                                println(w.node.Name+" getting from script_chan")
				if err == nil {
                                        w.node.Status = FinishWithoutError
					res_chan <- Respond(w.node.Name + " finished without error")
				} else {
                                        w.node.Status = FinishWithError
					res_chan <- Respond(fmt.Sprintf("%v", err))
				}
		        //case <-time.After(time.Second * 5):
			default:
                                println(w.node.Name+" getting from default")
				w.node.UpdateStatus()
                                if w.node.Status == Ready {
                                        println(w.node.Name + " is Ready")
                                        inst_chan <- Run
                                }
			}
		}
	}
}

func (w *Worker) Run(res_chan chan Respond) {
		res_chan <- Respond("Running " + w.node.Name)
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
			log.Printf("%s finished with error: %v", w.node.Name, err)
		}
		w.script_chan <- err
}
