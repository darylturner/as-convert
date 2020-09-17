package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertASplain(i string) (string, error) {
	x, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		return "", fmt.Errorf("failed to convert asplain to integer: %w", err)
	}
	if x > 4294967296 {
		return "", fmt.Errorf("not valid 4 byte asn")
	}

	left := (x / 65535)
	right := (x - (65535 * left) - left)
	asdot := fmt.Sprintf("%v.%v", left, right)

	return asdot, nil
}

func splitASdot(i string) ([]uint64, error) {
	elements := make([]uint64, 0)

	split := strings.Split(i, ".")
	if len(split) != 2 {
		return elements, fmt.Errorf("unexpected number of asdot elements")
	}

	for _, i := range split {
		x, err := strconv.ParseUint(i, 10, 64)
		if err != nil {
			return elements, fmt.Errorf("failed to convert asdot elements to integer: %w", err)
		}
		if x > 65535 {
			return elements, fmt.Errorf("asdot element(s) invalid")
		}
		elements = append(elements, x)
	}
	return elements, nil
}

func convertASdot(i string) (uint64, error) {
	e, err := splitASdot(i)
	if err != nil {
		return 0, fmt.Errorf("asdot split failed: %w", err)
	}

	asplain := ((e[0] * 65535) + e[1] + e[0])
	return asplain, nil
}

func main() {
	input := os.Args[1]
	if strings.Contains(input, ".") {
		out, err := convertASdot(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	} else {
		out, err := convertASplain(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	}
}
