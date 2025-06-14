package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanWithDefaultRules(t *testing.T) {
	diff := `
+++ b/config.txt
+api_key = "AKIA1234567890123456"
+password: "hunter2"
+BEGIN PRIVATE KEY
`
	flags := ScanDiff("config.txt", diff)

	assert.Equal(t, 3, len(flags))
	assert.Equal(t, "config.txt", flags[0].FileName)
	assert.NotEmpty(t, flags[0].RuleID)
	assert.NotEmpty(t, flags[0].RuleDesc)
}
