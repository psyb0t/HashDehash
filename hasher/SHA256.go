package hasher

import (
	"HashDehash/db"
	"crypto/sha256"
	"fmt"
)

func (h *Hasher) SHA256() (hash string) {
	dbToFrom := db.NewDB("hash_db/SHA256/to_from")
	dbFromTo := db.NewDB("hash_db/SHA256/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", sha256.Sum256([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
