# sandbox-crane

- https://sched.co/ytrA
- https://www.youtube.com/watch?v=Mi_fb5ToOa8
- https://github.com/google/go-containerregistry/tree/main/cmd/crane
- https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry

## crane.Pull

```bash
make busybox
```

Push busybox to [ghcr.io/nokamoto/sandbox-crane/busybox](https://ghcr.io/nokamoto/sandbox-crane/busybox).

```bash
SRC=ghcr.io/nokamoto/sandbox-crane/busybox:latest go run ./cmd/example
```

Pull busybox:latest and print sha256.

## crane.Append and crane.Push

```bash
make testdata.tar.gz
```

Make `testdata.tar.gz` from the `test` directory.

```bash
SRC=ghcr.io/nokamoto/sandbox-crane/busybox:latest NEW_TAG=ghcr.io/nokamoto/sandbox-crane/busybox:testdata go run ./cmd/example testdata.tar.gz
```

Append `testdata.tar.gz` to busybox:latest and push a new tag busybox:testdata.

```bash
make testdata
```

Run busybox:testdata and print a content of `test/testdata/a.txt`.
