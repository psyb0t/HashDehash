package hasher

import (
	"HashDehash/db"
	"crypto/sha256"
	"fmt"
)

func (h *Hasher) SHA224() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA224/to_from")
		dbFromTo := db.NewDB("hash_db/SHA224/from_to")

		h.hashes.SHA224, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA224 = fmt.Sprintf("%x", sha256.Sum224([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA224, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA224)
	}()
}
