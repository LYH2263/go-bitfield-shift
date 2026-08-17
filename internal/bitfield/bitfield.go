
package bitfield

type Bits struct{ words []uint64 }

func New(n int) *Bits { return &Bits{words: make([]uint64, (n+63)/64)} }

func (b *Bits) Set(bit int) {
	w, off := bit/64, bit%64
	b.words[w] |= 1 << off
}

func (b *Bits) Get(bit int) bool {
	w, off := bit/64, bit%64
	return b.words[w]&(1<<off) != 0
}
