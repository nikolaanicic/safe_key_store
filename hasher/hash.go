package hasher

import (
	"crypto/sha256"
)

const hashlen = 32
type Hash [hashlen]byte


func (h Hash) String() string{
	return string(h[:])
}


func FromBytes(b []byte) Hash{
	var hash Hash
	copy(hash[:],b)
	return hash
}

func (h Hash) Bytes() []byte{
	return h[:]
}

// creates [32]byte hash from passed data
func CreateHash(data []byte) Hash {
	return hash(data)
}

// salting function used in hashing
func getSalt(data []byte) []byte {
	res := sha256.Sum256(data)
	return res[:]
}


func (h Hash) Compare(h2 Hash) bool{
	for i:=0;i<hashlen;i++{
		if h[i] != h2[i]{
			return false
		}
	}
	return true
}


// hashing implementation that should not be exposed
// adds the salt to the data and then computes the hash
// salt is just the data hashed with sha256
func hash(data []byte) Hash {
	return sha256.Sum256(append(data,getSalt(data)...))
}