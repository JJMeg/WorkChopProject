package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	localIPV4 = "127.0.0.1"
	localIPV6 = "::1"
)

func main() {
	var ret []string

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("error: ", err.Error())
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipStr := ip.String()

			if isIPv4(ipStr) && !strings.Contains(ipStr, localIPV4) {
				ret = append(ret, ipStr)
			}
		}
	}

	for i, r := range ret {
		fmt.Printf("result %d: %+v\n", i, r)
	}
}

func isIPv4(str string) bool {
	return !strings.Contains(str, ":")
}

func isIPv6(str string) bool {
	return strings.Contains(str, ":")
}
