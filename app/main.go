package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Version struct {
	App         string `json:"app"`
	Environment string `json:"environment"`
	PRNumber    string `json:"pr_number"`
	GitSHA      string `json:"git_sha"`
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	version := Version{
		App:         "pr-preview-platform",
		Environment: getEnv("APP_ENV", "development"),
		PRNumber:    getEnv("PR_NUMBER", "local"),
		GitSHA:      getEnv("GIT_SHA", "local"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>PR Preview Platform</title></head>
<body>
<h1>PR Preview Platform</h1>
<p>Status: Running</p>
<p>Environment: %s</p>
<p>Pull Request: %s</p>
<p>Git Commit: %s</p>
</body>
</html>`, version.Environment, version.PRNumber, version.GitSHA)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(version)
	})

	port := getEnv("PORT", "8080")
	log.Printf("Starting preview application on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
