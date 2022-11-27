package main

import (
	"log"
	"os"
)

const (
    // Set in the Deployment manifest.
	terminationMessagePath = "/tmp/oomed-pod.log"
)

func main() {
	f, err := os.Create(terminationMessagePath)
	if err != nil {
		log.Fatalf("unable to create %s: %s", terminationMessagePath, err)
	}

	_, err = f.WriteString("OOMKilled")
	if err != nil {
		log.Fatalf("unable to write to %s: %s", terminationMessagePath, err)
	}
	f.Close()

	os.Exit(137)
}
