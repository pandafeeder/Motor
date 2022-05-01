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
        "motor/file_utils"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)
import "encoding/json"

//throw error if $2 contains anything other than space
var inputPat = regexp.MustCompile(`^\s*MRequireInput\s+(\S+)(.*)$`)
var outputPat = regexp.MustCompile(`^\s*MGenOutput\s(\S+)(.*)$`)
var unemptyPat = regexp.MustCompile(`\S`)
var pat_mapping = map[string]*regexp.Regexp{"input": inputPat, "output": outputPat}

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
                        node, err := makeNodeFromFile(f)
                        if err != nil {
                                panic(err)
                        }
                        nodes[i] = node
		}(i, f)
	}
	wg.Wait()
        graph_nodes, err := BuildTree(nodes)
        if err != nil {
                panic(err)
        }
        jbytes, _ := json.MarshalIndent(graph_nodes, "", "    ")
        fmt.Println(string(jbytes))
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

func makeNodeFromFile(file string) (Node, error) {
	fmt.Printf("Info(makeNodeFromFile): Parsing file %s\n", file)
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	inputs := make([]*File, 0)
	outputs := make([]*File, 0)
        lines, err := file_utils.ReadLinesFromFile(file)
        if err != nil {
                panic(err)
        }
        for _, line := range lines {
                file_obj, kind, err := GrabFileFromLine(line)
                if err != nil {
                        panic(err)
                }
                if kind == "input" && file_obj.Name != "" {
                        inputs = append(inputs, file_obj)
                }
                if kind == "output" && file_obj.Name != "" {
                        outputs = append(outputs, file_obj)
                }
	}
	file_name := strings.Split(filepath.Base(file), ".")[0]
        node := Node{
		Name:       file_name,
		Sourcefile: file,
		Inputs:     inputs,
		Outputs:    outputs,
		Parents:    []string{},
		Children:   []string{},
		Status:     InValid,
		Level:      -1,
	}
	return node, nil
}

func GrabFileFromLine(line string) (*File, string, error) {
        file_obj := File{}
        file_type := ""
        for kind, pat := range pat_mapping {
                file_type = kind
                if match := pat.FindStringSubmatch(line); match != nil {
                        if unemptyPat.MatchString(match[2]) {
			        err_msg := fmt.Sprintf("Error(makeNodeFromFile): `%s` syntax error as multi arguments for MRequireInput|MGenOutput\n", line)
                                return &File{}, "", errors.New(err_msg)
                        }
                        got_file := match[1]
                        if exists, _ := file_utils.CheckFileExistence(got_file); exists == true {
                                md5sum, err := file_utils.GetFileMd5sum(got_file)
                                if err != nil {
                                        panic(err)
                                }
                                file_obj.Md5sum = md5sum
                        }
                        file_obj.Name = filepath.Base(got_file)
                        file_obj.Path = got_file
                        break
                }
        }
        return &file_obj, file_type, nil
}
