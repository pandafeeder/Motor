/*
Tree related operations
*/

package tree

import (
        "os"
	"fmt"
	"motor/data_utils"
)

type pair struct {
        parents []*Node
        children []*Node
}

/*
set a node's parents/children according to its inputs/outputs
*/
func BuildTree(nodes []*Node) []*Node {
	edges := make([]*File, 0, 0)
	for _, node := range nodes {
		inputs := node.Inputs
		for _, i := range inputs {
			edges = append(edges, i)
		}
		outputs := node.Outputs
		for _, o := range outputs {
			edges = append(edges, o)
		}
	}
	edges = data_utils.UniqSliceOfPtrByVal(edges)
        mapping := make(map[*File]pair)
	for _, edge := range edges {
		parent_nodes := make([]*Node, 0, 0)
		children_nodes := make([]*Node, 0, 0)
		for _, node := range nodes {
			for _, file := range node.Outputs {
				if *edge == *file {
					parent_nodes = append(parent_nodes, node)
				}
			}
                        for _, file := range node.Inputs {
                                if *edge == *file {
                                        children_nodes = append(children_nodes, node)
                                }
                        }
		}
                if len(parent_nodes) > 1 {
                        fmt.Printf("Error(tree.BuildTree): File %v has multi parents declaring it as output\n", edge)
                        os.Exit(1)
                }
                mapping[edge] = pair{parent_nodes, children_nodes}
	}
        for _, e := range edges {
                apair := mapping[e]
                parents := apair.parents
                children := apair.children
                for _, p := range parents {
                        for _, c := range children {
                                if data_utils.IndexOf(p.Children, c.Name) == -1 {
                                        p.Children = append(p.Children, c.Name)
                                }
                                if data_utils.IndexOf(p.children, c) == -1 {
                                        p.children = append(p.children, c)
                                }
                                if data_utils.IndexOf(c.Parents, p.Name) == -1 {
                                        c.Parents = append(c.Parents, p.Name)
                                }
                                if data_utils.IndexOf(c.parents, p) == -1 {
                                        c.parents = append(c.parents, p)
                                }
                        }
                }
        }
        nodes = AnnotaeNodeLevel(nodes)
        return nodes
}

func AnnotaeNodeLevel(nodes []*Node) []*Node {
        for _, node := range nodes {
                depth := 0
                // think of this as a simple queue
                parents := node.parents
                for len(parents) > 0 {
                        for _, p := range parents {
                                parents = parents[1:]
                                if len(p.parents) > 0 {
                                        for _, upper_p := range p.parents {
                                                parents = append(parents, upper_p)
                                        }
                                }
                        }
                        depth += 1
                }
                node.Level = depth
        }
       return nodes
}

//func UpdateDependency(nodes []*Node) {
//}
//
//func FindNodeByName(nodes []*Node) *Node {
//}
//
//func FindNodeByLevel(nodes []*Node) []*Node {
//}
//func SanityCheckOnTree() error {
//}

