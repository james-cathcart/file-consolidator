package main

import (
	"dedupfs/internal/pkg/common"
	"dedupfs/internal/pkg/fsutils"
	"dedupfs/internal/pkg/runner"
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Starting app")

	collect := make(chan common.UniqueFile)
	uniqueFiles := make(map[string]common.UniqueFile)

	wg := sync.WaitGroup{}
	var deDup runner.Runner
	deDup = runner.NewDeDuplicator(
		uniqueFiles,
		collect,
		&wg,
	)

	deDup.SearchForDuplicates(common.SearchConfig{
		SearchDirs: []string{
			`data`,
		},
	})

	go func() {
		wg.Wait()
		close(collect)
	}()

	fmt.Println(`Printing Unique Files`)
	for file := range collect {
		if _, ok := uniqueFiles[file.Hash]; !ok {
			uniqueFiles[file.Hash] = file
		}
	}

	for _, v := range uniqueFiles {
		fmt.Printf("path: %s\n", v.FilePath)
	}

	fmt.Println("App finished...")
}

func hashFile(
	process chan string,
	collect chan common.UniqueFile,
	errs chan error,
	wg *sync.WaitGroup,
) {

	for file := range process {
		file := file // intermediate variable for go routine
		go func(filePath string) {

			hash, err := fsutils.HashFile(file)
			if err != nil {
				errs <- err
				return
			}
			collect <- common.UniqueFile{
				FilePath: file,
				Hash:     hash,
			}
		}(file)
	}
}

func processResult(collect chan common.UniqueFile, uniqueFiles map[string]common.UniqueFile, wg *sync.WaitGroup) {
	for file := range collect {
		if _, ok := uniqueFiles[file.Hash]; !ok {
			uniqueFiles[file.Hash] = file
		}
	}
}

func processError(errs chan error, reportedErrors []error, wg *sync.WaitGroup) {
	for err := range errs {
		reportedErrors = append(reportedErrors, err)
	}
}
