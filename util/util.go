package util

import (
	"io/ioutil"
	"os"
	"regexp"
	"syscall"
)

// Based upon https://www.npmjs.com/package/directory-tree
type File struct {
	Name        string `json:"name"`
	Owner       int    `json:"owner"`
	Size        int64  `json:"size"`
	Permissions string `json:"permissions"`
	IsDir       bool   `json:"isDir"`
	Children    []File `json:"children"`
	Path        string `json:"path"`
	Contents    string `json:"contents"`
	URI         string `json:"URI"`
}

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

// Reference for methodology: https://golangr.com/file-exists/
func FileExists(absPath string) bool {
	_, err := os.Stat(absPath)
	return err == nil
}

func IsDir(absPath string) (bool, error) {
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func GetFileContents(absPath string, relPath string) (File, error) {
	osFileInfo, osErr := os.Stat(absPath)
	if osErr != nil {
		return File{}, osErr
	}
	contents, readErr := ioutil.ReadFile(absPath)
	if readErr != nil {
		return File{}, readErr
	}
	owner := int(osFileInfo.Sys().(*syscall.Stat_t).Uid)
	file := File{
		Name:        osFileInfo.Name(),
		Size:        osFileInfo.Size(),
		Permissions: osFileInfo.Mode().Perm().String(),
		IsDir:       osFileInfo.IsDir(),
		Contents:    string(contents),
		Owner:       owner,
		Path:        relPath,
		Children:    []File{},
		URI:         relPath}
	return file, nil
}
