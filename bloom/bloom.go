package bloom

type Filter struct {
	bits []byte

	m uint64

	k uint64
}


func FromBytes(
	bits []byte,
) *Filter {

	return &Filter{
		bits: bits,
		m: uint64(len(bits) * 8),
		k: 7,
	}
}


func New(
	expectedItems uint64,
) *Filter {

	m := expectedItems * 10

	return &Filter{
		bits: make(
			[]byte,
			(m+7)/8,
		),
		m: m,
		k: 7,
	}
}

func (f *Filter) setBit(
	pos uint64,
) {

	bytePos := pos / 8

	bitPos := pos % 8

	f.bits[bytePos] |=
		(1 << bitPos)
}

func (f *Filter) testBit(
	pos uint64,
) bool {

	bytePos := pos / 8

	bitPos := pos % 8

	return (f.bits[bytePos] &
		(1 << bitPos)) != 0
}

func (f *Filter) Add(
	key string,
) {

	h1 := hash1(key)

	h2 := hash2(key)

	for i := uint64(0); i < f.k; i++ {

		pos :=
			(h1 + i*h2) % f.m

		f.setBit(pos)
	}
}

func (f *Filter) MayContain(
	key string,
) bool {

	h1 := hash1(key)

	h2 := hash2(key)

	for i := uint64(0); i < f.k; i++ {

		pos :=
			(h1 + i*h2) % f.m

		if !f.testBit(pos) {
			return false
		}
	}

	return true
}


func (f *Filter) Bits() []byte {
	return f.bits
}