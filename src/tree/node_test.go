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
	[]*Node{},
	[]*Node{},
	Valid,
}

func TestSingleNode2JSON(t *testing.T) {
	jbytes, _ := json.Marshal(&a_node)
	jstr := string(jbytes)
	expected_str := `{"name":"node1",` +
		`"sourcefile":"/fakepath/node1_sourcefile",` +
		`"inputs":[{"md5sum":"inf1md5sum","name":"inputfile","path":"/fakepath/inf1"}],` +
		`"output":[{"md5sum":"outf1md5sum","name":"outputfile","path":"/fakepath/outf1"}],` +
		`"parents":[],` +
		`"children":[],` +
		`"status":0}`
	if jstr != expected_str {
		t.Error("SingleNode JSON serialization get unexpected str")
	}
}

func TestSingleNodeFromJSON(t *testing.T) {
	json_str := `{"name":"node1",` +
		`"sourcefile":"/fakepath/node1_sourcefile",` +
		`"inputs":[{"md5sum":"inf1md5sum","name":"inputfile","path":"/fakepath/inf1"}],` +
		`"output":[{"md5sum":"outf1md5sum","name":"outputfile","path":"/fakepath/outf1"}],` +
		`"parents":[],` +
		`"children":[],` +
		`"status":0}`
	node := Node{}
	json.Unmarshal([]byte(json_str), &node)
	if !reflect.DeepEqual(node, a_node) {
                t.Error("SingleNode deserialization get unexpected Node struct")
        }
}

