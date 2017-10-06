package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func (h *Hasher) SHA3_224() (hash string) {
	dbToFrom := db.NewDB("hash_db/SHA3_224/to_from")
	dbFromTo := db.NewDB("hash_db/SHA3_224/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", sha3.Sum224([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
