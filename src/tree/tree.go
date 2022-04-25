/*
Tree related operations
*/

package tree

import (
        "fmt"
)

func BuildTree(nodes []*Node) {
        edges := make([]*File, 0, 0)
        for _, node := range nodes {
                inputs  := node.Inputs
                for _, i := range inputs {
                        edges = append(edges, i)
                }
                outputs := node.Outputs
                for _, o := range outputs {
                        edges = append(edges, o)
                }
        }
        edges = UniqSlice(edges)
        for _, edge := range edges {
                fmt.Println(edge)
        }
}

func UpdateDependency(nodes *[]Node) {

}
