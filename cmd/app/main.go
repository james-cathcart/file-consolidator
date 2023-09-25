package main

import (
	"dedupfs/internal/pkg/common"
	"dedupfs/internal/pkg/fsutils"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
)

func main() {

	fmt.Println("Starting app")

	uniqueFiles := make(map[string][]common.FileRecord)
	collect := make(chan common.FileRecord)
	wg := sync.WaitGroup{}

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

				if !d.IsDir() {

					filePath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, path)

					go func() {
						wg.Add(1)
						defer wg.Done()
						log.Printf("checking file: %s\n", filePath)
						hash, hashErr := fsutils.HashFile(filePath)
						if hashErr != nil {
							log.Println(hashErr)
						} else {
							collect <- common.FileRecord{
								FilePath: filePath,
								Hash:     hash,
							}
						}

					}()
				}

				return err
			})

	}

	go func() {
		wg.Wait()
		close(collect)
	}()

	for value := range collect {
		if _, ok := uniqueFiles[value.Hash]; !ok {
			uniqueFiles[value.Hash] = []common.FileRecord{
				{
					FilePath: value.FilePath,
					Hash:     value.Hash,
				},
			}
		} else {
			uniqueFiles[value.Hash] = append(uniqueFiles[value.Hash], common.FileRecord{
				FilePath: value.FilePath,
				Hash:     value.Hash,
			})
		}
	}

	fmt.Println("\nDuplication Report:")
	fmt.Println(`--------------------`)

	fmt.Printf("%d unique files were found, see duplication report below:\n\n", len(uniqueFiles))
	for hash, filesFound := range uniqueFiles {

		fmt.Printf("For Hash: %s, found:\n", hash)
		for _, fileFound := range filesFound {
			fmt.Printf("\t%s\n", fileFound.FilePath)
		}

	}

	fmt.Println("App finished...")
}
