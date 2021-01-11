package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

type record struct {
	ipaddress string
	count     int
}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func createTopList(accessLogPath string) map[string]int {
	TopList := make(map[string]int)

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
		ip := findIP(line)
		count := lookupRecord(ip, TopList)
		TopList[ip] = count + 1

	}
	return TopList
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage: topIP.go path_to_accelog_file")
		os.Exit(1)
	}
	aLogPath := arguments[1]

	topList := createTopList(aLogPath)
	printSortedMap(topList)

}

func lookupRecord(k string, m map[string]int) int {

	n, ok := m[k]
	if !ok {

		n = 0
	}
	return n
}

func printSortedMap(unsortedMap map[string]int) {
	keys := make([]string, 0, len(unsortedMap))
	for key := range unsortedMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return unsortedMap[keys[i]] > unsortedMap[keys[j]] })

	for n, key := range keys {
		if n > 5 {
			break
		}
		fmt.Printf("IP: %v Count: %v\n", key, unsortedMap[key])
	}

}
