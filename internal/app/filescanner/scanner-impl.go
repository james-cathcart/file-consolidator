package filescanner

import (
	"dedupfs/internal/pkg/common"
	"dedupfs/internal/pkg/fsutils"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
)

type SimpleImpl struct {
	collect   chan common.FileRecord
	waitGroup *sync.WaitGroup
}

func New(collect chan common.FileRecord, waitGroup *sync.WaitGroup) Scanner {
	return &SimpleImpl{
		collect:   collect,
		waitGroup: waitGroup,
	}
}

func (scan *SimpleImpl) ScanDir(verbose bool, dirs ...string) {

	for _, dir := range dirs {

		fileSystem := os.DirFS(dir)

		_ = fs.WalkDir(
			fileSystem,
			".",
			func(path string, d fs.DirEntry, err error) error {

				if !d.IsDir() {

					filePath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, path)
					scan.waitGroup.Add(1)
					go func() {
						defer scan.waitGroup.Done()
						if verbose {
							log.Printf("checking file: %s\n", filePath)
						}
						hash, hashErr := fsutils.HashFile(filePath)
						if hashErr != nil {
							log.Println(hashErr)
						} else {
							scan.collect <- common.FileRecord{
								FilePath: filePath,
								Hash:     hash,
							}
						}

					}()
				}

				return err
			})
	}
}
