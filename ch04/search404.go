package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type logRecord struct {
	ipaddress    string
	reqTime      string
	requestedURL string
}

func checkString(input string) bool {
	partIP := "(\" 404 \\d{3})"
	matchMe := regexp.MustCompile(partIP)
	//fmt.Printf("checking string: %v\n", input)
	return matchMe.MatchString(input)
}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func findEventTime(input string) string {
	timReg := "(\\d{2}\\/\\w{3}\\/\\d{4}:\\d{2}:\\d{2}:\\d{2} \\+\\d{4})"
	matchMe := regexp.MustCompile(timReg)
	return matchMe.FindString(input)
}

func findURL(input string) string {
	urlReg := "(\\b[A-Z]{3,} .* HTTP\\/\\d\\.\\d)"
	matchMe := regexp.MustCompile(urlReg)
	return matchMe.FindString(input)
}

func getRecords(accessLogPath string) map[string]logRecord {
	data := make(map[string]logRecord)
	f, err := os.Open(accessLogPath)
	if err != nil {
		fmt.Printf("Error in opening access log file %v\n", accessLogPath)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		if checkString(line) {
			ip := findIP(line)
			eventTime := findEventTime(line)
			url := findURL(line)

		}

	}
	return data
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage: topIP.go path_to_accelog_file")
		os.Exit(1)
	}
	aLogPath := arguments[1]

	getRecords(aLogPath)

}
