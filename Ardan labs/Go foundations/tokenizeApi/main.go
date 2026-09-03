package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"unicode/utf8"
)

func main() {
	api := API{
		log: slog.Default().With("component", "HTTP server"),
	}

	// Route
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)

	// Start server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

type API struct {
	log *slog.Logger
}

func (l *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// get the body from request
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024) // 1MB
	body, err := io.ReadAll(r.Body)
	if err != nil {
		l.log.Warn("Payload too large", "remote_addr", r.RemoteAddr, "", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if !utf8.Valid(body) { // protect against XSS
		l.log.Info("Invalid UTF-8 encoded text:", "remote_addr", r.RemoteAddr, ":", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// tokenize
	res := map[string][]string{
		"tokens": l.tokenize(body),
	}

	// marshalling
	if err := json.NewEncoder(w).Encode(res); err != nil {
		l.log.Error("Failed to marshal JSON", "", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func (l *API) tokenize(s []byte) []string {
	return regexp.MustCompile(`[a-zA-z]+`).FindAllString(string(s), -1)
}
