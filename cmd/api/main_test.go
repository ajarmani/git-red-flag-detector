package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajarmani/git-red-flag-detector/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestScanEndpoint(t *testing.T) {
	repo, commits := testutils.SetupTestRepo(t)

	// Use last commit
	payload := map[string]string{
		"repo":   repo,
		"commit": commits[len(commits)-1],
	}
	body, _ := json.Marshal(payload)

	r := setupRouter()
	req := httptest.NewRequest(http.MethodPost, "/scan", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var flags []map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &flags)
	assert.NoError(t, err)
	assert.NotEmpty(t, flags, "should detect red flags in commit diff")
}
