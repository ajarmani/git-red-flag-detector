package scanner

import (
	"regexp"
	"strings"
)

// RedFlag represents a matched risky line
type RedFlag struct {
	FileName string
	Line     string
	Pattern  string
}

// Predefined risky patterns
var redFlagPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)api[_-]?key\s*=\s*['"]?[a-z0-9-_]{10,}`),
	regexp.MustCompile(`(?i)secret\s*=\s*['"]?[a-z0-9-_]{8,}`),
	regexp.MustCompile(`(?i)password\s*[:=]\s*['"]?.{6,}`),
	regexp.MustCompile(`(?i)access[_-]?token\s*[:=]\s*['"]?[a-z0-9-_]{10,}`),
	regexp.MustCompile(`(?i)BEGIN\s+PRIVATE\s+KEY`),
}

// ScanDiff scans the diff output for red flags
func ScanDiff(fileName, diff string) []RedFlag {
	var matches []RedFlag
	lines := strings.Split(diff, "\n")

	for _, line := range lines {
		if !strings.HasPrefix(line, "+") || strings.HasPrefix(line, "+++") {
			continue // Only scan added lines, skip diff metadata
		}

		cleanLine := strings.TrimPrefix(line, "+")

		for _, pattern := range redFlagPatterns {
			if pattern.MatchString(cleanLine) {
				matches = append(matches, RedFlag{
					FileName: fileName,
					Line:     cleanLine,
					Pattern:  pattern.String(),
				})
			}
		}
	}

	return matches
}
