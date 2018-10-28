package utils

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProxyUserAgent(t *testing.T) {
	randUserAgent := GetProxyUserAgent()
	assert.NotEmpty(t, randUserAgent)
	assert.Equal(t, reflect.TypeOf(randUserAgent).Name(), "string")
}

func TestGetRandomIP(t *testing.T) {
	randIP := GetRandomIP()
	randIPParts := strings.Split(randIP, ".")
	assert.Equal(t, len(randIPParts), 4)

	for _, num := range randIPParts {
		_, err := strconv.Atoi(num)
		assert.NoError(t, err)
	}
}
