package handlers

import (
	"HashDehash/db"
	"HashDehash/utils"
	"net/http"
)

func Dehash(w http.ResponseWriter, r *http.Request) {
	dbMD5ToFrom := db.NewDB("hash_db/md5/to_from")

	params := utils.GetPathParams(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if len(params) == 0 || params[0] == "" {
			http.Error(w, "Hash string parameter not provided", 400)
			return
		}

		hash := params[0]
		input, err := dbMD5ToFrom.Get(hash)

		if err != nil {
			http.Error(w, "Hash original input not found", 404)
			return
		}

		w.Write([]byte(input))
	default:
		http.Error(w, "Method not allowed", 405)
	}
}
