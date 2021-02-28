package main

import (
	"flag"
	"fmt"
	"fs-explorer/util"
	"github.com/gin-gonic/gin"
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
		path := c.Request.URL.Path
		fmt.Println("Request received, fetching from ", rootDir, " with path ", path)
		// TODO URL Decode
		if util.IsForbiddenPath(path) {
			c.JSON(403, "Requested forbidden filesystem path")
			return
		}
		c.JSON(200, "About to get that for you")
	})
	router.Run(":3000")
}
