package handlers

import (
	"HashDehash/hasher"
	"HashDehash/utils"
	"encoding/json"
	"net/http"
)

func Hash(w http.ResponseWriter, r *http.Request) {
	params := utils.GetPathParams(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if len(params) == 0 || params[0] == "" {
			http.Error(w, "Hash string parameter not provided", 400)
			return
		}

		input := params[0]

		h := hasher.NewHasher()
		h.SetInput(input)
		h.Hash()

		hashes := h.GetHashes()

		hashList, err := json.MarshalIndent(hashes, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(hashList)

	default:
		http.Error(w, "Method not allowed", 405)
	}
}
