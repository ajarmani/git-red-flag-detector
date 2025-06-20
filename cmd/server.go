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
	lastN := flag.Int("last", 0, "Number of latest commits to scan")

	flag.Parse()

	if *repoPath == "" || (*commitHash == "" && *lastN == 0) {
		fmt.Println("Usage: go run cmd/server.go --repo /path/to/repo --commit <hash> OR --last <n>")
		os.Exit(1)
	}

	if *lastN > 0 {
		commits, err := resolver.GetLastNCommits(*repoPath, *lastN)
		if err != nil {
			fmt.Println("Error fetching commits:", err)
			os.Exit(1)
		}

		for _, hash := range commits {
			fmt.Printf("\nScanning commit: %s\n", hash)
			diffs, err := resolver.GetCommitDiff(*repoPath, hash)
			if err != nil {
				fmt.Printf("Error scanning %s: %v\n", hash, err)
				continue
			}

			for _, d := range diffs {
				flags := scanner.ScanDiff(d.FileName, d.Diff)
				for _, flag := range flags {
					fmt.Printf("RED FLAG in %s:\n   ➤ Rule: %s (%s)\n   ➤ Line: %s\n\n",
						flag.FileName, flag.RuleID, flag.RuleDesc, flag.Line)
				}
			}
		}
	} else {
		diffs, err := resolver.GetCommitDiff(*repoPath, *commitHash)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, d := range diffs {
			flags := scanner.ScanDiff(d.FileName, d.Diff)
			for _, flag := range flags {
				fmt.Printf("RED FLAG in %s:\n   ➤ Rule: %s (%s)\n   ➤ Line: %s\n\n",
					flag.FileName, flag.RuleID, flag.RuleDesc, flag.Line)
			}
		}
	}
}
