package hasher

import (
	"HashDehash/db"
	"crypto/sha512"
	"fmt"
)

func (h *Hasher) SHA512_224() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA512_224/to_from")
		dbFromTo := db.NewDB("hash_db/SHA512_224/from_to")

		h.hashes.SHA512_224, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA512_224 = fmt.Sprintf("%x", sha512.Sum512_224([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA512_224, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA512_224)
	}()
}
