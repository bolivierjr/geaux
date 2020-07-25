test:
	go clean -testcache
	go test -v ./cmd
	go test -v ./pkg/fs

.PHONY: all
