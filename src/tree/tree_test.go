/*
testing Tree operation
tested  Tree is like below:
               root
              /     \
           c1       c2
          / \       |
        c1_1 c1_2   c2_1
*/

package tree

import (
        "testing"
)

var (
        root_outf1 = File{"rootOutf1Md5sum", "root_outf1", "/fakepath/root_outf1"}
        root_outf2 = File{"rootOutf1Md5sum", "root_outf2", "/fakepath/root_outf2"}
        c1_outf1   = File{"rootOutf1Md5sum", "c1_outf1", "/fakepath/c1_outf1"}
        c1_outf2   = File{"rootOutf1Md5sum", "c1_outf2", "/fakepath/c1_outf2"}
        c2_outf1   = File{"rootOutf1Md5sum", "c2_outf1", "/fakepath/c2_outf1"}
        c1_1_outf1 = File{"rootOutf1Md5sum", "c1_1_outf1", "/fakepath/c1_1_outf1"}
        c1_1_outf2 = File{"rootOutf1Md5sum", "c1_1_outf2", "/fakepath/c1_1_outf2"}
        c1_2_outf1 = File{"rootOutf1Md5sum", "c1_2_outf1", "/fakepath/c1_2_outf1"}
)
var root = Node {
        "root",
        "/fakepath/root_sourcefile",
        []*File{},
        []*File{&root_outf1, &root_outf2},
        []*Node{},
        []*Node{},
        InValid,
}
var c1 = Node {
        "c1",
        "/fakepath/c1_sourcefile",
        []*File{&root_outf1, &root_outf2},
        []*File{&c1_outf1, &c1_outf2},
        []*Node{},
        []*Node{},
        InValid,
}
var c2 = Node {
        "c2",
        "/fakepath/c2_sourcefile",
        []*File{&root_outf2},
        []*File{&c2_outf1},
        []*Node{},
        []*Node{},
        InValid,
}
var c1_1 = Node {
        "c1_1",
        "/fakepath/c1_1_sourcefile",
        []*File{&c1_outf1, &c1_outf1},
        []*File{&c1_1_outf1, &c1_1_outf2},
        []*Node{},
        []*Node{},
        InValid,
}
var c1_2 = Node {
        "c1_2",
        "/fakepath/c1_2_sourcefile",
        []*File{&c1_outf2},
        []*File{&c1_2_outf1},
        []*Node{},
        []*Node{},
        InValid,
}
var c2_1 = Node {
        "c2_1",
        "/fakepath/c2_1_sourcefile",
        []*File{&c2_outf1},
        []*File{},
        []*Node{},
        []*Node{},
        InValid,
}

func TestBuildTree(t *testing.T) {
        nodes := []*Node{&root, &c1, &c2, &c1_1, &c1_2, &c2_1}
        BuildTree(nodes)
}

func TestUniqSlice(t *testing.T) {
        files := []*File{&root_outf1, &root_outf2, &root_outf1, &root_outf2, &c1_outf1, &c1_outf1, &c1_outf2}
        files = UniqSlice(files)

}
