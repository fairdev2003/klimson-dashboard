package helpers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type SystemStats struct {
	SystemOS  string `json:"system_os"`
	GoVersion string `json:"go_version"`
	Arch      string `json:"arch"`
	NumCPU    int    `json:"num_cpu"`

	Goroutines  int `json:"goroutines"`
	ThreadCount int `json:"thread_count"`

	MemoryAlloc string `json:"memory_alloc"`
	MemoryTotal string `json:"memory_total"`
	MemorySys   string `json:"memory_sys"`
	HeapObjects uint64 `json:"heap_objects"`
	NumGC       uint32 `json:"num_gc"`

	Uptime     string    `json:"uptime"`
	ServerTime time.Time `json:"server_time"`
	Timestamp  time.Time `json:"timestamp"`
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
		return parts[0]
	}
	return "N/A"
}
func FormatBytes(bytes uint64) string {
	if bytes == 0 {
		return "0 B"
	}

	base := math.Log(float64(bytes)) / math.Log(1024)
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB"}

	i := int(math.Floor(base))
	if i >= len(sizes) {
		i = len(sizes) - 1
	}

	val := float64(bytes) / math.Pow(1024, float64(i))

	if i == 0 {
		return fmt.Sprintf("%d %s", int(bytes), sizes[i])
	}

	return fmt.Sprintf("%.2f %s", val, sizes[i])
}
