package deduper

import "dedupfs/internal/pkg/common"

type DeDuper interface {
	DeDuplicate(files map[string][]common.FileRecord) (totalCount int, duplicateCount int, unique map[string]common.FileRecord, duplicates map[string][]common.FileRecord)
}
