package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "a27cb3fb765615517b6bdf2af862d8e1c9416102114aa7b16cd0ddcdcaa2e408"
		addressStr = "2854bd145bce5018ebeb61505981beb9d3a911ee"
		privKey    = NewPrivateKeyFromString(seed)
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}

func TestPrivKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)

	// Test with valid msg
	assert.True(t, sig.Verify(pubKey, msg))

	// Test with invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// Test with invalid public key
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
