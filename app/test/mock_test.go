package test

import (
	"bou.ke/monkey"
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func readFirstLine() string {
	open, err := os.Open("file.txt")
	defer open.Close()
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func replaceLine() string {
	line := readFirstLine()
	afterReplace := strings.ReplaceAll(line, "11", "00")
	return afterReplace
}

func TestReplace(t *testing.T) {
	monkey.Patch(readFirstLine, func() string {
		return "line11"
	})
	defer monkey.Unpatch(readFirstLine)
	res := replaceLine()
	assert.Equal(t, "line00", res)
}
