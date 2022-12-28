package encryptor

func Encrypt(key string, data string) ([]byte, error) {
	keyb := []byte(key)
	datab := []byte(data)
	return encryptAES(keyb, datab)
}

func Decrypt(key []byte, data []byte) ([]byte, error) {
	return decryptAES(key, data)
}
