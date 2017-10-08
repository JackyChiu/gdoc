package gdoc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Local browser constants
const (
	HOST = "http://localhost"
	PORT = ":6060"
)

// StartDocServer runs the godoc server and stops when you kill the process
func StartDocSever() *exec.Cmd {
	cmd := exec.Command("godoc", "-http", PORT)
	cmd.Start()
	return cmd
}

// OpenPackage opens the package in your local browser
func OpenPackage(pkg string) (string, error) {
	url, err := buildPackageURL(pkg)
	if err != nil {
		return "", err
	}
	if err = exec.Command("open", url).Run(); err != nil {
		return "", errors.Wrap(err, "can't open url")
	}
	return url, nil
}

// buildPackageURL creates the url for package doc
func buildPackageURL(pkg string) (string, error) {
	if path, ok := localPackagePath(pkg); ok {
		return formatURL(path), nil
	}
	if path, ok := standardPackagePath(pkg); ok {
		return formatURL(path), nil
	}
	return "", fmt.Errorf("[ERR] Package doesn't exist")
}

// localPackagePath gets the pkg path and checks if the package is a local package
func localPackagePath(pkg string) (string, bool) {
	path, err := filepath.Abs(pkg)
	if err != nil {
		return "", false
	}

	gopath := os.Getenv("GOPATH")
	_, err = os.Stat(path)
	if !strings.Contains(path, gopath) || err != nil {
		return "", false
	}
	return path[len(gopath+"/src/"):], true
}

// standardPackagePath get the pkg path and checks if the package is apart of the std package
func standardPackagePath(pkg string) (string, bool) {
	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		return "", false
	}

	goroot := strings.Trim(string(out), "\n")
	path := fmt.Sprintf("%v/src/%v", goroot, pkg)
	if _, err = os.Stat(path); err != nil {
		return "", false
	}
	return pkg, true
}

// formatURL formats the url to open in broswer
func formatURL(path string) string {
	return fmt.Sprintf("%v%v/pkg/%v", HOST, PORT, path)
}
