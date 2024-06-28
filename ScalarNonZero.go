package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/oasisprotocol/curve25519-voi/curve"
	"github.com/oasisprotocol/curve25519-voi/curve/scalar"
)

// ScalarNonZero is a wrapper around curve.Scalar that ensures that the scalar is non-zero.
type ScalarNonZero struct {
	scalar *scalar.Scalar
}

func RandomScalarNonZero() *ScalarNonZero {
	var sc scalar.Scalar
	s, err := sc.SetRandom(rand.Reader)
	if err != nil {
		panic(err)
	}

	return &ScalarNonZero{
		scalar: s,
	}
}

func (s *ScalarNonZero) One() *ScalarNonZero {
	var sc scalar.Scalar

	return &ScalarNonZero{
		scalar: sc.One(),
	}
}

func (s *ScalarNonZero) Invert() *ScalarNonZero {
	var sc scalar.Scalar

	return &ScalarNonZero{
		scalar: sc.Invert(s.scalar),
	}
}

func (s *ScalarNonZero) MulScalar(t *ScalarNonZero) *ScalarNonZero {
	var sc scalar.Scalar

	return &ScalarNonZero{
		scalar: sc.Mul(s.scalar, t.scalar),
	}
}

func (s *ScalarNonZero) Mul(g *GroupElement) *GroupElement {
	var rp curve.RistrettoPoint
	p := g.GetPoint()

	return &GroupElement{
		point: rp.Mul(p, s.scalar),
	}
}

func (s *ScalarNonZero) String() string {
	b, err := s.scalar.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", b)
}

func (s *ScalarNonZero) HexString() string {
	b, err := s.scalar.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}
