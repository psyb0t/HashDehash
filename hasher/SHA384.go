package hasher

import (
	"HashDehash/db"
	"crypto/sha512"
	"fmt"
)

func (h *Hasher) SHA384() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/SHA384/to_from")
		dbFromTo := db.NewDB("hash_db/SHA384/from_to")

		h.hashes.SHA384, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.SHA384 = fmt.Sprintf("%x", sha512.Sum384([]byte(h.input)))

		dbToFrom.Set(h.hashes.SHA384, h.input)
		dbFromTo.Set(h.input, h.hashes.SHA384)
	}()
}
