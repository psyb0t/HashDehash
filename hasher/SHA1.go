package hasher

import (
	"HashDehash/db"
	"crypto/sha1"
	"fmt"
)

func (h *Hasher) SHA1() (hash string) {
	dbToFrom := db.NewDB("hash_db/SHA1/to_from")
	dbFromTo := db.NewDB("hash_db/SHA1/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", sha1.Sum([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
