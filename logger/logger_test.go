package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestInfo(t *testing.T) {
	logVal := "Testing info log"
	output := captureOutput(func() {
		Info(logVal)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[INFO] [%s]\n", logVal)))

	log2Val := "Second info log"
	output = captureOutput(func() {
		Info(logVal, log2Val)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[INFO] [%s %s]\n", logVal, log2Val)))
}

func TestDebug(t *testing.T) {
	logVal := "Testing debug log"
	output := captureOutput(func() {
		Debug(logVal)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[DEBUG] [%s]\n", logVal)))

	log2Val := "Second debug log"
	output = captureOutput(func() {
		Debug(logVal, log2Val)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[DEBUG] [%s %s]\n", logVal, log2Val)))
}

func TestError(t *testing.T) {
	logVal := "Testing error log"
	output := captureOutput(func() {
		Error(logVal)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[ERROR] [%s]\n", logVal)))

	log2Val := "Second error log"
	output = captureOutput(func() {
		Error(logVal, log2Val)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("[ERROR] [%s %s]\n", logVal, log2Val)))
}

func TestPrint(t *testing.T) {
	logVal := "Testing print log"
	output := captureOutput(func() {
		Print(logVal)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("%s\n", logVal)))

	log2Val := "Second print log"
	output = captureOutput(func() {
		Print(logVal, log2Val)
	})

	assert.True(t, strings.HasSuffix(output, fmt.Sprintf("%s%s\n", logVal, log2Val)))
}
