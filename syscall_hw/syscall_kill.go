package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// THIS IS THE IN-CLASS ASSIGNMENT
func main() {
	pid := os.Getpid()
	fmt.Println("Current process PID:", pid)
	// replace signal that you kill it with

	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		fmt.Println("Error:", err)
	}
	var exitcode uint32 = 0
	err = syscall.TerminateProcess(handle, exitcode) // send exitcode of zero
	fmt.Println("Process did not terminate")
	// err is an error type object, we hope it is like '0'
	// error object is different from the TerminateProcess
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Process got to end")
	// i dont think that the process got terminated
}

// HW make github pull request
// hw, go through the tutorial linux journal article, ptrace. using ptrace inside a go program

// exit code with zero is a success. give exit code that it will terminate process. tell what the exit code is going to be
