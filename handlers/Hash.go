package handlers

import (
	"HashDehash/hasher"
	"HashDehash/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Hash(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		wordHash(w, r)
	case http.MethodPost:
		wordlistHash(w, r)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func wordHash(w http.ResponseWriter, r *http.Request) {
	params := utils.GetPathParams(r.URL.Path)
	if len(params) == 0 || params[0] == "" {
		http.Error(w, "Word parameter not provided", 400)
		return
	}

	word := params[0]

	h := hasher.NewHasher()
	h.SetInput(word)
	h.Hash()

	hashes := h.GetHashes()

	hashList, err := json.MarshalIndent(hashes, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(hashList)
}

func wordlistHash(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("wordlist")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()
	io.Copy(buf, file)

	words := strings.Split(buf.String(), "\n")
	count := 0
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		count += 1

		h := hasher.NewHasher()
		h.SetInput(word)
		h.Hash()
	}

	w.Write([]byte(fmt.Sprintf("Successfully hashed %d words", count)))
}
