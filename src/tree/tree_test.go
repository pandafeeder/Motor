/*
testing Tree operation
tested  Tree is like below:
                root
              /   \   \
           c1      \   c2
          / \       \  /
        c1_1 c1_2   c2_1
              \     /
              c1_1_1
*/

package tree

import (
	"reflect"
	"testing"
)

//import "fmt"
//import "encoding/json"

var (
	root_outf1 = File{"rootOutf1Md5sum", "root_outf1", "/fakepath/root_outf1"}
	root_outf2 = File{"rootOutf2Md5sum", "root_outf2", "/fakepath/root_outf2"}
	c1_outf1   = File{"c1_outf1Md5sum", "c1_outf1", "/fakepath/c1_outf1"}
	c1_outf2   = File{"c1_outf2Md5sum", "c1_outf2", "/fakepath/c1_outf2"}
	c2_outf1   = File{"c2_outf1Md5sum", "c2_outf1", "/fakepath/c2_outf1"}
	c1_1_outf1 = File{"c1_1_outf1Md5sum", "c1_1_outf1", "/fakepath/c1_1_outf1"}
	c1_1_outf2 = File{"c1_1_outf2Md5sum", "c1_1_outf2", "/fakepath/c1_1_outf2"}
	c1_2_outf1 = File{"c1_2_outf1Md5sum", "c1_2_outf1", "/fakepath/c1_2_outf1"}
	c2_1_outf1 = File{"c2_1_outf1Md5sum", "c2_1_outf1", "/fakepath/c2_1_outf1"}
)
var root = Node{
	"root",
	"/fakepath/root_sourcefile",
	[]*File{},
	[]*File{&root_outf1, &root_outf2},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c1 = Node{
	"c1",
	"/fakepath/c1_sourcefile",
	[]*File{&root_outf1, &root_outf2},
	[]*File{&c1_outf1, &c1_outf2},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c2 = Node{
	"c2",
	"/fakepath/c2_sourcefile",
	[]*File{&root_outf2},
	[]*File{&c2_outf1},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c1_1 = Node{
	"c1_1",
	"/fakepath/c1_1_sourcefile",
	[]*File{&c1_outf1, &c1_outf1},
	[]*File{&c1_1_outf1, &c1_1_outf2},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c1_2 = Node{
	"c1_2",
	"/fakepath/c1_2_sourcefile",
	[]*File{&c1_outf2},
	[]*File{&c1_2_outf1},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c2_1 = Node{
	"c2_1",
	"/fakepath/c2_1_sourcefile",
	[]*File{&c2_outf1, &root_outf1},
	[]*File{&c2_1_outf1},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}
var c1_1_1 = Node{
	"c1_1_1",
	"/fakepath/c2_1_sourcefile",
	[]*File{&c1_2_outf1, &c2_1_outf1},
	[]*File{},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}

var input_nodes = []Node{root, c1, c2, c1_1, c1_2, c2_1, c1_1_1}

func TestBuildTree(t *testing.T) {
        nodes, err := BuildTree(input_nodes)
        if err != nil {
                t.Errorf("Got error as '%s' which shouldn't", err)
        }
	expected_level := []int{0, 1, 1, 2, 2, 2, 3}
	returned_level := make([]int, 0)
	for _, node := range nodes {
		returned_level = append(returned_level, node.Level)
	}
	if !reflect.DeepEqual(expected_level, returned_level) {
		t.Error("Nodes Level not as expected")
		t.Errorf("Expecting: %v\n", expected_level)
		t.Errorf("Got:       %v\n", returned_level)
	}

        root_copy := nodes[0]
        c1_copy := nodes[1]
        c2_copy := nodes[2]
        c1_2_copy := nodes[4]
        c2_1_copy := nodes[5]
        c1_1_1_copy := nodes[6]

	if len(root_copy.parents) != 0 {
		t.Error("root's parents is not empty")
	}
	if len(root_copy.children) != 3 {
		t.Errorf("root's children not 3 but %d", len(root.children))
	}
	if c1_copy.parents[0] != root_copy {
		t.Errorf("c1 parent not root but %v", c1_copy.parents[0])
	}
        if !reflect.DeepEqual(c1_1_1_copy.parents, []*Node{c1_2_copy, c2_1_copy}) {
                t.Error("c1_1_1 parents not expected")
                t.Errorf("Expecting: [%v %v]", c1_2_copy, c2_1_copy)
                t.Errorf("Got:       %v", c1_1_1_copy.parents)
        }
        if !reflect.DeepEqual(c2_1_copy.parents, []*Node{root_copy, c2_copy}) {
                t.Error("c2_1 parents not expected")
                t.Errorf("Expecting: [%v %v]", root_copy, c2_copy)
                t.Errorf("Got:       %v", c2_1_copy.parents)
        }
	//jbytes, err := json.MarshalIndent(nodes, "", "    ")
	//if err != nil {
	//        fmt.Println(err)
	//} else {
	//        fmt.Println(string(jbytes))
	//}
}

var bad_node = Node{
        "bad_node",
	"/fakepath/bad_node",
	[]*File{&c1_outf1},
	[]*File{&root_outf2},
	[]string{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Unready,
	-1,
}

var input_circular_nodes = []Node{root, c1, c2, bad_node, c1_1, c1_2, c2_1, c1_1_1}
// expecting error as File has multi parents
func TestBuildTreeForCircular(t *testing.T) {
        _, err := BuildTree(input_circular_nodes)
        if err == nil {
                t.Error("Expecting error as `File root_outf2 has multi parents declaring it as output`")
        }
}

// expecting error as Circular dependency found
func TestCircular(t *testing.T) {
        nodes, err := BuildTree(input_nodes)
        if err != nil {
                t.Errorf("Got error as '%s' which shouldn't", err)
        }
        c1_copy := nodes[1]
        c1_1_copy := nodes[3]
        c1_copy.parents = append(c1_copy.parents, c1_1_copy)
        err = CheckCircularDependency(nodes)
        if err == nil {
                t.Error("Expecting error as `Circular dependency found for Node c1")
        }
}

