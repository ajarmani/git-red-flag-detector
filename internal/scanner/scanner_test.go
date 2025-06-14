package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanDiff(t *testing.T) {
	diff := `
diff --git a/secret.txt b/secret.txt
new file mode 100644
index 0000000..e69de29
+++ b/secret.txt
+api_key = "AKIA1234567890123456"
+password: "supersecret"
+normal_line = "this is fine"
+access_token: "abc123def456"
+BEGIN PRIVATE KEY
`
	expectedPatterns := []string{
		`(?i)api[_-]?key\s*=\s*['"]?[a-z0-9-_]{10,}`,
		`(?i)password\s*[:=]\s*['"]?.{6,}`,
		`(?i)access[_-]?token\s*[:=]\s*['"]?[a-z0-9-_]{10,}`,
		`(?i)BEGIN\s+PRIVATE\s+KEY`,
	}

	flags := ScanDiff("secret.txt", diff)

	assert.Equal(t, 4, len(flags), "Expected 4 red flags")
	matched := make(map[string]bool)
	for _, f := range flags {
		matched[f.Pattern] = true
		assert.Equal(t, "secret.txt", f.FileName)
		assert.NotEmpty(t, f.Line)
	}

	for _, pattern := range expectedPatterns {
		assert.True(t, matched[pattern], "Pattern not matched: "+pattern)
	}
}
