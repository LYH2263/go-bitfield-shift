
package bitfield

type Bits struct{ words []uint64 }

func New(n int) *Bits { return &Bits{words: make([]uint64, (n+63)/64)} }

func (b *Bits) Set(bit int) { b.words[bit/64] |= 1 << uint(bit%64) }

func (b *Bits) Get(bit int) bool { return b.words[bit/64]&(1<<uint(bit%64)) != 0 }
