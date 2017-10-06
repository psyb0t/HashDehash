package hasher

import (
	"HashDehash/db"
	"crypto/sha512"
	"fmt"
)

func (h *Hasher) SHA512() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA512/to_from")
		dbFromTo := db.NewDB("hash_db/SHA512/from_to")

		h.hashes.SHA512, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA512 = fmt.Sprintf("%x", sha512.Sum512([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA512, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA512)
	}()
}
