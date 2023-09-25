package filemigrator

import "dedupfs/internal/pkg/common"

type Migrator interface {
	MigrateUniqueFiles(destination string, scannedDirs []string, files map[string]common.FileRecord) (err error)
}
