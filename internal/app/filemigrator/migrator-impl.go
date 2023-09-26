package filemigrator

import (
	"dedupfs/internal/pkg/common"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type SimpleImpl struct{}

func New() Migrator {
	return &SimpleImpl{}
}

func (mig *SimpleImpl) MigrateUniqueFiles(destination string, scannedDirs []string, files map[string]common.FileRecord) (err error) {

	_, err = os.Stat(destination)
	if !os.IsNotExist(err) {
		fmt.Printf("directory must not exist, status: provided directory exists")
		return err
	}

	fmt.Printf("Creating directory: %s\n", destination)
	err = os.MkdirAll(destination, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, file := range files {

		fmt.Printf("migrating file: %s\n", file.FilePath)

		sourceFile, err := os.Open(file.FilePath)
		if err != nil {
			log.Println(err)
			continue
		}
		relativePath := mig.stripPrefixes(sourceFile.Name(), scannedDirs)
		newFilePath := fmt.Sprintf(
			"%s%s",
			destination,
			relativePath,
		)
		pathTokens := strings.Split(relativePath, string(os.PathSeparator))
		if len(pathTokens) > 1 {
			newDirPathRel := strings.Join(pathTokens[0:len(pathTokens)-1], string(os.PathSeparator))
			newDirPath := fmt.Sprintf("%s%c%s", destination, os.PathSeparator, newDirPathRel)
			err = os.MkdirAll(newDirPath, 0777)
			if err != nil {
				log.Println(err)
				continue
			}
		}
		newFile, err := os.Create(newFilePath)
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = io.Copy(newFile, sourceFile)
		if err != nil {
			log.Println(err)
			continue
		}
		err = sourceFile.Close()
		if err != nil {
			log.Println(err)
			continue
		}
		err = newFile.Close()
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return
}

func (mig *SimpleImpl) stripPrefixes(originalString string, prefixes []string) string {
	strippedString := originalString
	for _, prefix := range prefixes {
		if strings.HasPrefix(originalString, prefix) {
			strippedString = strings.TrimPrefix(originalString, prefix)
			break
		}
	}
	return strippedString
}
