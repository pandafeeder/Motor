/*
Scans a root dir where is the root path contains all action scripts
and builds a tree graph json file
action scripts needs to have below commands to declare dependencies
MRequireInput inputfile
MGenOutput outputfile

plan:
        supports multi arguments for MRequireInput/MGenOutput ?
*/

package main

import . "motor/tree"
import (
	"bufio"
        "errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//throw error if $2 contains anything other than space
var inputPat = regexp.MustCompile(`^\s*MRequireInput\s+(\S+)(.*)$`)
var outputPat = regexp.MustCompile(`^\s*MGenOutput\s(\S+)(.*)$`)
var unemptyPat = regexp.MustCompile(`\S`)

func main() {
	flag.Parse()
	root_dir := flag.Arg(0)
	if _, err := os.Stat(root_dir); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
	files := scanDir(root_dir)

	var wg sync.WaitGroup
        var nodes = make([]Node, len(files))
        for i, f := range files {
                wg.Add(1)
                go func(i int, f string) {
                        defer wg.Done()
                        nodes[i] = makeNodeFromFile(f)
                }(i, f)
        }
        wg.Wait()

        UpdateDependency(&nodes)
}

func scanDir(dir string) (files []string) {
	err := filepath.WalkDir(dir, func(path string, di fs.DirEntry, err error) error {
		// ignore files whose names starts with '.'
		if !di.IsDir() && !strings.HasPrefix(di.Name(), ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

func makeNodeFromFile(file string) (node Node) {
	fmt.Printf("INFO: Parsing file %s\n", file)
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	inputs := make([]string, 0)
	outputs := make([]string, 0)
	reader := bufio.NewReader(fh)
	for {
		line_bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(line_bytes)
		if i_match := inputPat.FindStringSubmatch(line); i_match != nil {
			if unemptyPat.MatchString(i_match[2]) {
				fmt.Printf("ERROR: file %s syntax error as multi arguments for MRequireInput\n", file)
                                panic(errors.New(""))
			}
			inputs = append(inputs, i_match[1])
		}
		if o_match := outputPat.FindStringSubmatch(line); o_match != nil {
			if unemptyPat.MatchString(o_match[2]) {
				fmt.Printf("ERROR: file %s syntax error as multi arguments for MGenOutput\n", file)
                                panic(errors.New(""))
			}
			outputs = append(outputs, o_match[1])
		}
	}
	node = Node{
		Sourcefile: file,
		Inputs:     inputs,
		Outputs:    outputs,
		Parents:    []string{},
		Children:   []string{},
		Level:      0,
	}
        return
}
