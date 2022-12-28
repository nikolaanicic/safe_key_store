package dirmanager

import (
	"os"
	"syscall"
)

// gives directories or files a HIDDEN attribute
func hide(path string) error {
	fptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	return syscall.SetFileAttributes(fptr, syscall.FILE_ATTRIBUTE_HIDDEN)
}

// removes everything found on the path
// either a file or a directory
func removeOnPath(path string) error{
	return os.RemoveAll(path)
}