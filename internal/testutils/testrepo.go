package testutils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// SetupTestRepo creates a temporary Git repo with 3 commits
// and returns the path + list of commit hashes
func SetupTestRepo(t *testing.T) (string, []string) {
	dir, err := os.MkdirTemp("", "git-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	run := func(args ...string) {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v failed: %v\nOutput:\n%s", args, err, out)
		}
	}

	run("git", "init")
	run("git", "config", "user.email", "test@example.com")
	run("git", "config", "user.name", "tester")

	var hashes []string
	redFlagLines := [][]string{
		{
			`api_key = "AKIA1234567890ABCDEF"`,
			`secret = "mysecretvalue"`,
		},
		{
			`password = "hunter2"`,
			`access_token = "tok_abc_1234567890"`,
		},
		{
			`-----BEGIN PRIVATE KEY-----`,
			`MIIEv...FAKEKEY...AB`, // Simulated
			`-----END PRIVATE KEY-----`,
		},
	}

	for i, lines := range redFlagLines {
		content := strings.Join(append([]string{fmt.Sprintf("line %d", i+1)}, lines...), "\n")
		file := filepath.Join(dir, "file.txt")
		err := os.WriteFile(file, []byte(content), 0644)
		if err != nil {
			t.Fatalf("write failed: %v", err)
		}

		run("git", "add", ".")
		run("git", "commit", "-m", fmt.Sprintf("Commit %d", i+1))

		out, _ := exec.Command("git", "-C", dir, "rev-parse", "HEAD").Output()
		hashes = append(hashes, strings.TrimSpace(string(out)))
	}

	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	return dir, hashes
}
