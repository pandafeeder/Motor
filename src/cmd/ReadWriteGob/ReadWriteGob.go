package main

import . "motor/tree"
import (
        "os"
        "encoding/gob"
)


//func writeGob(file string, obj interface{}) error {
//        fh, err := os.Create(file)
//        if err == nil {
//                encoder := gob.NewEncoder(fh)
//                encoder.Encode(obj)
//        }
//        fh.Close()
//        return err
//}
//
//func readGob(file string, obj interface{}) error {
//        fh, err := os.Open(file)
//        if err == nil {
//                decoder := gob.NewDecoder(fh)
//                err = decoder.Decode(obj)
//        }
//        fh.Close()
//        return err
//}

