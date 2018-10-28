package utils

import (
	"fmt"
	"strings"
)

/**
 * Get random proxy user agent
 */
func GetProxyUserAgent() string {
	var availableUserAgents = []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.4.8 (KHTML, like Gecko) Version/10.0.3 Safari/602.4.8",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
		"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	}

	var totalAvailableAgents = len(availableUserAgents)
	var randomIndex = GetRandom(1, totalAvailableAgents)
	return availableUserAgents[randomIndex-1]
}

/**
 * Get random proxy IP
 */
func GetRandomIP() string {
	ipValues := []int{GetRandom(1, 255), GetRandom(0, 255), GetRandom(0, 255), GetRandom(0, 255)}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ipValues)), "."), "[]")
}
