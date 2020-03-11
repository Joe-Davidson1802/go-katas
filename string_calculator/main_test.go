package stringcalc

import (
	"testing"
	"math/rand"
	"strconv"
	"strings"
)

func getRandomInputs(count int) (int, string) {
	var nums []string
	total := 0

	for i := 1;  i<=count; i++ {
		curr := rand.Intn(1000)
		nums = append(nums, strconv.Itoa(curr))
		total += curr
	}

	res := strings.Join(nums, ",")

	return total, res
}

func TestMultipleNumberInputs(t *testing.T) {
	for i := 1;  i<=50; i++ {
		want, inp := getRandomInputs(rand.Intn(10))
		if got := Calculate(inp); got != want {
			t.Errorf("Calculate() = %d, want %d", got, want)
		}
	}
}

func BenchmarkOneThousandInputs(b *testing.B) {
	for i := 1;  i<=1000; i++ {
		want, inp := getRandomInputs(rand.Intn(10))
		if got := Calculate(inp); got != want {
			b.Errorf("Calculate() = %d, want %d", got, want)
		}
	}
}