package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg = new(sync.WaitGroup)
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			var address = fmt.Sprintf("scanme.nmap.org:%d", j)
			var conn, err = net.Dial("tcp", address)

			if err != nil {
				// Port is closed or filtered
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
		wg.Wait()
	}
}
