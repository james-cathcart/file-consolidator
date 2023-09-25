package runner

import (
	"dedupfs/internal/pkg/common"
	"dedupfs/internal/pkg/fsutils"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
)

type DeDuplicator struct {
	data    map[string]common.UniqueFile
	collect chan common.UniqueFile
	wg      *sync.WaitGroup
}

func NewDeDuplicator(
	data map[string]common.UniqueFile,
	collect chan common.UniqueFile,
	wg *sync.WaitGroup,
) Runner {
	return &DeDuplicator{
		data:    data,
		collect: collect,
		wg:      wg,
	}
}

func (dd *DeDuplicator) SearchForDuplicates(
	config common.SearchConfig,
) {
	for _, dir := range config.SearchDirs {
		fmt.Printf("searching directory: %s\n", dir)
		go dd.searchDir(dir)
	}
}

func (dd *DeDuplicator) searchDir(dir string) {
	fsys := os.DirFS(dir)
	_ = fs.WalkDir(fsys, dir, func(path string, d fs.DirEntry, err error) error {
		filePath := fmt.Sprintf("%s%s%s", dir, os.PathSeparator, path)
		log.Printf("searching: %s\n", filePath)
		dd.wg.Add(1)
		go func(filePath string) {
			defer dd.wg.Done()
			hash, err := fsutils.HashFile(filePath)
			if err != nil {
				return
			}
			dd.collect <- common.UniqueFile{
				FilePath: filePath,
				Hash:     hash,
			}
		}(filePath)

		return err
	})

	fmt.Println(`Printing Unique Files`)
	for file := range dd.collect {
		if _, ok := dd.data[file.Hash]; !ok {
			dd.data[file.Hash] = file
		}
	}
}

func (dd *DeDuplicator) DeDuplicate(
	config common.DeDupeConfig,
	report *common.SearchReport,
) (
	deDupeReport common.DeDupeReport,
	err error,
) {

	return
}
