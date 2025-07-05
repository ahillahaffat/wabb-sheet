package main

import "strings"

func parseMessage(message string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(message, "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}

	return result
}
