
package bitfield

type Bits struct{ words []uint64 }

func New(n int) *Bits { return &Bits{words: make([]uint64, (n+63)/64)} }

func (b *Bits) Set(bit int) { b.words[0] |= 1 << bit } // BUG

func (b *Bits) Get(bit int) bool { return b.words[0]&(1<<bit) != 0 }
