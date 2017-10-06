package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/blake2b"
)

func (h *Hasher) BLAKE2b_384() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/BLAKE2b_384/to_from")
		dbFromTo := db.NewDB("hash_db/BLAKE2b_384/from_to")

		h.hashes.BLAKE2b_384, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.BLAKE2b_384 = fmt.Sprintf("%x",
			blake2b.Sum384([]byte(h.input)))

		dbToFrom.Set(h.hashes.BLAKE2b_384, h.input)
		dbFromTo.Set(h.input, h.hashes.BLAKE2b_384)
	}()
}
