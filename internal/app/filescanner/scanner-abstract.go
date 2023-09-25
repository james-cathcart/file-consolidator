package filescanner

type Scanner interface {
	ScanDir(verbose bool, dir ...string)
}
