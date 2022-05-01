package file_utils

import (
        "errors"
        "io"
        "os"
        "strings"
        "crypto/md5"
        "fmt"
)

func ReadLinesFromFile(file string) ([]string, error) {
        if _, err := os.Stat(file); os.IsNotExist(err) {
                return nil, err
        }
        data, err := os.ReadFile(file)
        if err != nil {
                return nil, err
        }
        return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n"), nil
}


func WriteLinesToFile(file string, data []string) error {
        data_bytes := []byte(strings.Join(data[:], "\n"))
        err := os.WriteFile(file, data_bytes, 0755)
        return err
}

func GetFileMd5sum(file string) (string, error) {
        fh, err := os.Open(file)
        if err != nil {
                return "", err
        }
        defer fh.Close()
        hash := md5.New()
        _, err = io.Copy(hash, fh)
        if err != nil {
                return "", err
        }
        return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func CheckFileExistence(file string) (bool, error) {
        if _, err := os.Stat(file); err == nil {
                return true, nil
        } else if errors.Is(err, os.ErrNotExist) {
                return false, nil
        } else {
                return false, err
        }
}
