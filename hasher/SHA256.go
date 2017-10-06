package hasher

import (
	"HashDehash/db"
	"crypto/sha256"
	"fmt"
)

func (h *Hasher) SHA256() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA256/to_from")
		dbFromTo := db.NewDB("hash_db/SHA256/from_to")

		h.hashes.SHA256, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA256 = fmt.Sprintf("%x", sha256.Sum256([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA256, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA256)
	}()
}
