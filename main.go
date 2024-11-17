package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/gops/agent"
)

func main() {
	if len(os.Args) > 1 {
		pid, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Invalid PID:", err)
		}
		analyzeProcess(pid)
	} else {
		listProcesses()
	}
}


func analyzeProcess(pid int) {
    fmt.Println("Analyzing process", pid)
    procInfo, err := GetProcessInfo(pid)
    if err != nil {
        log.Fatal("Error getting process information:", err)
    }
    fmt.Println("Process Information:")
    printProcessInfo(procInfo)


    netConns, err := GetNetworkConnections(pid)
    if err != nil {

        log.Fatal("Error getting network connections:", err)

    }

    fmt.Println("\nNetwork Connections:")
    printNetworkConnections(netConns)



}


func listProcesses() {


    procs, err := agent.Processes()
    if err != nil {
        log.Fatal("Error listing processes:", err)

    }

    fmt.Println("Running Processes:")

    for _, proc := range procs {


        procInfo, err := GetProcessInfo(proc.Pid)

        if err != nil { // handle error

            fmt.Printf("Error getting process information for PID %d: %v\n", proc.Pid, err)

            continue
        }
        printProcessInfo(procInfo)

        fmt.Println("-------------") // Separator between processes

    }


}
