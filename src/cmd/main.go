package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"

	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "1.0.0"

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "yoophi/golang-self-update-example")
	fmt.Println("latest", latest)
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}

func main() {
	fmt.Printf("main (1): %s\n", version)
	fmt.Println("updating ...")
	doSelfUpdate()
	fmt.Println("done ...")
	fmt.Printf("main (2): %s\n", version)
}
