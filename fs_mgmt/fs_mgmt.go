package fs_mgmt

import (
	"io/ioutil"
	"os"
	"syscall"
)

// Based upon https://www.npmjs.com/package/directory-tree
type FileDetails struct {
	Name        string        `json:"name"`
	Owner       int           `json:"owner"`
	Size        int64         `json:"size"`
	Permissions string        `json:"permissions"`
	IsDir       bool          `json:"isDir"`
	Children    []FilePreview `json:"children"`
	Path        string        `json:"path"`
	Contents    string        `json:"contents"`
	URI         string        `json:"URI"`
}

type FilePreview struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
	URI   string `json:"URI"`
}

func GetWorkingDir() string {
	rootDir, _ := os.Getwd()
	return rootDir
}

// Reference for approach: https://golangr.com/file-exists/
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

func GetFileContents(absPath string, relPath string) (FileDetails, error) {
	osFileInfo, osErr := os.Stat(absPath)
	if osErr != nil {
		return FileDetails{}, osErr
	}
	contents, readErr := ioutil.ReadFile(absPath)
	if readErr != nil {
		return FileDetails{}, readErr
	}
	owner := int(osFileInfo.Sys().(*syscall.Stat_t).Uid)
	file := FileDetails{
		Name:        osFileInfo.Name(),
		Size:        osFileInfo.Size(),
		Permissions: osFileInfo.Mode().Perm().String(),
		IsDir:       osFileInfo.IsDir(),
		Contents:    string(contents),
		Owner:       owner,
		Path:        relPath,
		Children:    []FilePreview{},
		URI:         relPath}
	return file, nil
}
