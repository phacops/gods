package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Hash struct {
	values map[string]string
}

func parseFile(path string) *Hash {
	file, err := os.Open(path)
	hash := &Hash{values: make(map[string]string)}

	if err != nil {
		return hash
	}

	defer file.Close()

	var buffer []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		buffer = strings.Split(scanner.Text(), "=")

		if len(buffer) == 2 {
			hash.values[buffer[0]] = buffer[1]
		}
	}

	return hash
}

func (h *Hash) SearchForInt(fields []string) int {
	for _, field := range fields {
		if _, exists := h.values[field]; exists {
			return h.GetInt(field)
		}
	}

	return 0
}

func (h *Hash) GetInt(field string) int {
	if convertedValue, err := strconv.Atoi(h.values[field]); err == nil {
		return convertedValue
	}

	return 0
}
