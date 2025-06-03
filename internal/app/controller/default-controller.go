package controller

import (
	"dedupfs/internal/app/deduper"
	"dedupfs/internal/app/filemigrator"
	"dedupfs/internal/app/filescanner"
	"dedupfs/internal/app/view"
)

type DefaultController struct {
	dedupUtil deduper.DeDuper
	fScanner  filescanner.Scanner
	fMigrator filemigrator.Migrator
	view      view.View
}

func NewDefaultController(deduperUtil deduper.DeDuper, fileScanner filescanner.Scanner, fileMigrator filemigrator.Migrator, view view.View) IDefaultController {
	return &DefaultController{
		dedupUtil: deduperUtil,
		fScanner:  fileScanner,
		fMigrator: fileMigrator,
		view: view,
	}
}

func (c *DefaultController) Scan() {

	

}
