package scanner

import (
	"strings"
)

// RedFlag represents a matched risky line
type RedFlag struct {
	FileName string
	Line     string
	RuleID   string
	RuleDesc string
}

// Predefined red flag rules
var defaultRules = []Rule{
	NewRegexRule("AWS_KEY", "AWS API key", `(?i)api[_-]?key\s*=\s*['"]?[a-z0-9-_]{10,}`),
	NewRegexRule("SECRET", "Hardcoded secret", `(?i)secret\s*=\s*['"]?[a-z0-9-_]{8,}`),
	NewRegexRule("PASSWORD", "Hardcoded password", `(?i)password\s*[:=]\s*['"]?.{6,}`),
	NewRegexRule("TOKEN", "Access token", `(?i)access[_-]?token\s*[:=]\s*['"]?[a-z0-9-_]{10,}`),
	NewRegexRule("PRIVATE_KEY", "Private Key block", `(?i)BEGIN\s+PRIVATE\s+KEY`),
}

// ScanDiff scans a diff using the given rules
func ScanDiff(fileName, diff string) []RedFlag {
	return ScanWithRules(fileName, diff, defaultRules)
}

// ScanWithRules allows custom rule scanning
func ScanWithRules(fileName, diff string, rules []Rule) []RedFlag {
	var matches []RedFlag
	lines := strings.Split(diff, "\n")

	for _, line := range lines {
		if !strings.HasPrefix(line, "+") || strings.HasPrefix(line, "+++") {
			continue
		}

		cleanLine := strings.TrimPrefix(line, "+")

		for _, rule := range rules {
			if rule.Match(cleanLine) {
				matches = append(matches, RedFlag{
					FileName: fileName,
					Line:     cleanLine,
					RuleID:   rule.ID(),
					RuleDesc: rule.Description(),
				})
			}
		}
	}

	return matches
}
