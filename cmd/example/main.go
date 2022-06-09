package main

import (
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
)

func main() {
	src := os.Getenv("SRC")
	var opts []crane.Option

	img, err := crane.Pull(src, opts...)
	if err != nil {
		log.Fatalf("err Pull: %s: %s", src, err)
	}

	digest, err := img.Digest()
	if err != nil {
		log.Fatalf("err Pull: %s: %s", src, err)
	}
	log.Printf("Pull: %s", digest)
}
