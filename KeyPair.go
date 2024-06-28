package main

import "github.com/oasisprotocol/curve25519-voi/curve"

type KeyPair struct {
	pk *GroupElement
	sk *ScalarNonZero
}

func NewKeyPair() *KeyPair {
	G := NewGroupElement(curve.RISTRETTO_BASEPOINT_POINT)

	y := RandomScalarNonZero()
	gy := y.Mul(G)

	return &KeyPair{gy, y}
}

func (kp *KeyPair) GetPublicKey() *GroupElement {
	return kp.pk
}

func (kp *KeyPair) GetPrivateKey() *ScalarNonZero {
	return kp.sk
}
