package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func (h *Hasher) RIPEMD160() (hash string) {
	dbToFrom := db.NewDB("hash_db/RIPEMD160/to_from")
	dbFromTo := db.NewDB("hash_db/RIPEMD160/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hasher := ripemd160.New()
	hasher.Write([]byte(h.input))

	hash = fmt.Sprintf("%x", hasher.Sum(nil))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
