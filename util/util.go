package util

import (
	"io/ioutil"
	"os"
	"regexp"
)

func IsForbiddenPath(input string) bool {
	// Forbid '..' which could move out of hosted dir
	// Forbid '~' to prevent referencing user home dir
	REGEXP := `(\.\.|~)`
	match, err := regexp.Match(REGEXP, []byte(input))
	if err != nil {
		return true
	}
	return match
}

func GetWorkingDir() string {
	rootDir, _ := os.Getwd()
	return rootDir
}

func FileExists(absPath string) bool {
	_, err := os.Stat(absPath)
	return os.IsExist(err)
}

func IsDir(absPath string) (bool, error) {
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func GetFileContents(absPath string) (string, error) {
	contents, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(contents), err
}
