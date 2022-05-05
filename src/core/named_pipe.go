package core

import (
        "bufio"
        "runtime"
        "os"
        "sync"
        "syscall"
        "motor/file_utils"
)

// make the named pipe singleton
var single_pipe *MotorPipe
var once sync.Once
var lock sync.Mutex

type MotorPipe struct {
        File string
}


func GetPipe() *MotorPipe {
        once.Do(func () {
                file_sig := file_utils.GenFileSig()
                file := ".motor_ipc_pipe_"+file_sig+".ipc"
                if err := syscall.Mkfifo(file, 0666); err != nil {
                        panic(err)
                }
                single_pipe = &MotorPipe{
                        File: file,
                }
        })
        return single_pipe
}


func (m *MotorPipe) Read() (msg string) {
        if _, err := os.Stat(m.File); os.IsNotExist(err) {
                panic(err)
        }
        fh, _ := os.OpenFile(m.File, os.O_RDWR, os.ModeNamedPipe)
        defer fh.Close()
        for {
                reader := bufio.NewReader(fh)
                line, _, err := reader.ReadLine()
                if err == nil {
                        msg = string(line)
                }
                break
        }
        return
}

func (m *MotorPipe) Write(cmd string) {
        lock.Lock()
        if _, err := os.Stat(m.File); os.IsNotExist(err) {
                panic(err)
        }
        fh, err := os.OpenFile(m.File, os.O_WRONLY, 0777)
        defer fh.Close()
        defer lock.Unlock()
        if err != nil {
                panic(err)
        }
        fh.WriteString(cmd+"\n")
        runtime.Goexit()
}


