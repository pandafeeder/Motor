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
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c1 = Node{
	"c1",
	"/fakepath/c1_sourcefile",
	[]*File{&root_outf1, &root_outf2},
	[]*File{&c1_outf1, &c1_outf2},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c2 = Node{
	"c2",
	"/fakepath/c2_sourcefile",
	[]*File{&root_outf2},
	[]*File{&c2_outf1},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c1_1 = Node{
	"c1_1",
	"/fakepath/c1_1_sourcefile",
	[]*File{&c1_outf1, &c1_outf1},
	[]*File{&c1_1_outf1, &c1_1_outf2},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c1_2 = Node{
	"c1_2",
	"/fakepath/c1_2_sourcefile",
	[]*File{&c1_outf2},
	[]*File{&c1_2_outf1},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c2_1 = Node{
	"c2_1",
	"/fakepath/c2_1_sourcefile",
	[]*File{&c2_outf1, &root_outf1},
	[]*File{&c2_1_outf1},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}
var c1_1_1 = Node{
	"c1_1_1",
	"/fakepath/c2_1_sourcefile",
	[]*File{&c1_2_outf1, &c2_1_outf1},
	[]*File{},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	InValid,
	-1,
}

var nodes = []*Node{&root, &c1, &c2, &c1_1, &c1_2, &c2_1, &c1_1_1}

func TestBuildTree(t *testing.T) {
	nodes = BuildTree(nodes)
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
	if len(root.parents) != 0 {
		t.Error("root's parents is not empty")
	}
	if len(root.children) != 3 {
		t.Errorf("root's children not 3 but %d", len(root.children))
	}
	if c1.parents[0] != &root {
		t.Errorf("c1 parent not root but %v", c1.parents[0])
	}
        if !reflect.DeepEqual(c1_1_1.parents, []*Node{&c1_2, &c2_1}) {
                t.Error("c1_1_1 parents not expected")
                t.Errorf("Expecting: [%v %v]", c1_2, c2_1)
                t.Errorf("Got:       %v", c1_1_1.parents)
        }
        if !reflect.DeepEqual(c2_1.parents, []*Node{&root, &c2}) {
                t.Error("c2_1 parents not expected")
                t.Errorf("Expecting: [%v %v]", root, c2)
                t.Errorf("Got:       %v", c2_1.parents)
        }
	//jbytes, err := json.MarshalIndent(nodes, "", "    ")
	//if err != nil {
	//        fmt.Println(err)
	//} else {
	//        fmt.Println(string(jbytes))
	//}
}
