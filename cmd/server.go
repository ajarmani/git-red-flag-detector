package main

import (
	"fmt"

	"github.com/ajarmani/git-red-flag-detector/internal/resolver"
	"github.com/ajarmani/git-red-flag-detector/internal/scanner"
)

func main() {
	// For testing purposes, using hello-world repo
	repoPath := "/Users/mbalakrishna/Desktop/Desktop/Golang/hello-world"
	commitHash := "2215cb0"

	diffs, err := resolver.GetCommitDiff(repoPath, commitHash)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, d := range diffs {
		flags := scanner.ScanDiff(d.FileName, d.Diff)
		for _, flag := range flags {
			fmt.Printf("RED FLAG in %s:\n   ➤ Pattern: %s\n   ➤ Line: %s\n\n", flag.FileName, flag.Pattern, flag.Line)
		}
	}
}
