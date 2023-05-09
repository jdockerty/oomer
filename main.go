package main

import (
	"log"
	"os"
)

const (
	// Set in the Deployment manifest.
	terminationMessagePath = "/tmp/oomed-pod.log"
	message                = "OOMKilled"
)

func main() {
	f, err := os.Create(terminationMessagePath)
	if err != nil {
		log.Fatalf("unable to create %s: %s", terminationMessagePath, err)
	}
	defer f.Close()

	_, err = f.WriteString(message)
	if err != nil {
		log.Fatalf("unable to write to %s: %s", terminationMessagePath, err)
	}

	log.Println(message)
	os.Exit(137)
}
