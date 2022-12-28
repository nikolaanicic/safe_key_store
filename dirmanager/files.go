package dirmanager

import (
	"os"
)

// creates a files on the given path and gives in the hidden attribute
func createHiddenFile(path string) (*os.File,error) {
	f, err := os.Create(path)

	if err != nil{
		return nil, err
	}

	if err := hide(path); err != nil{
		f.Close()
		return nil, err
	}

	return f, err
}


func ReadFile(path string) ([]byte, error){
	return os.ReadFile(path)
}

// creates a hidden file and writes the data to it	
func WriteFile(path string, data []byte) error{
	f, err := createHiddenFile(path)
	
	if err != nil{
		return err
	}
	
	defer f.Close()

	if _, err := f.Write(data); err != nil{
		return err
	}
	return nil
}


// removes a file on the path
func RemoveFile(path string) error{
	return removeOnPath(path)
}


// edits a file by deleting and then recreating the same file with 
// the new data
func EditFile(path string, newdata []byte) error{
	if err := removeOnPath(path); err != nil{
		return err
	}

	if err := WriteFile(path,newdata); err != nil{
		return err
	}

	return nil
}

