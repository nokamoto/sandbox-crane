
.PHONY: all
all:
	go fmt ./...
	go test ./...

.PHONY: all
busybox:
	echo $$AUTH | docker login ghcr.io -u $$GITHUB_USER --password-stdin
	docker pull busybox
	docker tag busybox ghcr.io/nokamoto/sandbox-crane/busybox:latest
	docker push ghcr.io/nokamoto/sandbox-crane/busybox:latest
	docker rmi busybox
	docker rmi ghcr.io/nokamoto/sandbox-crane/busybox

