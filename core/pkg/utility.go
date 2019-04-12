package adbt

import (
	"log"
	"os"
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
	exist, err := isPathExists(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
