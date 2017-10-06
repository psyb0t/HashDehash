package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func (h *Hasher) SHA3_384() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA3_384/to_from")
		dbFromTo := db.NewDB("hash_db/SHA3_384/from_to")

		h.hashes.SHA3_384, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA3_384 = fmt.Sprintf("%x", sha3.Sum384([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA3_384, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA3_384)
	}()
}
