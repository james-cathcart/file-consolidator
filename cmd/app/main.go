package main

import (
	"dedupfs/internal/pkg/fsutils"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	fmt.Println("Starting app")

	var dirs []string
	for _, v := range os.Args[1:] {
		dirs = append(dirs, v)
	}

	for _, dir := range dirs {

		fileSystem := os.DirFS(dir)

		err := fs.WalkDir(
			fileSystem,
			".",
			func(path string, d fs.DirEntry, err error) error {

				log.Printf("walking path: %s", path)

				if d == nil {
					log.Println("warn: DirEntry is nil")
				} else {
					log.Printf("DirEntry name: %s", d.Name())
				}

				if !d.IsDir() {

					filePath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, path)
					log.Printf("hashing file: %s", filePath)

					var hash string
					hash, hashErr := fsutils.HashFile(filePath)
					if hashErr != nil {
						log.Println(hashErr)
					}

					fmt.Printf("file: %s, hash: %s\n", path, hash)
				}

				return err
			})

		if !errors.Is(err, fs.SkipDir) {
			log.Println(err)
		}
	}

	fmt.Println("App finished...")
}
