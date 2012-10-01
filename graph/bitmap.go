package graph

type Bitmap struct {
	table []uint32
}

func NewBitmap(len uint) (*Bitmap) {
	return &Bitmap{make([]uint32, (len + 31)>>5)}
}

func (b *Bitmap) Set(pos uint, val bool) {
	if val {
		b.table [pos >> 5] |= 1<<(pos & 31)
	} else {
		b.table[pos >> 5] &^= 1<<(pos & 31)
	}
}

func (b *Bitmap) Get(pos uint) bool {
	return (b.table[pos>>5] & (1<<(pos & 31))) != 0
}
