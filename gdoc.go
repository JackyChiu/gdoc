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

// buildPackageURL gets the url of the specificed package doc
func buildPackageURL(pkg string) (string, error) {
	if isStdPackage(pkg) {
		return fmt.Sprintf("%v%v/pkg/%v", HOST, PORT, pkg), nil
	}

	path, err := filepath.Abs(pkg)
	if err != nil {
		return "", err
	}

	gopath := os.Getenv("GOPATH")
	if !strings.Contains(path, gopath) {
		return "", fmt.Errorf("[ERR] Make sure your package is in your $GOPATH")
	}
	pkgPath := path[len(gopath+"/src/"):]

	return fmt.Sprintf("%v%v/pkg/%v", HOST, PORT, pkgPath), nil
}

// isStdPackage checks if the specified package is apart of the std library
func isStdPackage(pkg string) bool {
	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		return false
	}

	goroot := strings.Trim(string(out), "\n")
	path := fmt.Sprintf("%v/src/%v", goroot, pkg)
	if _, err = os.Stat(path); err != nil {
		return false
	}
	return true
}
