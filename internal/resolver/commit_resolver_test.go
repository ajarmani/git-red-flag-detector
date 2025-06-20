package resolver

import (
	"os"
	"testing"

	"github.com/ajarmani/git-red-flag-detector/internal/testutils"
)

var testRepoPath string
var testCommitHashes []string

func TestMain(m *testing.M) {
	// Global test setup
	repo, hashes := testutils.SetupTestRepo(&testing.T{})
	testRepoPath = repo
	testCommitHashes = hashes

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func TestGetLastNCommits(t *testing.T) {
	got, err := GetLastNCommits(testRepoPath, 2)
	if err != nil {
		t.Fatalf("error getting commits: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 commits, got %d", len(got))
	}
	if got[0] != testCommitHashes[2] || got[1] != testCommitHashes[1] {
		t.Errorf("unexpected commit order. got %v", got)
	}
}

func TestGetCommitDiff(t *testing.T) {
	diffs, err := GetCommitDiff(testRepoPath, testCommitHashes[2])
	if err != nil {
		t.Fatalf("error getting diff: %v", err)
	}
	if len(diffs) == 0 {
		t.Errorf("expected at least one file diff")
	}
}
