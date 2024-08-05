package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const hostsFilePath = `C:\Windows\System32\drivers\etc\hosts`
const loopbackAddress = "127.0.0.1"

func addToHostsFile(domain string) bool {
	f, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
		return false
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s %s\n", loopbackAddress, domain))

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func removeFromHostsFile(domain string) error {
	contents, err := os.ReadFile(hostsFilePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(contents), "\n")
	var newLines []string
	toRemove := fmt.Sprintf("%s %s", loopbackAddress, domain)

	for _, line := range lines {
		if strings.TrimSpace(line) != toRemove {
			newLines = append(newLines, line)
		}
	}

	return os.WriteFile(hostsFilePath, []byte(strings.Join(newLines, "\n")), 0644)
}

func getBlockedDomains() ([]string, error) {
	contents, err := os.ReadFile(hostsFilePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(contents), "\n")
	var blockedDomains []string

	for _, line := range lines {
		if strings.HasPrefix(line, loopbackAddress) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				domain := fields[1]
				blockedDomains = append(blockedDomains, domain)
			}
		}
	}

	return blockedDomains, nil
}
