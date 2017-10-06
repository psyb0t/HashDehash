package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/md4"
)

func (h *Hasher) MD4() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/MD4/to_from")
		dbFromTo := db.NewDB("hash_db/MD4/from_to")

		h.hashes.MD4, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		md4hasher := md4.New()
		md4hasher.Write([]byte(h.input))

		h.hashes.MD4 = fmt.Sprintf("%x", md4hasher.Sum(nil))

		dbToFrom.Set(h.hashes.MD4, h.input)
		dbFromTo.Set(h.input, h.hashes.MD4)
	}()
}
