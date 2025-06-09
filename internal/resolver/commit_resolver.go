package resolver

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type FileDiff struct {
	FileName string
	Diff     string
}

// GetCommitDiff returns diffs for all files changed in a commit.
func GetCommitDiff(repoPath, commitHash string) ([]FileDiff, error) {
	cmd := exec.Command("git", "-C", repoPath, "show", commitHash, "--pretty=format:", "--unified=0")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to get git diff: %v", err)
	}

	raw := out.String()
	files := strings.Split(raw, "diff --git ")

	var result []FileDiff

	for _, fileBlock := range files[1:] {
		lines := strings.SplitN(fileBlock, "\n", 2)
		header := lines[0]
		body := ""
		if len(lines) > 1 {
			body = lines[1]
		}

		// Parse filename from header like: a/path/to/file b/path/to/file
		parts := strings.Fields(header)
		if len(parts) < 2 {
			continue
		}

		filename := strings.TrimPrefix(parts[1], "b/")
		result = append(result, FileDiff{
			FileName: filename,
			Diff:     body,
		})
	}

	return result, nil
}
