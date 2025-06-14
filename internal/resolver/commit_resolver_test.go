package resolver

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

func setupTestRepo(t *testing.T) (string, []string) {
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

	// Init git repo
	run("git", "init")
	run("git", "config", "user.email", "test@example.com")
	run("git", "config", "user.name", "tester")

	// Create and commit files
	var hashes []string
	for i := 1; i <= 3; i++ {
		file := filepath.Join(dir, "file.txt")
		content := []byte("line " + strconv.Itoa(i) + "\n")
		err := os.WriteFile(file, content, 0644)
		if err != nil {
			t.Fatalf("write failed: %v", err)
		}
		run("git", "add", ".")
		run("git", "commit", "-m", "Commit "+strconv.Itoa(i))

		out, _ := exec.Command("git", "-C", dir, "rev-parse", "HEAD").Output()
		hashes = append(hashes, string(out[:len(out)-1]))
	}

	return dir, hashes
}

func TestGetLastNCommits(t *testing.T) {
	repo, hashes := setupTestRepo(t)
	got, err := GetLastNCommits(repo, 2)
	if err != nil {
		t.Fatalf("error getting commits: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 commits, got %d", len(got))
	}
	if got[0] != hashes[2] || got[1] != hashes[1] {
		t.Errorf("unexpected commit order. got %v", got)
	}
}

func TestGetCommitDiff(t *testing.T) {
	repo, hashes := setupTestRepo(t)
	diffs, err := GetCommitDiff(repo, hashes[2])
	if err != nil {
		t.Fatalf("error getting diff: %v", err)
	}
	if len(diffs) == 0 {
		t.Errorf("expected at least one file diff")
	}
}
