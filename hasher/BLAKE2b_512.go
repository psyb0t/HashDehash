package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/blake2b"
)

func (h *Hasher) BLAKE2b_512() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/BLAKE2b_512/to_from")
		dbFromTo := db.NewDB("hash_db/BLAKE2b_512/from_to")

		h.hashes.BLAKE2b_512, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.BLAKE2b_512 = fmt.Sprintf("%x",
			blake2b.Sum512([]byte(h.input)))

		dbToFrom.Set(h.hashes.BLAKE2b_512, h.input)
		dbFromTo.Set(h.input, h.hashes.BLAKE2b_512)
	}()
}
