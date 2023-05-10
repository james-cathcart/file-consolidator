package runner

import "dedupfs/internal/pkg/common"

type Runner interface {
	SearchForDuplicates(config common.SearchConfig)
	DeDuplicate(
		config common.DeDupeConfig,
		report *common.SearchReport,
	) (common.DeDupeReport, error)
}
