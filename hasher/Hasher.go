package hasher

type Hasher struct {
	input string
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) SetInput(input string) {
	h.input = input
}

func (h *Hasher) AllHashes() (hashes map[string]string) {
	hashes = make(map[string]string)

	hashes["MD5"] = h.MD5()
	hashes["MD4"] = h.MD4()
	hashes["SHA1"] = h.SHA1()
	hashes["SHA224"] = h.SHA224()
	hashes["SHA256"] = h.SHA256()
	hashes["SHA384"] = h.SHA384()
	hashes["SHA512"] = h.SHA512()
	hashes["RIPEMD160"] = h.RIPEMD160()
	hashes["SHA3-224"] = h.SHA3_224()
	hashes["SHA3-256"] = h.SHA3_256()
	hashes["SHA3-384"] = h.SHA3_384()
	hashes["SHA3-512"] = h.SHA3_512()
	hashes["SHA512-224"] = h.SHA512_224()
	hashes["SHA512-256"] = h.SHA512_256()
	hashes["BLAKE2s-256"] = h.BLAKE2s_256()
	hashes["BLAKE2b-256"] = h.BLAKE2b_256()
	hashes["BLAKE2b-384"] = h.BLAKE2b_384()
	hashes["BLAKE2b-512"] = h.BLAKE2b_512()

	return
}
