package hasher

import "sync"

type Hashes struct {
	MD5         string
	MD4         string
	SHA1        string
	SHA224      string
	SHA256      string
	SHA384      string
	SHA512      string
	RIPEMD160   string
	SHA3_224    string
	SHA3_256    string
	SHA3_384    string
	SHA3_512    string
	SHA512_224  string
	SHA512_256  string
	BLAKE2s_256 string
	BLAKE2b_256 string
	BLAKE2b_384 string
	BLAKE2b_512 string
}

type Hasher struct {
	input  string
	hashes *Hashes
	wg     sync.WaitGroup
}

func NewHasher() *Hasher {
	return &Hasher{hashes: &Hashes{}}
}

func (h *Hasher) SetInput(input string) {
	h.input = input
}

func (h *Hasher) Hash() {
	h.MD5()
	h.MD4()
	h.SHA1()
	h.SHA224()
	h.SHA256()
	h.SHA384()
	h.SHA512()
	h.RIPEMD160()
	h.SHA3_224()
	h.SHA3_256()
	h.SHA3_384()
	h.SHA3_512()
	h.SHA512_224()
	h.SHA512_256()
	h.BLAKE2s_256()
	h.BLAKE2b_256()
	h.BLAKE2b_384()
	h.BLAKE2b_512()

	h.wg.Wait()
}

func (h *Hasher) GetHashes() *Hashes {
	return h.hashes
}
