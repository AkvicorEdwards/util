package util

import (
	"testing"
)

func TestAES(t *testing.T) {
	testString := "Akvicor"
	testKey := "6dac2a12159afc4ee5100f6a950052a2"
	testIv := "xb27xzkb49s01byh"
	testEncrypted := "nPH37q+v2zGpXmdYKAE/JA=="
	enc := NewAES().EncryptCBC([]byte(testString), []byte(testKey), []byte(testIv))
	if !enc.Encrypted() {
		t.Errorf("expected true got false")
	}
	if enc.Base64().String() != testEncrypted {
		t.Errorf("expected %s got %s", testEncrypted, enc.Base64().String())
	}

	dec := enc.Decrypt([]byte(testKey), []byte(testIv))
	if dec.Encrypted() {
		t.Errorf("expected false got true")
	}
	if dec.String() != testString {
		t.Errorf("expected %s got %s", testString, dec.String())
	}

	dec = NewAES().DecryptCBC(enc.Bytes(), []byte(testKey), []byte(testIv))
	if dec.Encrypted() {
		t.Errorf("expected false got true")
	}
	if dec.String() != testString {
		t.Errorf("expected %s got %s", testString, dec.String())
	}
}
