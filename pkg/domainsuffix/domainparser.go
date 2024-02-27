package domainsuffix

import (
	"strings"
)

// ParseDomain attempts to find the domain's suffix in the public suffix list.
// If found, it returns the first delegated part along with the suffix.
// If not found, it returns the last two parts of the domain as a fallback.
func ParseDomain(domain string) (string, string, error) {
	// Ensure the list is loaded.
	initOnce.Do(loadPublicSuffixList)

	parts := strings.Split(domain, ".")
	for i := range parts {
		potentialSuffix := strings.Join(parts[i:], ".")
		if isPublicSuffix(potentialSuffix) {
			if i > 0 {
				delegatedPart := strings.Join(parts[i-1:], ".")
				return delegatedPart, potentialSuffix, nil
			}
			break
		}
	}
	// If no matching public suffix, use the last two parts as a fallback
	if len(parts) >= 2 {
		fallback := strings.Join(parts[len(parts)-2:], ".")
		return fallback, "", nil // Returning an empty string for the suffix to indicate a fallback scenario
	}
	return domain, "", nil // Fallback to the original domain if it's too short, with an empty suffix
}

// isPublicSuffix checks if the given domain part is in the public suffix list.
func isPublicSuffix(domainPart string) bool {
	for _, suffix := range publicSuffixes {
		if domainPart == suffix {
			return true
		}
	}
	return false
}
