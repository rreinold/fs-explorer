package main

import (
	"flag"
	"fmt"
	. "fs-explorer/fs_mgmt"
	"fs-explorer/util"
	"path"

	"github.com/gin-gonic/gin"
)

// Default to working directory
var DefaultDir string = "."

// Only GET is supported in v0.1.0
var SupportedMethod = "GET"

var rootDir string

func main() {

	flag.StringVar(&rootDir, "d", DefaultDir, "Path to directory to host")
	flag.Parse()
	if rootDir == DefaultDir {
		rootDir = GetWorkingDir()
	}
	fmt.Println("We host at: " + rootDir)
	initialize(rootDir)
}

// Start up web server, configure route(s)
func initialize(rootDir string) {
	router := gin.Default()
	router.NoRoute(processRequest)
	router.Run()
}

// Handle all HTTP requests
func processRequest(c *gin.Context) {
	// Enforce supported methods
	if c.Request.Method != SupportedMethod {
		c.JSON(501, "Received an unsupported method")
		return
	}

	relPath := c.Request.URL.Path
	fmt.Println("Request received, fetching from ", rootDir, " with relative path ", relPath)

	// Ensure path is not outside hosted dir
	if util.IsForbiddenPath(relPath) {
		c.JSON(403, "Requested forbidden filesystem path")
		return
	}

	// Check if file exists
	absPath := path.Join(rootDir, relPath)
	fmt.Println("Looking for file: ", absPath)
	if !FileExists(absPath) {
		c.JSON(404, "File not found: "+relPath)
		return
	}
	// Handle dir requests
	isDir, dirError := IsDir(absPath)
	if dirError != nil {
		c.JSON(500, "Unable to access file")
		return
	}
	if isDir {
		dir, dirErr := GetDir(absPath, relPath)
		if dirErr != nil {
			c.JSON(500, "Unable to access dir")
			return
		}
		c.JSON(200, dir)
		return
	} else {
		// Handle file requests
		fileDetails, readErr := GetFileDetails(absPath, relPath)
		if readErr != nil {
			c.JSON(500, "Unable to access file")
			return
		}
		c.JSON(200, fileDetails)
	}

}
