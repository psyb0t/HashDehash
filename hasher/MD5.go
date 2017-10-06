package hasher

import (
	"HashDehash/db"
	"crypto/md5"
	"fmt"
)

func (h *Hasher) MD5() (hash string) {
	dbToFrom := db.NewDB("hash_db/MD5/to_from")
	dbFromTo := db.NewDB("hash_db/MD5/from_to")

	hash, err := dbFromTo.Get(h.input)
	if err == nil {
		return
	}

	hash = fmt.Sprintf("%x", md5.Sum([]byte(h.input)))

	dbToFrom.Set(hash, h.input)
	dbFromTo.Set(h.input, hash)

	return
}
