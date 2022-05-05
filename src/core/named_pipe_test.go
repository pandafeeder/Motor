package core

import (
	"os"
	"testing"
)

func TestGetPipe(t *testing.T) {
	p1 := GetPipe()
	p2 := GetPipe()
	if p1.File != p2.File {
		t.Error("Expected single pipe obj")
	}

	cmd2sent := []string{"cmd1", "cmd2", "cmd3", "cmd4", "cmd5"}
	for _, cmd := range cmd2sent {
		go p1.Write(cmd)
		msg := p1.Read()
		if msg != cmd {
			t.Error("Msg readback not equal to sent")
		}
	}
	defer os.Remove(p1.File)
}
