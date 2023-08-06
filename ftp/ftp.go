package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "114.115.213.117:21")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dial)
}
