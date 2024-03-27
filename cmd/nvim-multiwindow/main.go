package main

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
)

func main() {
	pid := os.Getpid()
	process, err := ps.FindProcess(pid)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Process is %s\n", process.Executable())
	fmt.Printf("Process ID %d\n", process.Pid())

	parentProcess, err := ps.FindProcess(process.PPid())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parent Process is %s\n", parentProcess.Executable())
	fmt.Printf("Parent Process ID %d\n", parentProcess.Pid())

}
