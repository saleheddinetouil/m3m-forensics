package main

import (
	"fmt"
	"github.com/google/gops/agent"
)



// ProcessInfo stores process details
type ProcessInfo struct {
	Pid        int
	Executable string
	Cmdline    []string
}



func GetProcessInfo(pid int) (*ProcessInfo, error) {

	proc, err := agent.NewProcess(pid)
	if err != nil {
		return nil, err
	}



	cmdline, err := proc.Cmdline()
	if err != nil {
		return nil, err
	}



	exe, err := proc.Executable()
	if err != nil {

		return nil, err

	}

	return &ProcessInfo{
		Pid:        pid,
		Executable: exe,
		Cmdline:    cmdline,
	}, nil
}




func printProcessInfo(procInfo *ProcessInfo) {
    fmt.Printf("PID: %d\n", procInfo.Pid)
    fmt.Printf("Executable: %s\n", procInfo.Executable)
    fmt.Printf("Cmdline: %v\n", procInfo.Cmdline)
}
