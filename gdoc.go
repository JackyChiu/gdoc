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

	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	url, err := buildPackageURL(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)

	err = exec.Command("open", url).Run()
	if err != nil {
		log.Fatal(err)
	}

	godoc.Wait()
}

func buildPackageURL(dir string) (string, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("[ERR] %v doesn't exsit", dir)
	}

	gopath := os.Getenv("GOPATH")
	if !strings.Contains(path, gopath) {
		return "", fmt.Errorf("[ERR] Make sure your package is in your $GOPATH")
	}

	pkg := path[len(gopath+"/src"):]
	return fmt.Sprintf("%v%v/pkg%v", HOST, PORT, pkg), nil
}
