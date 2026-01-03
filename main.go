package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

func Workers(Host string, ports chan int, results chan int) {
	for p := range ports {
		// adress := fmt.Sprintf("scanme.nmap.org:%d", p)

		adress := fmt.Sprintf("%s:%d", Host, p)

		conn, err := net.Dial("tcp", adress)

		if err != nil {
			// there is errror port is closed
			results <- 0
			// fmt.Printf("port %d is closed\n", p)
		} else {
			results <- p
			fmt.Printf("PORT %d is open\n", p)
			_ = conn.Close()

		}
	}
}

func main() {
	HostPtr := flag.String("Host", "scanme.nmap.org", "Website to Scan ")

	flag.Parse()

	portch := make(chan int, 100)
	resultch := make(chan int)

	var openports []int

	for i := 1; i <= 100; i++ {
		go Workers(*HostPtr, portch, resultch)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			portch <- i
		}
		close(portch)
	}()

	for i := 0; i < 1024; i++ {
		p := <-resultch

		if p != 0 {
			openports = append(openports, p)
		}
	}

	close(resultch)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
