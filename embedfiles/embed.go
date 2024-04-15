package embedfiles

import (
	"io/fs"
	"embed"
	"log"
	"os"
	"fmt"
)

//go:embed web/*
var static embed.FS

// EmbedFiles embeds the files in the web and share folders
func EmbedFiles() error {
	return nil
}

// GetWebFS returns the web folder
func GetWebFS() fs.FS {
	// Create a new file system from the embedded web content
	webFS, err := fs.Sub(static, ".")
	if err != nil {
		log.Fatal(err)
	}
	return webFS
}

func EmbedFilesToDisk() error {

	folder := ".static"

	_, err := os.Stat(folder)
	if err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	fmt.Println("Writing embedded files to disk...")
	err = writeFilesToDisk(folder+"/html", "web/html")
	if err != nil {
		return err
	}

	err = writeFilesToDisk(folder+"/css", "web/css")
	if err != nil {
		return err
	}

	err = writeFilesToDisk(folder+"/js", "web/js")
	if err != nil {
		return err
	}

	err = writeFilesToDisk(folder+"/share", "share")
	if err != nil {
		return err
	}

	return nil
}

func writeFilesToDisk(toPath, fromDir string) error {

	err := os.MkdirAll(toPath, 0755)
	if err != nil {
		return err
	}

	files, err := static.ReadDir(fromDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		
		rFile := fromDir+"/"+file.Name()
		// fmt.Println(rFile)
		data, err := static.ReadFile(rFile)
		if err != nil {
			return err
		}

		wFile := toPath + "/" + file.Name()
		err = os.WriteFile(wFile, data, 0644)
		if err != nil {
			return err
		}
	}
	
	return nil
}