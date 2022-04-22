/*

 */

package node

import "encoding/json"

type Node struct {
	Sourcefile string   `json:"sourcefile"`
	Inputs     []string `json:"inputs"`
	Outputs    []string `json:"output"`
	Parents    []string `json:"parents"`
	Children   []string `json:"children"`
	Level      uint     `json:"level"`
}

func (n *Node) ToJson() ([]byte, error) {
	return json.Marshal(n)
}

func BuildGraph(nodes ...Node) {
}

func UpdateDependency(nodes *[]Node) {
}
