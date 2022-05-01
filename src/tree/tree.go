/*
Tree related operations
*/

package tree

import (
	"fmt"
	"motor/data_utils"
        "errors"
)

type pair struct {
        parents []*Node
        children []*Node
}

/*
BuildTree copied each node's value 
and operated on copied node's ptr afterwards
AKA original Node object untouched
*/
func BuildTree(input_nodes []Node) ([]*Node, error) {
        nodes := make([]*Node, 0)
        for id, _ := range input_nodes {
                nodes = append(nodes, &input_nodes[id])
        }
	edges := make([]*File, 0)
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
		parent_nodes := make([]*Node, 0)
		children_nodes := make([]*Node, 0)
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
                        return nil, errors.New(fmt.Sprintf("Error(tree.BuildTree): File %s has multi parents declaring it as output\n", edge))
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
        has_err := CheckCircularDependency(nodes)
        if has_err != nil {
                return nil, errors.New(fmt.Sprintf("Error(tree.BuildTree): %s",has_err))
        }
        return nodes, nil
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

func CheckCircularDependency(nodes []*Node) error {
        for _, node := range nodes {
                parents := node.parents
                children := node.children
                if len(parents) == 0 || len(children) == 0 {
                        continue
                }
                for _, p := range parents {
                        for _, c := range children {
                                if p == c {
                                        return errors.New("Circular dependency found for Node "+node.Name)
                                }
                        }
                }

        }
        return nil
}


//func UpdateDependency(nodes []*Node) {
//}
//
//func FindNodeByName(nodes []*Node) *Node {
//}
//
//func FindNodeByLevel(nodes []*Node) []*Node {
//}

