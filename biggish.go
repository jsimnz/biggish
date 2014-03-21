package biggish

import (
	"errors"
	"math/big"

	"github.com/jsimnz/errgo"
)

const (
	// Generic int sizes
	SIZE32 = Size(32 << iota)
	SIZE64
	SIZE128
	SIZE256

	// Signed and unsigned numbers
	SIGNED = iota
	UNSIGNED
	NOSIGN

	// errors
	ErrTooBig = iota
)

func init() {
	errgo.Register(ErrTooBig, "Given int %v too large for %v sized []byte backend")
}

type Size int

type Int struct {
	// size of int (in bytes)
	size int
	// numeric sign
	sign int
	// byte array backend
	val []byte
}

// Create a new int with an initial value x, and backed by a n sized byte array
func NewInt(x int, size Size) (*Int, error) {
	i := &Int{
		val:  make([]byte, size),
		sign: SIGNED,
	}
	bigi := big.NewInt(int64(x))
	b := bigi.Bytes()
	if len(b) < size {
		i.val = b
		return i, nil
	} else {
		return nil, errgo.New(ErrTooBig, x, size)
	}
}

func (i *Int) SetSign(s sign) {
	i.sign = s
}

func (z *Int) Add(x, y *Int) *Int {
	bigz := z.getBig()
	bigx := x.getBig()
	bigy := y.getBig()

	bigz.Add(bigx, bigy)

}

func (z *Int) getBig() *big.Int {
	i := big.NewInt(int64(0))
	i.SetBytes(z.val)
	return i
}

func (i *Int) maxVal() *big.Int {
	bigi := newBig()
}

func (i *Int) catchOverflow() {
	if len(i.val) > size {
		if i.sign == SIGNED {

		}
	}
}

func getInt(x *big.Int) *Int {

}

func newBig() *big.Int {
	return big.NewInt(0)
}
