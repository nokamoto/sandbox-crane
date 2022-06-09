
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

testdata.tar.gz:
	tar cf testdata.tar.gz test
	tar tf testdata.tar.gz

.PHONY: testdata
testdata:
	docker rmi ghcr.io/nokamoto/sandbox-crane/busybox:testdata
	docker run --rm ghcr.io/nokamoto/sandbox-crane/busybox:testdata cat test/testdata/a.txt
