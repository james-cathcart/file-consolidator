package fsutils

import (
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"
)

func HashFile(filePath string) (hash string, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer func(fileToClose *os.File) {
		err := fileToClose.Close()
		if err != nil {
			log.Printf("error: could not close file %v", fileToClose)
		}
	}(file)

	fileInfo, err := file.Stat()
	if fileInfo.IsDir() {
		err = errors.New(`error: directory, skipping comparison`)
		return
	}

	hashWriter := sha256.New()
	_, err = io.Copy(hashWriter, file)
	if err != nil {
		return
	}

	hash = string(hashWriter.Sum(nil))

	return
}
