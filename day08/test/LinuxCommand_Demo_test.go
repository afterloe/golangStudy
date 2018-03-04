package test

import (
	"testing"
	"os/exec"
	"log"
	"strings"
)

func Test_Demo_demo(t *testing.T) {
	//errorNotFound := errors.New("executable file not found in $PATH")
	arg := "/Users/afterloe"
	cmd := exec.Command("ls", arg)
	log.Println("Running command and waiting for it finish")
	tpl, err := cmd.Output()
	if nil != err {
		t.Error("exec command ls " + arg + " failed.")
		return
	}
	str := string(tpl)
	files := strings.Split(str, "\n")
	log.Println(files)
	log.Println(len(files))
	for _, file := range files {
		log.Printf("item is -> %s", file)
	}
}
