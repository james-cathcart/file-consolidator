package deduper

import "dedupfs/internal/pkg/common"

type DeDuper interface {
	DeDuplicate(files map[string][]common.FileRecord) common.DeDupeOutput
}
