package main

import (
	"log"
	"os"

	"github.com/oasisprotocol/curve25519-voi/curve"
	"github.com/sirupsen/logrus"
)

func main() {
	println("GoPP")

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetLevel(logrus.TraceLevel)
	log.SetOutput(os.Stdout)

	// test_encdec()

	poolAm := NewFactorPool()
	poolT := NewFactorPool()
	keyServer := NewKeyServer(poolAm, poolT)
	accessManager := NewAccessManager(poolAm)
	transcryptor := NewTranscryptor(poolT)
	storage := NewStorageFacility(keyServer)

	// Store the doctor's secret in the key server (this is a secret, does it make sense to store it on a keyserver? Why
	// does the secret leave the doctor?)
	doctor_secret := keyServer.GenerateFactor("DOC")

	smartwatch_senddata("950000024", keyServer, accessManager, transcryptor, storage)
	doctor_retrieval(doctor_secret, "950000024", keyServer, accessManager, transcryptor, storage)
	// doctor_retrieval(doctor_secret, "950000012", keyServer, accessManager, transcryptor, storage)
}

// This function is the scenario where the smartwatch sends a few pieces of data to the storage facility
func smartwatch_senddata(
	bsn string,
	keyServer *KeyServer,
	accessManager *AccessManager,
	transcryptor *Transcryptor,
	storage *StorageFacility,
) {
	println("This is work that is done inside our smartwatch that knows our BSN")

	// Create some random data that we want to send to the storage facility. This random data is an AES key which can be
	// used to encrypt large blobs of data.
	aes_key1 := RandomGroupElement()
	enc_data1 := encrypt(aes_key1, keyServer.Global.GetPublicKey())
	println("**************************************************************")
	println("AES key1 generated: ", aes_key1.HexString())
	println("**************************************************************")

	ppid_a := GeneratePseudonym(bsn, keyServer.Global.GetPublicKey())
	// println("Ppid_a: ", ppid_a.HexString())
	// We can rerandomize our polymorphic pseudonym here as many times as we like.
	ppid_a = RerandomizeLocal(ppid_a)
	// println("Ppid_a (rerandomized): ", ppid_a.HexString())
	// We must transform our ppid_a so it can be accepted by the storage facility as identified with "SF"
	ppid_a_sf := TransformForDestination(ppid_a, "SF", transcryptor, accessManager)
	// println("Ppid_a_sf: ", ppid_a_sf.HexString())
	storage.Store(*ppid_a_sf, *enc_data1)

	// Send more information (another aes key) to the SF. It should be stored in the same LEP as the first key.
	// SF cannot decrypt the data though, but it can decrypt the pseudonym
	aes_key2 := RandomGroupElement()
	enc_data2 := encrypt(aes_key2, keyServer.Global.GetPublicKey())
	println("**************************************************************")
	println("AES key2 generated: ", aes_key2.HexString())
	println("**************************************************************")

	ppid_a = RerandomizeLocal(ppid_a)
	ppid_a_sf = TransformForDestination(ppid_a, "SF", transcryptor, accessManager)
	storage.Store(*ppid_a_sf, *enc_data2)

	println("All done. The smartwatch has send 2 pieces of data to the storage facility for this BSN.")
}

func doctor_retrieval(
	doctor_secret *ScalarNonZero,
	bsn string,
	keyServer *KeyServer,
	accessManager *AccessManager,
	transcryptor *Transcryptor,
	storage *StorageFacility,
) {
	println("\n\nAt the doctors. The doctor wants to retrieve the data for a specific BSN")

	// Generate a pseudonym for this bsn and transform it so we can fetch info from the SF
	ppid_a := GeneratePseudonym(bsn, keyServer.Global.GetPublicKey())
	// println("Ppid_a: ", ppid_a.HexString())
	ppid_a_sf := TransformForDestination(ppid_a, "SF", transcryptor, accessManager)
	// println("Ppid_a_sf: ", ppid_a_sf.HexString())
	items := storage.Retrieve(*ppid_a_sf)

	println("Data retrieved from the storage for bsn: ", len(items))

	for _, item := range items {
		// With the "item" alone, we cannot decrypt the data. We need to ask the transcryptor AND the access manager to
		// help us. This is done by rekeying the data. Note that we rerandomize the data before we rekey it so the same
		// data is not sent to the transcryptor and access manager deterministically.
		data_doc := rerandomize(&item, RandomScalarNonZero())
		data_doc = transcryptor.Rekey("DOC", data_doc)
		data_doc = accessManager.Rekey("DOC", data_doc)

		// Now we can decrypt the data with the doctor's secret
		data := decrypt(data_doc, doctor_secret)
		logrus.Infof("Decrypted AES key: %s", data.HexString())
	}
}

func test_encdec() {
	G := NewGroupElement(curve.RISTRETTO_BASEPOINT_POINT)
	s := RandomScalarNonZero()
	p := s.Mul(G)

	value := RandomGroupElement()
	logrus.Debugf("V: %v\n", value.String())

	encrypted := encrypt(value, p)
	logrus.Debugf("E: %v\n", &encrypted)

	decrypted := decrypt(encrypted, s)
	logrus.Debugf("D: %v\n", decrypted.String())
}

func TransformForDestination(ppid *ElGamal, dest string, transcryptor *Transcryptor, accessManager *AccessManager) *ElGamal {
	ppid = RerandomizeLocal(ppid)

	ppid_dest := transcryptor.Rks(dest, ppid)
	ppid_dest = accessManager.Rks(dest, ppid_dest)

	// This is the data we need to send to the dest
	logrus.Debugf("! epid(A)@%s: %s", dest, ppid_dest.HexString())

	return ppid_dest
}
