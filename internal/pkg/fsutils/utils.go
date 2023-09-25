package fsutils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func HashFile(filePath string, rootPath string) (hash string, err error) {

	path := fmt.Sprintf("%s%c%s", rootPath, os.PathSeparator, filePath)
	file, err := os.Open(path)
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
