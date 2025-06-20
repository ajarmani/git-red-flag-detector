package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ajarmani/git-red-flag-detector/internal/resolver"
	"github.com/ajarmani/git-red-flag-detector/internal/scanner"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ScanRequest struct {
	RepoPath   string `json:"repo"`
	CommitHash string `json:"commit"`
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	diffs, err := resolver.GetCommitDiff(req.RepoPath, req.CommitHash)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error resolving commit: %v", err), http.StatusInternalServerError)
		return
	}

	var allFlags []scanner.RedFlag
	for _, d := range diffs {
		flags := scanner.ScanDiff(d.FileName, d.Diff)
		allFlags = append(allFlags, flags...)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allFlags)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/scan", scanHandler)

	fmt.Println("REST server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
