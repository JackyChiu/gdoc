package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	HOST = "http://localhost"
	PORT = ":6060"
)

func main() {
	godoc := exec.Command("godoc", "-http", PORT)
	godoc.Start()

	pkg := "."
	if len(os.Args) > 1 {
		pkg = os.Args[1]
	}

	url, err := buildPackageURL(pkg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Serving on: %v", url)

	err = exec.Command("open", url).Run()
	if err != nil {
		log.Fatal(err)
	}

	godoc.Wait()
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
