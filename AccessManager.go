package main

// Just like the transcryptor, the access manager is needed in order to rekey data in order to decrypt it. The access
// manager however can also decide if the data is actually allowed to be decrypted. For proof-of-concept purposes, it
// will always allow access.

type AccessManager struct {
	Pool *FactorPool
}

func NewAccessManager(pool *FactorPool) *AccessManager {
	return &AccessManager{pool}
}

func (am *AccessManager) GetFactor(id string) (Factor, error) {
	// logrus.Tracef("[AM] GetFactor: %s", id)
	return am.Pool.Get(id)
}

func (am *AccessManager) Rks(id string, data *ElGamal) *ElGamal {
	f, err := am.GetFactor(id)
	if err != nil {
		return nil
	}

	return rks(data, &f.K, &f.S)
}

func (am *AccessManager) Rekey(id string, data *ElGamal) *ElGamal {
	f, err := am.GetFactor(id)
	if err != nil {
		return nil
	}

	return rekey(data, &f.K)
}
