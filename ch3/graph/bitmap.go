package graph

type Bitmap []uint32

func (Bitmap) Init(len uint) Bitmap {
	return make([]uint32, (len+31)>>5)
}

func (b *Bitmap) Set(pos uint, val bool) {
	if val {
		(*b)[pos>>5] |= 1 << (pos & 31)
	} else {
		(*b)[pos>>5] &^= 1 << (pos & 31)
	}
}

func (b *Bitmap) Get(pos uint) bool {
	return ((*b)[pos>>5] & (1 << (pos & 31))) != 0
}
