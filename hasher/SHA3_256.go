package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func (h *Hasher) SHA3_256() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA3_256/to_from")
		dbFromTo := db.NewDB("hash_db/SHA3_256/from_to")

		h.hashes.SHA3_256, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA3_256 = fmt.Sprintf("%x", sha3.Sum256([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA3_256, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA3_256)
	}()
}
