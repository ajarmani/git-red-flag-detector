# Git Red Flag Detector — v1.0

Detect hardcoded secrets, API keys, passwords, and other risky patterns in Git commits — via CLI or REST API.

---

## Overview

**Git Red Flag Detector** scans Git commit diffs for dangerous patterns like AWS keys, access tokens, private keys, etc., using a pluggable rule engine.

It supports both:

- **Command-line usage**
- **REST API endpoint**
- **Custom rule engine**
- **Git commit scanning**
- **Full unit and API test coverage**

---

## Features (v1)

- Scan any single commit or last N commits in a Git repo
- Rule-based detection engine (extensible)
- CLI and REST interface
- Automated tests using in-memory Git repos
- Pluggable rule system with metadata: RuleID, RuleDesc
- Modular, clean folder structure

---

## Tech Stack

| Layer         | Technology           |
|---------------|----------------------|
| Language      | Go (Golang)          |
| Git Resolver  | `git show` parsing   |
| API Server    | Chi + net/http       |
| Rule Engine   | Regex-based rules    |
| Tests         | `testing`+`httptest` |
| Repo Setup    | Temp Git repos       |

## API USAGE

- Start REST Server : go run cmd/api/main.go

- POST /scan

    curl -X POST http://localhost:8080/scan \
    -H "Content-Type: application/json" \
    -d '{
            "repo": "/absolute/path/to/git/repo",
            "commit": "commit_hash_here"
        }'

- Response format

    [
        {
            "fileName": "file.txt",
            "line": "secret = \"abc123\"",
            "ruleID": "SECRET",
            "ruleDesc": "Hardcoded secret"
        }
    ]


## CLI USAGE

- Scan particular commit hash
    go run cmd/server.go --repo /absolute/path/to/repo --commit <commit_hash>

- Scan last n commits
    go run cmd/server.go --repo /absolute/path/to/repo --last n


## Default Rules

| Rule ID       | Description         |
| ------------- | ------------------- |
| `AWS_KEY`     | AWS-style API keys  |
| `SECRET`      | Hardcoded secrets   |
| `PASSWORD`    | Hardcoded passwords |
| `TOKEN`       | Access tokens       |
| `PRIVATE_KEY` | Private key blocks  |

## License
- MIT — feel free to use, extend, or contribute.

## Next Steps (v2 and beyond)
- Storage for scan results
- Configurable rule
- Kafka/Redis ingestion pipeline
