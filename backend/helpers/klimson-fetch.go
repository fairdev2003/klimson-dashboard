package helpers

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type SystemStats struct {
	SystemOS    string    `json:"system_os"`
	Arch        string    `json:"arch"`
	NumCPU      int       `json:"num_cpu"`
	Goroutines  int       `json:"goroutines"`
	MemoryAlloc string    `json:"memory_alloc"`
	Uptime      string    `json:"uptime"`
	Timestamp   time.Time `json:timestamp`
}

func GetUbuntuVersion() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Ubuntu (Unknown version)"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.Trim(parts[1], "\"")
			}
		}
	}
	return "Ubuntu"
}

func GetMemoryUsage() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "N/A"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
		}
		if strings.HasPrefix(line, "MemAvailable:") {
		}
	}
	return "Optimized via /proc/meminfo"
}

func GetSystemUptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "N/A"
	}
	parts := strings.Fields(string(data))
	if len(parts) > 0 {
		return parts[0] + " seconds"
	}
	return "N/A"
}
