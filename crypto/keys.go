package crypto

import "crypto/ed25519"

const (
	privKeyLen = 64
	pubKeyLen  = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, pubKeyLen)
	copy(b, p.key[32:])

	return &PublicKey{
		key: b,
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}
