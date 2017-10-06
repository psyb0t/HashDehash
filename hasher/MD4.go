package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/md4"
)

func (h *Hasher) MD4() (hash string) {
	dbToFrom := db.NewDB("hash_db/MD4/to_from")
	dbFromTo := db.NewDB("hash_db/MD4/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	md4hasher := md4.New()
	md4hasher.Write([]byte(h.input))

	hash = fmt.Sprintf("%x", md4hasher.Sum(nil))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
