package hasher

import (
	"HashDehash/db"
	"crypto/sha1"
	"fmt"
)

func (h *Hasher) SHA1() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA1/to_from")
		dbFromTo := db.NewDB("hash_db/SHA1/from_to")

		h.hashes.SHA1, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA1 = fmt.Sprintf("%x", sha1.Sum([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA1, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA1)
	}()
}
