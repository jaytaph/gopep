package main

import (
	"encoding/hex"
	"fmt"

	"github.com/oasisprotocol/curve25519-voi/curve"
)

// This is not much more than a wrapper around curve.RistrettoPoint
type GroupElement struct {
	point *curve.RistrettoPoint
}

func NewGroupElementFromBytes(b []byte) *GroupElement {
	var p curve.RistrettoPoint

	_, err := p.SetUniformBytes(b)
	if err != nil {
		panic(err)
	}

	return &GroupElement{
		point: &p,
	}

}

func NewGroupElement(p *curve.RistrettoPoint) *GroupElement {
	return &GroupElement{
		point: p,
	}
}

func RandomGroupElement() *GroupElement {
	var p curve.RistrettoPoint
	if _, err := p.SetRandom(nil); err != nil {
		panic(err)
	}

	return &GroupElement{
		point: &p,
	}
}

func (g *GroupElement) GetPoint() *curve.RistrettoPoint {
	return g.point
}

func (g *GroupElement) Add(other *GroupElement) *GroupElement {
	var s curve.RistrettoPoint
	p := s.Add(g.point, other.point)

	return &GroupElement{
		point: p,
	}
}

func (g *GroupElement) IsIdentity() bool {
	return g.point.IsIdentity()
}

func (g *GroupElement) Sub(other *GroupElement) *GroupElement {
	var s curve.RistrettoPoint
	p := s.Sub(g.point, other.point)

	return &GroupElement{
		point: p,
	}
}

func (g *GroupElement) String() string {
	b, err := g.point.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", b)
}

func (g *GroupElement) HexString() string {
	b, err := g.point.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

func (g *GroupElement) Encode() []byte {
	b, err := g.point.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return b
}
