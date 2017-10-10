package handlers

import (
	"HashDehash/db"
	"HashDehash/utils"
	"io/ioutil"
	"net/http"
	"path"
)

func Dehash(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := utils.GetPathParams(r.URL.Path)
		if len(params) == 0 || params[0] == "" {
			http.Error(w, "Hash string parameter not provided", 400)
			return
		}

		hash := params[0]

		algo_dirs, err := ioutil.ReadDir("hash_db")
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		for _, algo_dir := range algo_dirs {
			if algo_dir.IsDir() {
				dbToFrom := db.NewDB(path.Join("hash_db",
					algo_dir.Name(), "to_from"))

				word, err := dbToFrom.Get(hash)

				if err == nil {
					w.Write([]byte(word))
					return
				}
			}
		}

		http.Error(w, "Hash original input not found", 404)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}
