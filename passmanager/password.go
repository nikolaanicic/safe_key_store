package passmanager

import (
	"fmt"
	"safe_key_store/encryptor"
	"safe_key_store/hasher"

	"github.com/danieljoos/wincred"
)

// hashes the key so it can't be read in the credential manager
func preparekey(key string) string{
	return hasher.CreateHash([]byte(key)).String()
}

// edits existing credential in the store
func EditCredential(username string, newpassword string) error{
	// gets the credential to check if it exists
	cred, err := getCredential(username)
	if err != nil{
		return fmt.Errorf("credential with username %s doesn't exist",username)
	}

	key := preparekey(username)
	data, err := encryptor.Encrypt(key,newpassword)
	
	if err != nil{
		return err
	}


	copy(cred.CredentialBlob,data)
	return cred.Write()
}

// stores the credential in the windows credential manager store
func StoreNewCredential(username string, password string) error{

	// gets the credential to check if one already exists
	cred, _ := getCredential(username)

	// if the credential doesn't exist it creates it
	// it will overwrite any existing credential with the username passed to the function
	if cred == nil{
		cred = wincred.NewGenericCredential(username)
	}

	data, err := encryptor.Encrypt(username, password)

	if err != nil{
		return err
	}
	
	cred.CredentialBlob = data
	return cred.Write()
}

// gets the credential from windows credential manager
func getCredential(target string) (*wincred.GenericCredential,error){
	cred, err := wincred.GetGenericCredential(target)

	if err != nil{
		return nil, err
	}

	return cred, nil
}


// returns the password encrypted in the credential with the passed target
func GetPassword(username string) (string, error){
	cred, err := getCredential(username)
	if err != nil{
		return "",err
	}

	pass, err := encryptor.Decrypt([]byte(username),cred.CredentialBlob)
	if err != nil{
		return "", err
	}
	return string(pass), nil
}

// takes data from windows credential manager and converts it to a string
// func credentialToString(cred *wincred.Credential) (string, error){
// 	data, err := encryptor.Decrypt([]byte(cred.TargetName),cred.CredentialBlob)
// 	if err != nil{
// 		return "", err
// 	}

// 	return string(data), nil
// }
