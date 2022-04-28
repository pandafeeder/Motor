/*
testing Node JSON serialization and deserialization
*/

package tree

import (
	"encoding/json"
	"testing"
        "reflect"
)
var (
	inf1       = File{"inf1md5sum", "inputfile", "/fakepath/inf1"}
	outf1      = File{"outf1md5sum", "outputfile", "/fakepath/outf1"}
)
var a_node = Node{
	"node1",
	"/fakepath/node1_sourcefile",
	[]*File{&inf1},
	[]*File{&outf1},
	[]string{},
	[]string{},
	[]*Node{},
	[]*Node{},
	Valid,
        -1,
}

func TestSingleNode2JSON(t *testing.T) {
	jbytes, _ := json.Marshal(&a_node)
	jstr := string(jbytes)
	expected_str := `{"Name":"node1",` +
		`"Sourcefile":"/fakepath/node1_sourcefile",` +
		`"Inputs":[{"Md5sum":"inf1md5sum","Name":"inputfile","Path":"/fakepath/inf1"}],` +
		`"Outputs":[{"Md5sum":"outf1md5sum","Name":"outputfile","Path":"/fakepath/outf1"}],` +
		`"Parents":[],` +
		`"Children":[],` +
		`"Status":0,` +
                `"Level":-1}`
	if jstr != expected_str {
		t.Error("SingleNode JSON serialization get unexpected str")
	}
        jstr2, err := a_node.ToJSONString()
        if err != nil {
                t.Error("Node struct method ToJSONString failed")
        }
        if jstr2 != expected_str {
                t.Error("Node struct method ToJSONString returns unexpedted result")
        }
}

func TestSingleNodeFromJSON(t *testing.T) {
	json_str := `{"Name":"node1",` +
		`"Sourcefile":"/fakepath/node1_sourcefile",` +
		`"Inputs":[{"Md5sum":"inf1md5sum","Name":"inputfile","Path":"/fakepath/inf1"}],` +
		`"Outputs":[{"Md5sum":"outf1md5sum","Name":"outputfile","Path":"/fakepath/outf1"}],` +
		`"Parents":[],` +
		`"Children":[],` +
		`"Status":0,` +
                `"Level":-1}`
	node := Node{}
        err := json.Unmarshal([]byte(json_str), &node)
        if err != nil {
                t.Error("Failed to deserialize json from string")
        }
        node2, err := JSONStringToNode(json_str)
        if err != nil {
                t.Error("Failed to deserialize json from string using JSONStringToNode")
        }
        node2.parents = []*Node{}
        node2.children = []*Node{}
        if ! reflect.DeepEqual(node2, a_node) {
                t.Error("node retured by JSONStringToNode after patching filed value Not euqal ot original node")
        }
}

