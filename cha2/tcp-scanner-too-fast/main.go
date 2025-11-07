package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			var address = fmt.Sprintf("scanme.nmap.org:%d", j)
			var conn, err = net.Dial("tcp", address)

			if err != nil {
				// Port is closed or filtered
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
