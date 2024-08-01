package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/McFlanky/blocker/crypto"
	"github.com/McFlanky/blocker/proto"
	"github.com/cbergoon/merkletree"

	pb "google.golang.org/protobuf/proto"
)

type TxHash struct {
	hash []byte
}

func NewTxHash(hash []byte) TxHash {
	return TxHash{
		hash: hash,
	}
}

func (h TxHash) CalculateHash() ([]byte, error) {
	return h.hash, nil
}

func (h TxHash) Equals(other merkletree.Content) (bool, error) {
	equals := bytes.Equal(h.hash, other.(TxHash).hash)
	return equals, nil
}

func VerifyBlock(b *proto.Block) bool {
	if len(b.Transactions) == 0 {
		if !VerifyRootHash(b) {
			fmt.Println("!!--- INVALID root hash")
			return false
		}
	}
	if len(b.PublicKey) != crypto.PubKeyLen {
		fmt.Println("!!--- INVALID public key length")
		return false
	}
	if len(b.Signature) != crypto.SignatureLen {
		fmt.Println("!!--- INVALID signature key length")
		return false
	}
	var (
		sig    = crypto.SignatureFromBytes(b.Signature)
		pubKey = crypto.PublicKeyFromBytes(b.PublicKey)
		hash   = HashBlock(b)
	)
	if !sig.Verify(pubKey, hash) {
		fmt.Printf("root: %+v\n", hex.EncodeToString(b.Header.RootHash))
		fmt.Println("!!--- INVALID block signature")
		return false
	}
	return true
}

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	if len(b.Transactions) > 0 {
		tree, err := GetMerkleTree(b)
		if err != nil {
			panic(err)
		}
		b.Header.RootHash = tree.MerkleRoot()
	}

	hash := HashBlock(b)
	sig := pk.Sign(hash)
	b.PublicKey = pk.Public().Bytes()
	b.Signature = sig.Bytes()

	return sig
}

func VerifyRootHash(b *proto.Block) bool {
	tree, err := GetMerkleTree(b)
	if err != nil {
		return false
	}
	valid, err := tree.VerifyTree()
	if err != nil {
		return false
	}

	if !valid {
		return false
	}

	return bytes.Equal(b.Header.RootHash, tree.MerkleRoot())
}

func GetMerkleTree(b *proto.Block) (*merkletree.MerkleTree, error) {
	list := make([]merkletree.Content, len(b.Transactions))
	for i := 0; i < len(b.Transactions); i++ {
		list[i] = NewTxHash(HashTransaction(b.Transactions[i]))
	}
	// Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// HashBlock returns a SHA256 of the header
func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
