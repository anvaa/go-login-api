package filefunc

import (
	
	// "io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func IsExists(path string) bool {
	if _, err := os.Stat(os.Getenv("WORKING_FOLDER") + path); os.IsNotExist(err) {
		return false
	}
	return true
}


// Create folder
func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

// Create file
func CreateFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}
}

// Delete file
func DeleteWebFile(webpath string) error {
	path := os.Getenv("WORKING_FOLDER") + webpath
	_, err := os.Stat(path); os.IsExist(err)
	if err != nil {
		return err
	}
	os.Remove(path)
	return nil
}

// Delete folder with content
func DeleteFolder(path string) {
	if _, err := os.Stat(path); os.IsExist(err) {
		os.RemoveAll(path)
	}
}

// GetFileList returns a list of files in a directory
func GetFileList(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// GetFileListByExt returns a list of files in a directory with a specific extension
func GetFileListByExt(dir string, ext string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// create image file from byte stream
