package passmanager

import (
	"fmt"
	"safe_key_store/hasher"

	"github.com/danieljoos/wincred"
)

const (
	masterpasskey = "master"
)

// verifies master password
func VerifyMasterPassword(pass string) error {
	cred, err := getCredential(masterpasskey)
	if err != nil {
		return err
	}


	masterpassh := hasher.FromBytes(cred.CredentialBlob)
	passh := hasher.CreateHash([]byte(pass))

	if !passh.Compare(masterpassh){
		return fmt.Errorf("invalid master password: %s",pass)
	}

	return nil
}


// this function sets the master password
func SetMasterPassword(pass string) error{
	passhash := hasher.CreateHash([]byte(pass))
	
	mastercred := wincred.NewGenericCredential(masterpasskey)
	mastercred.CredentialBlob = passhash.Bytes()

	if err := mastercred.Write(); err != nil{
		return err
	}
	return nil
}


