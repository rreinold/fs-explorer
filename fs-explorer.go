package main

import (
	"flag"
	"fmt"
	"fs-explorer/util"
	"github.com/gin-gonic/gin"
	"path"
)

func main() {

	var DEFAULT_DIR string = "."

	var rootDir string
	flag.StringVar(&rootDir, "Hosted directory", DEFAULT_DIR, " (Default: '.' )")
	flag.Parse()
	if rootDir == DEFAULT_DIR {
		rootDir = util.GetWorkingDir()
	}
	fmt.Println("We go host some stuff at " + rootDir)
	initialize(rootDir)
}

func initialize(rootDir string) {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		relativePath := c.Request.URL.Path
		fmt.Println("Request received, fetching from ", rootDir, " with relative path ", relativePath)
		// TODO URL Decode
		if util.IsForbiddenPath(relativePath) {
			c.JSON(403, "Requested forbidden filesystem path")
			return
		}
		absPath := path.Join(rootDir, relativePath)
		if util.FileExists(absPath) {
			c.JSON(404, "File not found: "+relativePath)
		}
		isDir, dirError := util.IsDir(absPath)
		if dirError != nil {
			c.JSON(500, "Unable to access file")
		}
		if isDir {
			c.JSON(200, "About to get that whole dir at "+absPath)
			return
		} else {
			contents, readErr := util.GetFileContents(absPath)
			if readErr != nil {
				c.JSON(500, "Unable to read file")
				return
			}
			c.JSON(200, contents)
		}

	})
	router.Run(":3000")
}
