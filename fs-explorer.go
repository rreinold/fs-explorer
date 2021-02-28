package main

import (
	"flag"
	"fmt"
	"fs-explorer/util"
	"github.com/gin-gonic/gin"
	"path"
)

var DefaultDir string = "."
var rootDir string

// Based upon https://www.npmjs.com/package/directory-tree
type File struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Size        int    `json:"size"`
	Permissions string `json:"permissions"`
	IsDir       bool   `json:"isDir"`
	Children    []File `json:"children"`
	Path        string `json:"path"`
}

func main() {

	flag.StringVar(&rootDir, "d", DefaultDir, "Directory to host (Default: '.' )")
	flag.Parse()
	if rootDir == DefaultDir {
		rootDir = util.GetWorkingDir()
	}
	fmt.Println("We go host some stuff at " + rootDir)
	initialize(rootDir)
}

func initialize(rootDir string) {
	router := gin.Default()
	router.NoRoute(processRequest)
	router.Run(":3000")
}

func processRequest(c *gin.Context) {
	relativePath := c.Request.URL.Path
	fmt.Println("Request received, fetching from ", rootDir, " with relative path ", relativePath)
	// TODO URL Decode
	if util.IsForbiddenPath(relativePath) {
		c.JSON(403, "Requested forbidden filesystem path")
		return
	}
	absPath := path.Join(rootDir, relativePath)
	fmt.Println("Looking for file: ", absPath)
	if !util.FileExists(absPath) {
		c.JSON(404, "File not found: "+absPath)
		return

	}
	isDir, dirError := util.IsDir(absPath)
	if dirError != nil {
		c.JSON(500, "Unable to access file")
		return
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

}
