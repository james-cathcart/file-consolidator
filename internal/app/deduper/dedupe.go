package deduper

import "dedupfs/internal/pkg/common"

type SimpleImpl struct {
}

func New() DeDuper {
	return &SimpleImpl{}
}

func (dedupe *SimpleImpl) DeDuplicate(files map[string][]common.FileRecord) (totalCount int, duplicateCount int, unique map[string]common.FileRecord, duplicates map[string][]common.FileRecord) {

	totalCount = 0
	duplicateCount = 0

	unique = make(map[string]common.FileRecord)
	duplicates = make(map[string][]common.FileRecord)

	for k, v := range files {
		if _, ok := unique[k]; !ok {
			unique[k] = v[0]
			duplicates[k] = v[1:]
			totalCount += len(duplicates[k]) + 1
			duplicateCount += len(duplicates[k])
		}
	}

	return
}
