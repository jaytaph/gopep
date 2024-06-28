package main

import "github.com/sirupsen/logrus"

// / Here we store data
type StorageFacility struct {
	secret ScalarNonZero
	data   map[string][]ElGamal
}

func NewStorageFacility(ks *KeyServer) *StorageFacility {
	return &StorageFacility{
		*ks.GenerateFactor("SF"),
		make(map[string][]ElGamal),
	}
}

func (sf *StorageFacility) Store(pid ElGamal, data ElGamal) {
	logrus.Tracef("[SF] Store: %s %s", pid, data)

	key := decrypt(&pid, &sf.secret).HexString()
	logrus.Tracef("[SF] Decrypted key: %s", key)

	if _, ok := sf.data[key]; !ok {
		sf.data[key] = make([]ElGamal, 0)
	}

	sf.data[key] = append(sf.data[key], data)
}

func (sf *StorageFacility) Retrieve(ppid ElGamal) []ElGamal {
	logrus.Tracef("[SF] Retrieve: %s", ppid)

	key := decrypt(&ppid, &sf.secret).HexString()
	logrus.Tracef("[SF] Decrypted key: %s", key)

	if data, ok := sf.data[key]; ok {
		return data
	}

	return nil
}
