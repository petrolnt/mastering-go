package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checkIP(ipParts []string) (string, int) {
	invalidPart := ""
	partIP := "(^25[0-5]$|^2[0-4][0-9]$|^1[0-9][0-9]$|^[1-9]?[0-9]$)"
	matchMe := regexp.MustCompile(partIP)
	octetNumber := 0
	for x, octet := range ipParts {

		isMatch := matchMe.MatchString(octet)
		if isMatch != true {
			invalidPart = octet
			octetNumber = x + 1
		}
	}
	return invalidPart, octetNumber
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage: checkIP.go someipaddress")
		os.Exit(1)
	}
	ip := arguments[1]
	ipParts := strings.Split(ip, ".")
	if len(ipParts) < 4 {
		fmt.Println("IP address is wrong")
		os.Exit(1)
	}
	invalidOctet, octetNumber := checkIP(ipParts)
	if len(invalidOctet) != 0 {
		fmt.Printf("The error in %v octet of IP address: %v\n", octetNumber, invalidOctet)
	} else {
		fmt.Printf("IP address %v is valid\n", ip)
	}

}
