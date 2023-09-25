package common

// SearchConfig defines the criteria for searching for duplicates
type SearchConfig struct {
	SearchDirs []string
}

// SearchReport provides the result of the duplicate search
type SearchReport struct {
	Unique    map[string]string
	Duplicate map[string][]string
}

// DeDupeConfig defines the criteria for the deduplication process
type DeDupeConfig struct {
	UniqueDir string
}

// DeDupeReport provides the result of the deduplication process
type DeDupeReport struct {
	Errs []error
}

type FilesFound struct {
	UniqueFiles map[string]string
}

type FileRecord struct {
	FilePath string
	Hash     string
}
