package main

import (
	"dedupfs/internal/pkg/fsutils"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	fmt.Println("Starting app")

	var dirs []string
	for _, v := range os.Args[1:] {
		dirs = append(dirs, v)
	}

	for _, dir := range dirs {

		fileSystem := os.DirFS(dir)

		err := fs.WalkDir(
			fileSystem,
			".",
			func(path string, d fs.DirEntry, err error) error {

				log.Printf("walking path: %s", path)

				if d == nil {
					log.Println("warn: DirEntry is nil")
				} else {
					log.Printf("DirEntry name: %s", d.Name())
				}

				if !d.IsDir() {
					log.Printf("hashing file: %s", path)

					var hash string
					hash, hashErr := fsutils.HashFile(path, dir)
					if hashErr != nil {
						log.Println(hashErr)
					}

					fmt.Printf("file: %s, hash: %s\n", path, hash)
				}

				return err
			})

		if !errors.Is(err, fs.SkipDir) {
			log.Println(err)
		}
	}

	fmt.Println("App finished...")
}

//
//func MyWalkFunc(path string, d fs.DirEntry, err error) error {
//
//	return err
//}
//
//func hashFile(
//	process chan string,
//	collect chan common.UniqueFile,
//	errs chan error,
//	wg *sync.WaitGroup,
//) {
//
//	for file := range process {
//		file := file // intermediate variable for go routine
//		go func(filePath string) {
//
//			hash, err := fsutils.HashFile(file)
//			if err != nil {
//				errs <- err
//				return
//			}
//			collect <- common.UniqueFile{
//				FilePath: file,
//				Hash:     hash,
//			}
//		}(file)
//	}
//}
//
//func processResult(collect chan common.UniqueFile, uniqueFiles map[string]common.UniqueFile, wg *sync.WaitGroup) {
//	for file := range collect {
//		if _, ok := uniqueFiles[file.Hash]; !ok {
//			uniqueFiles[file.Hash] = file
//		}
//	}
//}
//
//func processError(errs chan error, reportedErrors []error, wg *sync.WaitGroup) {
//	for err := range errs {
//		reportedErrors = append(reportedErrors, err)
//	}
//}
