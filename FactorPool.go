package main

import (
	"errors"
)

// We need both a K and S to rekey/shuffle. However, we have split these numbers into two parts. One part is stored
// in the transcryptor, the other part is stored in the access manager. This is done to prevent the transcryptor
// from being able to decrypt the data.
// So we have a factorpool for the access_manager, and a factorpool for the transcryptor. To get the full K and S
// we need to combine the two factors:
//
//	  K = k_am * k_t
//	  S = s_am * s_t
//
//	Only when we rekey/shuffle (rks) with these two factors, we can decrypt the data with the private key of
//
// that destination.
type Factor struct {
	K ScalarNonZero
	S ScalarNonZero
}

// A factor pool holds a list of factors (K,S numbers) for each destination. Note that K and S are not the complete
// numbers, but merely a part of a whole number (see above).
type FactorPool struct {
	Factors map[string]Factor
}

func NewFactorPool() *FactorPool {
	return &FactorPool{make(map[string]Factor)}
}

func (fp *FactorPool) Get(id string) (Factor, error) {
	// logrus.Tracef("[FP] Get: %s", id)

	factor, ok := fp.Factors[id]
	if !ok {
		return Factor{}, errors.New("Factor not found")
	}
	return factor, nil
}

func (fp *FactorPool) Set(id string, k ScalarNonZero) {
	// logrus.Tracef("[FP] Set: %s: %s", id, k)
	fp.Factors[id] = Factor{k, *RandomScalarNonZero()}
}
