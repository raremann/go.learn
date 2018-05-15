package github

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
)

func getTempFile() string {
	d := "/tmp/.gohack."
	p := os.Getpid()
	return d + strconv.Itoa(p)
}

func ForkWait() {
	e := os.Getenv("GOHACKEDITOR")
	if e == "" {
		e = "emacs"
	}
	f := getTempFile()
	log.Printf("forking %s", e)
	pid, err := syscall.ForkExec(e, []string{e, f}, &syscall.ProcAttr{Files: []uintptr{0, 1, 2}})
	if err != nil {
		panic(err.Error())
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		panic(err.Error())
	}
	state, err := proc.Wait()
	if err != nil {
		panic(err.Error())
	}
	println("string:", state.String())
	println("pid:", state.Pid())
	println("Parent exiting")
}

func readTempFile(f string) string {
	// We expect the file to be relatively small. Use ioutil.ReadFile
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Could not read file %s: %s", f, err)
	}
	return string(data)
}

func GetUpdateContents() string {
	ForkWait()
	f := getTempFile()
	contents := readTempFile(f)
	log.Printf("file %s: contents %s", f, contents)
	return contents
}
