package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/blake2b"
)

func (h *Hasher) BLAKE2b_256() (hash string) {
	dbToFrom := db.NewDB("hash_db/BLAKE2b_256/to_from")
	dbFromTo := db.NewDB("hash_db/BLAKE2b_256/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", blake2b.Sum256([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
