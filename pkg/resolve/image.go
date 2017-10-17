package resolve

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ernoaapa/can/pkg/fs"
	log "github.com/sirupsen/logrus"
)

// Image try to resolve what container image should be used to run project in the directory
func Image(projectDir string) (string, error) {
	if isNodeProject(projectDir) {
		return "docker.io/library/node:latest", nil
	} else if isGolangProject(projectDir) {
		return "docker.io/library/golang:latest", nil
	}

	return "", fmt.Errorf("Unable to resolve container image for project in directory [%s]", projectDir)
}

func isNodeProject(projectDir string) bool {
	nodePackageFile := filepath.Join(projectDir, "package.json")
	log.Debugf("Checking does [%s] file exist, if does use Node container image", nodePackageFile)
	if fs.FileExist(nodePackageFile) {
		return true
	}
	return false
}

var golangDirs = []string{".", "pkg", "cmd"}

func isGolangProject(projectDir string) bool {
	for _, goDir := range golangDirs {
		if containsFiles(filepath.Join(projectDir, goDir), ".go") {
			return true
		}
	}
	return false
}

func containsFiles(path, extension string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Warnf("Unable to list files in [%s] directory. Cannot detect golang project", path)
		return false
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), extension) {
			return true
		}
	}
	return false
}