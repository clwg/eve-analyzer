package domainsuffix

import (
	"testing"
)

func TestParseDomain(t *testing.T) {
	// Mock the publicSuffixes for testing without needing the actual file
	publicSuffixes = []string{"co.uk", "com", "uk.co"}

	tests := []struct {
		domain         string
		expectedSuffix string
		expectedPart   string
		description    string
	}{
		{"www.example.com", "com", "example.com", "Standard domain"},
		{"www.example.co.uk", "co.uk", "example.co.uk", "Domain with a second-level public suffix"},
		{"w1.w2.w3.example.uk.co", "uk.co", "example.uk.co", "Domain with an unusual public suffix"},
		{"example.com", "com", "example.com", "Domain without subdomains"},
		{"www.example", "", "www.example", "Domain not in list"},
	}

	for _, test := range tests {
		delegatedPart, suffix, err := ParseDomain(test.domain)
		if err != nil {
			t.Errorf("Test '%s' failed with error: %s", test.description, err)
		}
		if delegatedPart != test.expectedPart || suffix != test.expectedSuffix {
			t.Errorf("Test '%s' failed: expected (%s, %s), got (%s, %s)", test.description, test.expectedPart, test.expectedSuffix, delegatedPart, suffix)
		}
	}
}
