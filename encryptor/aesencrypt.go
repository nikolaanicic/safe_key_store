package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"safe_key_store/hasher"
)

const (
	aeskeylen = 32
	bit_128_bytelen = 16
)


func encryptAES(key []byte, data []byte) ([]byte, error){

	paddedkey := checkFixKeyAES(key)

	c, err := aes.NewCipher(paddedkey[:])

	if err != nil{
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil{
		return nil, err
	}

	nonce := make([]byte,gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil{
		return nil, err
	}
	
	return gcm.Seal(nonce,nonce,data,nil), nil
}


func decryptAES(key []byte, data []byte) ([]byte, error){

	paddedkey := checkFixKeyAES(key)

	c, err := aes.NewCipher(paddedkey[:])
	
	if err != nil{
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil{
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize{
		return nil, fmt.Errorf("invalid nonce size")
	}

	nonce, cipherbytes := data[:nonceSize],data[nonceSize:]
	plain, err := gcm.Open(nil,nonce,cipherbytes,nil)

	if err != nil{
		return nil, err
	}

	return plain, nil
}

// uses hasher to hash passed key and creates aes 32 byte key for later use
func checkFixKeyAES(key []byte) [aeskeylen]byte{
	newkey := [aeskeylen]byte{}

	keyhash := hasher.CreateHash(key)
	copy(newkey[:],keyhash[0:aeskeylen])

	return newkey
}