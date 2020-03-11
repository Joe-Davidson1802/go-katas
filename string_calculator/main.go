package stringcalc

import (
	"strconv"
	"strings"
)

func Calculate(inp string) int {
	if inp == "" {
		return 0
	}

	nums := strings.Split(inp, ",")
	total := 0

	for _, num := range nums {
		numVal, _ := strconv.Atoi(num)
		total += numVal
	}

	return total
}