package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <nmap-output-file>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Maps to store open ports by protocol
	tcpPorts := make(map[int]bool)
	udpPorts := make(map[int]bool)

	// Regex to match port lines
	// Matches lines like: "3128/tcp  open  squid-http"
	portRegex := regexp.MustCompile(`^(\d+)/(tcp|udp)\s+open\s+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and headers
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "Nmap") || 
		   strings.HasPrefix(line, "Host") || strings.HasPrefix(line, "Other") || 
		   strings.HasPrefix(line, "Not shown") || strings.HasPrefix(line, "PORT") {
			continue
		}

		// Check if line matches port pattern
		matches := portRegex.FindStringSubmatch(line)
		if matches != nil {
			port, err := strconv.Atoi(matches[1])
			if err != nil {
				continue
			}

			protocol := matches[2]
			if protocol == "tcp" {
				tcpPorts[port] = true
			} else if protocol == "udp" {
				udpPorts[port] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert maps to sorted slices
	var tcpList []int
	for port := range tcpPorts {
		tcpList = append(tcpList, port)
	}
	sort.Ints(tcpList)

	var udpList []int
	for port := range udpPorts {
		udpList = append(udpList, port)
	}
	sort.Ints(udpList)

	// Build output string
	var output []string

	if len(tcpList) > 0 {
		tcpStr := "T:"
		for i, port := range tcpList {
			if i > 0 {
				tcpStr += ","
			}
			tcpStr += strconv.Itoa(port)
		}
		output = append(output, tcpStr)
	}

	if len(udpList) > 0 {
		udpStr := "U:"
		for i, port := range udpList {
			if i > 0 {
				udpStr += ","
			}
			udpStr += strconv.Itoa(port)
		}
		output = append(output, udpStr)
	}

	// Print the result
	if len(output) > 0 {
		fmt.Println(strings.Join(output, ","))
	}
}
