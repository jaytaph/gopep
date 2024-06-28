package main

import "github.com/sirupsen/logrus"

// @todo:  figure out how the keyserver works in the global picture. We generate a factor by
//  using (ka_am * ka_t) * private_key. What do we do with this information and why do we store ka_am and ka_t in the
// pools?

// Key server holds all keys. The global key and the numbers for each factor. These factors are
// allowed outside the keyserver.
type KeyServer struct {
	Global *KeyPair
	PoolAm *FactorPool
	PoolT  *FactorPool
}

func NewKeyServer(poolAm *FactorPool, poolT *FactorPool) *KeyServer {
	return &KeyServer{
		NewKeyPair(),
		poolAm,
		poolT,
	}
}

func (ks *KeyServer) GetGlobalPublicKey() *GroupElement {
	return ks.Global.pk
}

func (ks *KeyServer) GenerateFactor(id string) *ScalarNonZero {
	logrus.Tracef("[KS] GenerateFactor: %s", id)

	// Generate Ka@am and Ka@t
	ka_am := RandomScalarNonZero()
	ka_t := RandomScalarNonZero()
	ka := ka_am.MulScalar(ka_t)

	ks.PoolAm.Set(id, *ka_am)
	ks.PoolT.Set(id, *ka_t)

	// The private key for "a"
	return ka.MulScalar(ks.Global.GetPrivateKey())
}
