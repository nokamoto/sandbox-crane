# sandbox-crane

- https://sched.co/ytrA
- https://www.youtube.com/watch?v=Mi_fb5ToOa8
- https://github.com/google/go-containerregistry/tree/main/cmd/crane
- https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry

## crane.Pull

```bash
make busybox
```

Push busybox to [ghcr.io/nokamoto/sandbox-crane/busybox](https://github.com/nokamoto/sandbox-crane/pkgs/container/sandbox-crane%2Fbusybox).

```bash
SRC=ghcr.io/nokamoto/sandbox-crane/busybox:latest go run ./cmd/example
```

Pull busybox:latest and print sha256.
