package test

import (
	"strings"
	"testing"
)

/**
go test -bench .
*/

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	for i := 2; i < 10; i++ {
		Fib(5)
	}
}

func AppendStrAdd(n int, s string) string {
	res := s
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}

func AppendStrBuffer(n int, str string) string {

	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 基准测试 AppendStrAdd
func BenchmarkAppendStrAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendStrAdd(1000, "a") // 这里的 1000 是拼接次数，可以调整
	}
}

// 基准测试 AppendStrBuffer
func BenchmarkAppendStrBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendStrBuffer(1000, "a") // 这里的 1000 是拼接次数，可以调整
	}
}
