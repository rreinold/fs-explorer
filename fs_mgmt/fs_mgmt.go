package fs_mgmt

import (
	"fmt"
	"io/ioutil"
	"os"
	. "path/filepath"
	"syscall"
)

// Based upon https://www.npmjs.com/package/directory-tree
type FileDetails struct {
	Name        string        `json:"name"`
	Owner       int           `json:"owner"`
	Size        int64         `json:"size"`
	Permissions string        `json:"permissions"`
	IsDir       bool          `json:"isDir"`
	Children    []FilePreview `json:"links"`
	Path        string        `json:"path"`
	Contents    string        `json:"contents"`
}

type FilePreview struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
	Href  string `json:"href"`
	Type  string `json:"type"`
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

func GetFileDetails(absPath string, relPath string) (FileDetails, error) {
	osFileInfo, osErr := os.Stat(absPath)
	if osErr != nil {
		return FileDetails{}, osErr
	}
	contents, readErr := ioutil.ReadFile(absPath)
	if readErr != nil {
		return FileDetails{}, readErr
	}
	// TODO try catch
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
	}
	return file, nil
}

func getFilePreview(absPath string, relPath string) (FilePreview, error) {
	DefaultType := "GET"
	osFileInfo, osErr := os.Stat(absPath)
	if osErr != nil {
		return FilePreview{}, osErr
	}

	file := FilePreview{
		Name:  osFileInfo.Name(),
		IsDir: osFileInfo.IsDir(),
		Path:  relPath,
		Href:  relPath,
		Type:  DefaultType}
	return file, nil
}

// Get a Directory's FileDetails, and children's FilePreviews
func GetDir(absPath string, relPath string) (FileDetails, error) {
	osFileInfo, osErr := os.Stat(absPath)
	if osErr != nil {
		return FileDetails{}, osErr
	}
	// TODO try catch
	owner := int(osFileInfo.Sys().(*syscall.Stat_t).Uid)
	filePreviews := []FilePreview{}
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		fmt.Println("Unable to access directory")
		return FileDetails{}, err
	}

	for _, f := range files {
		rel := Join(relPath, f.Name())
		abs := Join(absPath, f.Name())
		filePreview, previewErr := getFilePreview(abs, rel)
		if previewErr != nil {
			fmt.Println("Unable to fetch file preview: ", previewErr.Error())
			return FileDetails{}, previewErr
		}
		filePreviews = append(filePreviews, filePreview)
	}

	file := FileDetails{
		Name:        osFileInfo.Name(),
		Size:        osFileInfo.Size(),
		Permissions: osFileInfo.Mode().Perm().String(),
		IsDir:       osFileInfo.IsDir(),
		Owner:       owner,
		Path:        relPath,
		Children:    filePreviews,
	}

	return file, nil
}
