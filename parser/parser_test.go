package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/ashwanthkumar/golang-utils/io"
	"github.com/stretchr/testify/assert"
)

func TestParseAllURLs(t *testing.T) {
	htmlBody, err := io.ReadFullyFromFile("../tests/golangcode.html")
	assert.NoError(t, err)

	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	assert.NoError(t, err)

	websiteHost := "golangcode.com"
	allURLs := ParseAllURLs(document, websiteHost)
	expectedURLs := []string{
		"http://golangcode.com/tags",
		"http://golangcode.com/about",
		"http://golangcode.com/search",
		"http://golangcode.com/index.xml",
		"http://golangcode.com/about",
		"http://golangcode.com/parsing-dates",
		"http://golangcode.com/execute-a-command",
		"http://golangcode.com/sorting-an-array-of-numeric-items",
		"http://golangcode.com/creating-temp-files",
		"http://golangcode.com/base-64-encode-decode",
		"http://golangcode.com/http-get-with-timeout",
		"http://golangcode.com/get-http-method-from-request",
		"http://golangcode.com/lambda-pdf-generator-from-s3",
		"http://golangcode.com/substring-num-of-characters",
		"http://golangcode.com/middleware-on-handlers",
		"http://golangcode.com/page/2",
		"http://golangcode.com/index.xml",
		"http://golangcode.com/privacy",
	}
	assert.Equal(t, expectedURLs, allURLs)
}
