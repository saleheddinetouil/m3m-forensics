package main

import (
    "fmt"

	"github.com/google/gops/agent"
)


// NetConn represents a network connection.
type NetConn struct {

	LocalAddr  string
	RemoteAddr string
}



// GetNetworkConnections retrieves network connection information for a process.
func GetNetworkConnections(pid int) ([]NetConn, error) {
	proc, err := agent.NewProcess(pid)
	if err != nil {
		return nil, err
	}



	conns, err := proc.Connections()
	if err != nil {
		return nil, err
	}


	var netConns []NetConn
	for _, c := range conns {
		netConns = append(netConns, NetConn{LocalAddr: c.LocalAddr().String(), RemoteAddr: c.RemoteAddr().String()})
	}



	return netConns, nil


}





func printNetworkConnections(connections []NetConn) {

    if len(connections) == 0{
        fmt.Println("No Network connections found")
        return
    }
	for _, conn := range connections {
		fmt.Printf("Local Address: %s, Remote Address: %s\n", conn.LocalAddr, conn.RemoteAddr)

	}
}
