package utils

import (
	"errors"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

/**
 * Get random number between given range
 */
func GetRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/**
 * Validates the provided URL
 */
func IsValidURL(linkURL string) bool {
	u, err := url.Parse(linkURL)
	if err != nil || len(u.Hostname()) == 0 {
		return false
	}
	return true
}

/**
 * Get host for the provided URL
 */
func GetHost(linkURL string) (string, error) {
	if IsValidURL(linkURL) {
		u, _ := url.Parse(linkURL)
		return u.Hostname(), nil
	}
	return "", errors.New("Invalid URL")
}

/**
 * Create valid url w.r.t. website URL
 */
func CreateValidURL(linkURL string, websiteHost string) string {
	u, err := url.Parse(linkURL)
	if err != nil {
		return linkURL
	}

	baseURL, err := url.Parse("http://" + websiteHost)
	if err != nil {
		return websiteHost
	}

	validURL := baseURL.ResolveReference(u)
	validURL.Scheme = "http"
	return validURL.String()
}

/**
 * Get crawlable URL
 */
func GetCrawlableURL(linkURL string) string {
	if IsValidURL(linkURL) {
		u, _ := url.Parse(linkURL)
		u.Scheme = "http"
		u.RawQuery = ""
		u.Fragment = ""
		return strings.TrimRight(u.String(), "/")
	}
	return linkURL
}
