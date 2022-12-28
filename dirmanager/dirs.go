package dirmanager

import "os"

// creates directory at the given path
// and gives it HIDDEN attribute
func CreateHiddenDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return hide(path)
}


// returns names of all files and subdirs on the given path
func ListFiles(dirpath string) ([]string, error){
	entries, err := os.ReadDir(dirpath)
	if err != nil{
		return nil, err
	}

	files := []string{}
	for _, ent := range entries{
		files = append(files,ent.Name())
	}

	return files, nil
}


// removes a directory if it exists
func RemoveDir(path string) error{
	return removeOnPath(path)
}
