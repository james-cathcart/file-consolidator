package fsutils

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func HashFile(filePath string) (hash string, err error) {

	//log.Printf("hashing file: %s\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer func(fileToClose *os.File) {
		err := fileToClose.Close()
		if err != nil {
			log.Printf("HashFile: could not close file %v", fileToClose)
		}
	}(file)

	hashWriter := sha256.New()
	_, err = io.Copy(hashWriter, file)
	if err != nil {
		return
	}

	hash = string(hashWriter.Sum(nil))

	return
}
