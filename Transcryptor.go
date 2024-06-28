package main

type Transcryptor struct {
	Pool *FactorPool
}

func NewTranscryptor(pool *FactorPool) *Transcryptor {
	return &Transcryptor{pool}
}

func (t *Transcryptor) GetFactor(id string) (Factor, error) {
	return t.Pool.Get(id)
}

// Rekeys and reshuffles data based on the keys found in the transcryptor pool
func (t *Transcryptor) Rks(id string, data *ElGamal) *ElGamal {
	f, err := t.GetFactor(id)
	if err != nil {
		return nil
	}

	return rks(data, &f.K, &f.S)
}

// Rekeys data based on the keys found in the transcryptor pool
func (t *Transcryptor) Rekey(id string, data *ElGamal) *ElGamal {
	f, err := t.GetFactor(id)
	if err != nil {
		return nil
	}

	return rekey(data, &f.K)
}
