package main

import (
	"dedupfs/internal/pkg/common"
	"dedupfs/internal/pkg/fsutils"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	fmt.Println("Starting app")

	uniqueFiles := make(map[string][]common.FileRecord)

	var dirs []string
	for _, v := range os.Args[1:] {
		dirs = append(dirs, v)
	}

	for _, dir := range dirs {

		fileSystem := os.DirFS(dir)

		_ = fs.WalkDir(
			fileSystem,
			".",
			func(path string, d fs.DirEntry, err error) error {

				if d == nil {
					log.Println("warn: DirEntry is nil")
				}

				if !d.IsDir() {

					filePath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, path)
					log.Printf("checking file: %s\n", filePath)

					var hash string
					hash, hashErr := fsutils.HashFile(filePath)
					if hashErr != nil {
						log.Println(hashErr)
					} else {
						if _, ok := uniqueFiles[hash]; !ok {
							uniqueFiles[hash] = []common.FileRecord{
								{
									FilePath: filePath,
									Hash:     hash,
								},
							}
						} else {
							uniqueFiles[hash] = append(uniqueFiles[hash], common.FileRecord{
								FilePath: filePath,
								Hash:     hash,
							})
						}
					}
				}

				return err
			})

	}

	fmt.Println(`Duplication Report:`)
	fmt.Println(`--------------------`)

	fmt.Printf("%d unique files were found, see duplication report below:\n", len(uniqueFiles))
	for hash, filesFound := range uniqueFiles {

		fmt.Printf("For Hash: %s, found:\n", hash)
		for _, fileFound := range filesFound {
			fmt.Printf("\t%s\n", fileFound.FilePath)
		}

	}

	fmt.Println("App finished...")
}
