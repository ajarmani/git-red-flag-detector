package main

import (
	"fmt"

	"github.com/ajarmani/git-red-flag-detector/internal/resolver"
)

func main() {
	repoPath := "insert/local/path/to/your/repo"
	commitHash := "insert_commit_hash_here"

	diffs, err := resolver.GetCommitDiff(repoPath, commitHash)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, d := range diffs {
		fmt.Printf("File: %s\n", d.FileName)
		fmt.Println(d.Diff)
	}
}
