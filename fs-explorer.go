package main

import (
	"flag"
	"fmt"
	. "fs-explorer/fs_mgmt"
	"fs-explorer/util"
	"path"

	"github.com/gin-gonic/gin"
)

var DefaultDir string = "."
var rootDir string

func main() {

	flag.StringVar(&rootDir, "d", DefaultDir, "Directory to host (Default: '.' )")
	flag.Parse()
	if rootDir == DefaultDir {
		rootDir = GetWorkingDir()
	}
	fmt.Println("We go host some stuff at " + rootDir)
	initialize(rootDir)
}

func initialize(rootDir string) {
	router := gin.Default()
	router.NoRoute(processRequest)
	router.Run()
}

func processRequest(c *gin.Context) {

	relPath := c.Request.URL.Path
	fmt.Println("Request received, fetching from ", rootDir, " with relative path ", relPath)

	if util.IsForbiddenPath(relPath) {
		c.JSON(403, "Requested forbidden filesystem path")
		return
	}

	absPath := path.Join(rootDir, relPath)
	fmt.Println("Looking for file: ", absPath)
	if !FileExists(absPath) {
		c.JSON(404, "File not found: "+relPath)
		return

	}
	isDir, dirError := IsDir(absPath)
	if dirError != nil {
		c.JSON(500, "Unable to access file")
		return
	}
	if isDir {
		dir, dirErr := GetDir(absPath, relPath)
		if dirErr != nil {
			c.JSON(500, "Unable to access dir")
		}
		c.JSON(200, dir)
		return
	} else {
		contents, readErr := GetFileDetails(absPath, relPath)
		if readErr != nil {
			c.JSON(500, "Unable to access file")
			return
		}
		c.JSON(200, contents)
	}

}
