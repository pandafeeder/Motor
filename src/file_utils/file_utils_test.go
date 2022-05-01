package file_utils

import (
        "os"
        "testing"
        "fmt"
        "reflect"
)


func TestReadLinesAndWriteLines(t *testing.T) {
        data := []string{"abc", "ABC", "xyz", "XYZ"}
        tmpfile, err := os.CreateTemp("", "fortesting.log")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        err = WriteLinesToFile(tmpfile.Name(), data)
        defer os.Remove(tmpfile.Name())
        if err != nil {
                t.Error(err)
        }
        lines, err := ReadLinesFromFile(tmpfile.Name())
        if err != nil {
                t.Error(err)
        }
        if !reflect.DeepEqual(data, lines) {
                t.Error("content write to file and read back from that file not queal")
        }
        hash, err := GetFileMd5sum(tmpfile.Name())
        if err != nil {
                t.Error(err)
        }
        if hash == "" {
                t.Error("Failed to get md5sum")
        }
        exists, err := CheckFileExistence(tmpfile.Name())
        if err != nil {
                t.Error(err)
        }
        if exists != true {
                t.Errorf("Existence checking for file %s should be true", tmpfile.Name())
        }
}
