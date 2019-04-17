package adbt

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func isPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createFolderIfNotExist(folderPath string) {
	execPath, _ := os.Executable()
	execPath = filepath.Dir(execPath)
	realPath := filepath.Join(execPath, folderPath)
	exist, err := isPathExists(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		absPath, _ := filepath.Abs(realPath)
		fmt.Println("Creating folder" + absPath)
		err = os.Mkdir(absPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createFileIfNotExist(filePath string) {
	execPath, _ := os.Executable()
	execPath = filepath.Dir(execPath)
	realPath := filepath.Join(execPath, filePath)
	exist, err := isPathExists(realPath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		f, err := os.Create(realPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}
