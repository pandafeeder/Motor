/*
A node representing a action script
*/

package tree

import "encoding/json"

type Status uint8

const (
	Valid Status = iota
	InValid
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
