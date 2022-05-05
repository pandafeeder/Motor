package core

import . "motor/tree"
import (
        "os"
        "encoding/gob"
        "motor/file_utils"
)

func Writegob(file string, obj interface{}) error {
        fh, err := os.Create(file)
        defer fh.Close()
        if err == nil {
                encoder := gob.NewEncoder(fh)
                encoder.Encode(obj)
        }
        return err
}

func ReadGob(file string, obj interface{}) error {
        fh, err := os.Open(file)
        defer fh.Close()
        if err == nil {
                decoder := gob.NewDecoder(fh)
                err = decoder.Decode(obj)
        }
        return err
}

func WriteSnapshot(nodes []*Node) {
        file_sig := file_utils.GenFileSig()
        filename := ".motor_"+file_sig+".gob"
        Writegob(filename, nodes)
}

func ReadSnapshot(file string) {
}
