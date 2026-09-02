package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"unicode/utf8"
)

func main() {
	// Route
	http.HandleFunc("POST /tokenize", tokenizeHandler)

	// Start server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// get the body from request
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024) // 1MB
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	if !utf8.Valid(body) { // protect against XSS
		http.Error(w, "Invalid UTF-8 encoded text", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// tokenize
	tokens, err := tokenize(body)
	if err != nil {
		http.Error(w, "Tokenizer failed", http.StatusInternalServerError)
		return
	}

	type TokenResponse struct {
		Tokens []string `json:"tokens"`
	}
	res := TokenResponse{Tokens: tokens}

	// marshalling
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
}

func tokenize(s []byte) ([]string, error) {
	re := regexp.MustCompile(`[a-zA-z]+`)
	tokens := re.FindAllString(string(s), -1)
	// TODO error handling?
	return tokens, nil
}
