package csv

import "strings"

func New(rawData string) ([][]string, error) {
	var output [][]string

	rawData = strings.TrimSpace(rawData)
	lines := strings.Split(rawData, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		output = append(output, strings.Split(line, ","))
	}

	return output, nil
}
