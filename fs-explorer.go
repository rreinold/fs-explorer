package main

import (
	"flag"
	"fmt"
	"fs-explorer/util"
	"github.com/gin-gonic/gin"
	"path"
)

func main() {

	rootDir := flag.String("d", ".", "Hosted directory (Default: '.' )")
	flag.Parse()
	fmt.Println("We go host some stuff at " + *rootDir)
	initialize(rootDir)
}

func initialize(rootDir *string) {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		relativePath := c.Request.URL.Path
		fmt.Println("Request received, fetching from ", rootDir, " with relative path ", relativePath)
		// TODO URL Decode
		if util.IsForbiddenPath(relativePath) {
			c.JSON(403, "Requested forbidden filesystem path")
			return
		}
		absPath := path.Join(*rootDir, relativePath)
		c.JSON(200, "About to get that for you at "+absPath)
	})
	router.Run(":3000")
}
