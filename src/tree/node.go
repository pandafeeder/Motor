/*
A node representing a action script
*/

package tree

import "encoding/json"
import "motor/file_utils"

type Status string

/*
Ready : to be queued 
Queued: waiting for running
MissingDependency : missing dependency
Running : running
FinishWithError : finished with error
FinishWithoutError : finished without error
Skipped : skipped
Unready : not ready, this mainly means dependencies ok, but user stopped or parent node failed
Tracing: dependencies are being generated
*/
const (
	Ready Status = "Ready"
        Queued Status = "Queued"
        MissingDependency Status = "MissingDependency"
        Running Status = "Running"
        FinishWithError Status = "FinishWithError"
        FinishWithoutError Status = "FinishWithoutError"
        Skipped Status = "Skipped"
        Unready Status = "Unready"
        Tracing Status = "Tracing"
)

// use Md5sum to indicate file existence
type File struct {
	Md5sum string `json:"Md5sum"`
	Name  string `json:"Name"`
	Path  string `json:"Path"`
}

func (f File) String() string {
        return f.Name
}

// expose Parents/Children as slice of string whose value is other nodes' Name
// can't serilaze circled json
type Node struct {
	Name       string   `json:"Name"`
	Sourcefile string   `json:"Sourcefile"`
	Inputs     []*File  `json:"Inputs"`
	Outputs    []*File  `json:"Outputs"`
	Parents    []string `json:"Parents"`
	Children   []string `json:"Children"`
        Missing    []string `json:"Missing"`
	parents    []*Node
	children   []*Node
	Status     Status `json:"Status"`
	Level      int  `json:"Level"`
}

func (n *Node) ToJSONBytes() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Node) ToJSONString() (string, error) {
	bytes, err := n.ToJSONBytes()
	return string(bytes), err
}

// parents/children field of returned Node is always nil
func JSONStringToNode(str string) (Node, error) {
	node := Node{}
	err := json.Unmarshal([]byte(str), &node)
	return node, err
}

func (n Node) String() string {
        return n.Name
}

// sololy by inputs and outputs
// not checking parents' status here
// relay such checking to Worker
func (n *Node) UpdateStatus() {
        //println("Updating status for node "+n.Name)
        if n.Status == FinishWithoutError {
                return
        }
        if n.Status == FinishWithError {
                return
        }
        if n.Status == Running {
                return
        }
        if len(n.Missing) > 0 {
                n.Status = MissingDependency
                return
        }
        if len(n.Inputs) == 0 {
                n.Status = Ready
                return
        }
        all_inputs_ok := 1
        for _, in := range n.Inputs {
                exists, _ := file_utils.CheckFileExistence(in.Path)
                if !exists {
                        all_inputs_ok = 0
                        break
                }
        }
        if all_inputs_ok == 1 {
                n.Status = Ready
        } else {
                n.Status = Tracing
        }
}

func (n *Node) SetStatus(stat Status) {
        n.Status = stat
}

