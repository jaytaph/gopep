package main

import (
	"encoding/hex"

	"github.com/oasisprotocol/curve25519-voi/curve"
)

type ElGamal struct {
	B *GroupElement
	C *GroupElement
	Y *GroupElement
}

func (g ElGamal) Encode() []byte {
	b := make([]byte, 96)
	copy(b[0:32], g.B.Encode())
	copy(b[32:64], g.C.Encode())
	copy(b[64:96], g.Y.Encode())

	return b
}

func (g ElGamal) HexString() string {
	return hex.EncodeToString(g.Encode())
}

func encrypt(msg *GroupElement, public_key *GroupElement) *ElGamal {
	if public_key.IsIdentity() {
		panic("pubkey is identity")
	}

	r := RandomScalarNonZero()
	G := NewGroupElement(curve.RISTRETTO_BASEPOINT_POINT)

	return &ElGamal{
		B: r.Mul(G),
		C: msg.Add(r.Mul(public_key)),
		Y: public_key,
	}
}

func decrypt(s *ElGamal, secret_key *ScalarNonZero) *GroupElement {
	return s.C.Sub(secret_key.Mul(s.B))
}

func rerandomize(v *ElGamal, s *ScalarNonZero) *ElGamal {
	G := NewGroupElement(curve.RISTRETTO_BASEPOINT_POINT)

	return &ElGamal{
		B: s.Mul(G).Add(v.B),
		C: s.Mul(v.Y).Add(v.C),
		Y: v.Y,
	}
}

func rekey(v *ElGamal, k *ScalarNonZero) *ElGamal {
	return &ElGamal{
		B: k.Invert().Mul(v.B),
		C: v.C,
		Y: k.Mul(v.Y),
	}
}

func reshuffle(v *ElGamal, n *ScalarNonZero) *ElGamal {
	return &ElGamal{
		B: n.Mul(v.B),
		C: n.Mul(v.C),
		Y: v.Y,
	}
}

// Combined rekey(k) and reshuffle(n)
func rks(v *ElGamal, k *ScalarNonZero, n *ScalarNonZero) *ElGamal {
	return &ElGamal{
		B: n.MulScalar(k.Invert()).Mul(v.B),
		C: n.Mul(v.C),
		Y: k.Mul(v.Y),
	}
}
