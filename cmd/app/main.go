package main

import (
	"dedupfs/internal/app/deduper"
	"dedupfs/internal/app/filemigrator"
	"dedupfs/internal/app/filescanner"
	"dedupfs/internal/pkg/common"
	"fmt"
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
		if v == `-v` {
			continue
		}
		dirs = append(dirs, v)
	}

	verbose := false

	if os.Args[len(os.Args)-1] == `-v` {
		verbose = true
	}

	scan := filescanner.New(collect, &wg)
	scan.ScanDir(verbose, dirs...)

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

	deDupe := deduper.New()
	totalCount, duplicateCount, unique, duplicate := deDupe.DeDuplicate(uniqueFiles)

	report := common.ScanReport{
		UniqueFiles:     unique,
		DuplicateFiles:  duplicate,
		FilesFound:      uniqueFiles,
		TotalFiles:      totalCount,
		DuplicatesFound: duplicateCount,
		UniqueFileCount: totalCount - duplicateCount,
	}

	fmt.Println("\nDuplication Report:")
	fmt.Println("--------------------")
	fmt.Printf("-> Total Files Scanned: %d\n", report.TotalFiles)
	fmt.Printf("-> Unique Files Found: %d\n", report.UniqueFileCount)
	fmt.Printf("-> Duplicates Found: %d\n\n", report.DuplicatesFound)

	fmt.Println("Copy unique files to new directory? (yes/no)")
	var copyUnique string
	_, _ = fmt.Scanln(&copyUnique)
	if copyUnique == `yes` {
		fmt.Println("Destination directory (must not exist):")
		var destination string
		_, _ = fmt.Scanln(&destination)
		fmt.Printf("Copying unique files to %s\n", destination)

		mig := filemigrator.New()
		err := mig.MigrateUniqueFiles(destination, dirs, report.UniqueFiles)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("App finished...")
}
