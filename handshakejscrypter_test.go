package handshakejscrypter_test

import (
	"../handshakejscrypter"
	"log"
	"testing"
)

const (
	ORIGINAL_PLAINTEXT = "some really long plaintext"
)

func TestEncryptionAndDecryption(t *testing.T) {
	handshakejscrypter.Setup("somesecretsaltofaspecificlength.") //32 bytes

	cipher := handshakejscrypter.Encrypt(ORIGINAL_PLAINTEXT)
	log.Println(cipher)

	plaintext := handshakejscrypter.Decrypt(cipher)
	log.Println(plaintext)

	if plaintext != ORIGINAL_PLAINTEXT {
		t.Errorf("Incorrect decrypted plaintext: " + plaintext)
	}
}
