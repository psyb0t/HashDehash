package hasher

import (
	"HashDehash/db"
	"crypto/md5"
	"fmt"
)

func (h *Hasher) MD5() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/MD5/to_from")
		dbFromTo := db.NewDB("hash_db/MD5/from_to")

		h.hashes.MD5, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.MD5 = fmt.Sprintf("%x", md5.Sum([]byte(h.input)))

		dbToFrom.Set(h.hashes.MD5, h.input)
		dbFromTo.Set(h.input, h.hashes.MD5)
	}()
}
