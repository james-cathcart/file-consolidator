package common

// ScanReport defines the criteria for searching for duplicates
type ScanReport struct {
	FilesFound      map[string][]FileRecord
	UniqueFiles     map[string]FileRecord
	DuplicateFiles  map[string][]FileRecord
	TotalFiles      int
	DuplicatesFound int
	UniqueFileCount int
}

type DeDupeOutput struct {
	UniqueFiles    map[string]FileRecord
	DuplicateFiles map[string][]FileRecord
	TotalCount     int
	DuplicateCount int
}

type FileRecord struct {
	FilePath string
	Hash     string
}
