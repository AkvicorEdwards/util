package util

import "fmt"

func ExampleTcpPortIsOpen() {
	ip := "172.16.0.1"
	fmt.Println(TcpPortIsOpen(ip, "22"))
	fmt.Println(TcpPortIsOpen(ip, "80"))
	fmt.Println(TcpPortIsOpen(ip, "443"))
	fmt.Println(TcpPortIsOpen(ip, "444"))

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleTcpPortIsOpenByAddr() {
	fmt.Println(TcpPortIsOpenByAddr("172.16.0.1:22"))
	fmt.Println(TcpPortIsOpenByAddr("172.16.0.1:80"))
	fmt.Println(TcpPortIsOpenByAddr("172.16.0.1:443"))
	fmt.Println(TcpPortIsOpenByAddr("172.16.0.1:444"))

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleGetAvailablePort() {
	p, err := GetAvailablePort()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
