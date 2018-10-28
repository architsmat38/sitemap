package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	minVal := 1
	maxVal := 5
	randVal := GetRandom(minVal, maxVal)
	assert.True(
		t,
		(randVal >= minVal && randVal <= maxVal),
		fmt.Sprintf("Expected randVal between %d and %d, got %d", minVal, maxVal, randVal),
	)
}

func TestIsValidURL(t *testing.T) {
	assert.True(t, IsValidURL("http://www.example.com/"))
	assert.True(t, IsValidURL("http://example.com/"))
	assert.True(t, IsValidURL("http://example.com/xyz?a=1"))
	assert.False(t, IsValidURL("example.com"))
	assert.False(t, IsValidURL("example"))
}

func TestGetHost(t *testing.T) {
	host, err := GetHost("http://example.com/welcome.html")
	assert.NoError(t, err)
	expectedHost := "example.com"
	assert.Equal(t, expectedHost, host, fmt.Sprintf("Expected %s, got %s", expectedHost, host))
}

func TestCreateValidURL(t *testing.T) {
	assert.Equal(t, CreateValidURL("http://example.com/xyz.html", "example.com"), "http://example.com/xyz.html")
	assert.Equal(t, CreateValidURL("/xyz.html", "example.com"), "http://example.com/xyz.html")
	assert.Equal(t, CreateValidURL("xyz", "example.com"), "http://example.com/xyz")
}

func TestGetCrawlableURL(t *testing.T) {
	assert.Equal(t, GetCrawlableURL("http://example.com/"), "http://example.com")
	assert.Equal(t, GetCrawlableURL("http://example.com/xyz.php"), "http://example.com/xyz.php")
	assert.Equal(t, GetCrawlableURL("http://example.com/abc?q=test"), "http://example.com/abc")
	assert.Equal(t, GetCrawlableURL("http://example.com/abc?q=test#1231231"), "http://example.com/abc")
}
