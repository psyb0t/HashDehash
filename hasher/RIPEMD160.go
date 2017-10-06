package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func (h *Hasher) RIPEMD160() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/RIPEMD160/to_from")
		dbFromTo := db.NewDB("hash_db/RIPEMD160/from_to")

		h.hashes.RIPEMD160, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		hasher := ripemd160.New()
		hasher.Write([]byte(h.input))

		h.hashes.RIPEMD160 = fmt.Sprintf("%x", hasher.Sum(nil))

		dbToFrom.Set(h.hashes.RIPEMD160, h.input)
		dbFromTo.Set(h.input, h.hashes.RIPEMD160)
	}()
}
