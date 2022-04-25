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

type File struct {
	Mdsum string `json:"md5sum"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}

type Node struct {
	Name       string  `json:"name"`
	Sourcefile string  `json:"sourcefile"`
	Inputs     []*File `json:"inputs"`
	Outputs    []*File `json:"output"`
	Parents    []*Node `json:"parents"`
	Children   []*Node `json:"children"`
	Status     Status  `json:"status"`
}

func (n *Node) ToJSONBytes() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Node) ToJSONString() (string) {
        bytes, err := n.ToJSONBytes()
        if err != nil {
                panic(err)
        }
        return string(bytes)
}
