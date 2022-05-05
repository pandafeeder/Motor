package core

import . "motor/tree"

import "encoding/json"
import "fmt"
import (
	"log"
	"sync"
)

type Scheduler struct {
	Nodes         []*Node
	Pipe          *MotorPipe
	InstructChans map[string]chan Instruct
	RespondChan   chan Respond
}

func (s *Scheduler) Run() {
	s.MakeChanForEachNode()
	s.Init()

	// TODO: writeout snapshot here
	jbytes, err := json.MarshalIndent(s.Nodes, "", "    ")
	if err != nil {
	        panic(err)
	}
	fmt.Println(string(jbytes))

	var wg sync.WaitGroup
	for _, node := range s.Nodes {
		wg.Add(1)
		inst_chan := s.InstructChans[node.Name]
		go func(n *Node, inst_c chan Instruct, res_c chan Respond) {
			defer wg.Done()
			w := Worker{node: n, script_chan: make(chan error)}
			w.MainLoop(inst_c, res_c)
		}(node, inst_chan, s.RespondChan)
	}

	// init run
	for _, v := range s.InstructChans {
		v <- Run
	}

	go func() {
		//for msg := range s.RespondChan {
		//	log.Println(msg)
		//}
                for {
                        select {
                        case res := <-s.RespondChan:
                                log.Println(res)
                        }
                }
	}()

	for {
		msg := s.Pipe.Read()
		println(msg)
	}

	wg.Wait()
}

//init each node's status
func (s *Scheduler) Init() {
	var wg sync.WaitGroup
	for _, node := range s.Nodes {
		wg.Add(1)
		go func(n *Node) {
			defer wg.Done()
			n.UpdateStatus()
		}(node)
	}
	wg.Wait()
}

func (s *Scheduler) Handle(p chan Instruct) {
}

func (s *Scheduler) MakeChanForEachNode() {
	var inst_channels = make(map[string]chan Instruct)
	var res_channels = make(chan Respond)//, len(s.Nodes))
	for _, node := range s.Nodes {
		inst_channels[node.Name] = make(chan Instruct)
	}
	s.InstructChans = inst_channels
	s.RespondChan = res_channels
}
