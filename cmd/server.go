package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ajarmani/git-red-flag-detector/internal/resolver"
	"github.com/ajarmani/git-red-flag-detector/internal/scanner"
)

func main() {
	repoPath := flag.String("repo", "", "Absolute path to the Git repo")
	commitHash := flag.String("commit", "", "Commit hash to scan")
	flag.Parse()

	if *repoPath == "" || *commitHash == "" {
		fmt.Println("Usage: go run cmd/server.go --repo /path/to/repo --commit <hash>")
		os.Exit(1)
	}

	diffs, err := resolver.GetCommitDiff(*repoPath, *commitHash)
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
