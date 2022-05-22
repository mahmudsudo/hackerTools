package main

import (
	"fmt"
	"net"
	"sync"
)


func main(){
	// _,err:= net.Dial("tcp","scanme.nmap.org:80")
	// if err == nil {
	// 	fmt.Println("connection successful")
	// }
	// for i := 1; i <= 1024; i++ {
	// 	address := fmt.Sprintf("scanme.nmap.org:%d", i)
	// 	fmt.Println(address)
	// 	conn, err := net.Dial("tcp", address)
	// 	if err != nil {
	// 	// port is closed or filtered.
	// 	continue
	// 	}
	// 	conn.Close()
	// 	fmt.Printf("%d open\n", i)
	// 	}
	var wg sync.WaitGroup
	for i := 1024; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
		address := fmt.Sprintf("localhost:%d", j)
		conn, err := net.Dial("tcp", address)
		if err != nil {
		return
		}
		conn.Close()
		fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
	fmt.Println(p)
	wg.Done()
	}
	}