package hasher

import (
	"HashDehash/db"
	"crypto/sha512"
	"fmt"
)

func (h *Hasher) SHA512_256() (hash string) {
	dbToFrom := db.NewDB("hash_db/SHA512_256/to_from")
	dbFromTo := db.NewDB("hash_db/SHA512_256/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", sha512.Sum512_256([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
