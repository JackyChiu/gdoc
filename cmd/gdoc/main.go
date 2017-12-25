package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JackyChiu/gdoc"
)

func main() {
	cmd := gdoc.StartDocSever()
	defer cmd.Wait()

	pkg := "."
	if len(os.Args) > 1 {
		pkg = os.Args[1]
	}

	// wait for doc server to start
	time.Sleep(time.Second)

	url, err := gdoc.OpenPackage(pkg)
	if err != nil {
		log.Fatalf("Couldn't open package: %v", err)
	}
	fmt.Printf("opened at %v", url)
}
