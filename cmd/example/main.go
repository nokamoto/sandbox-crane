package main

import (
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
)

func main() {
	src := os.Getenv("SRC")
	newTag := os.Getenv("NEW_TAG")
	var opts []crane.Option

	cfg := authn.AuthConfig{
		Username: os.Getenv("GITHUB_USER"),
		Password: os.Getenv("AUTH"),
	}

	opts = append(opts, crane.WithAuth(authn.FromConfig(cfg)))

	img, err := crane.Pull(src, opts...)
	if err != nil {
		log.Fatalf("err Pull: %s: %s", src, err)
	}

	digest, err := img.Digest()
	if err != nil {
		log.Fatalf("err Pull: %s: %s", src, err)
	}
	log.Printf("Pull: %s", digest)

	if len(os.Args) > 1 {
		paths := os.Args[1:]
		img, err = crane.Append(img, paths...)

		if err != nil {
			log.Fatalf("err Append: %s: %s", src, err)
		}

		digest, err := img.Digest()
		if err != nil {
			log.Fatalf("err Append: %s: %s", src, err)
		}
		log.Printf("Append: %s", digest)

		if err := crane.Push(img, newTag, opts...); err != nil {
			log.Fatalf("err Push: %s: %s: %s", src, newTag, err)
		}

		log.Printf("Push: %s: %s", src, newTag)
	}
}
