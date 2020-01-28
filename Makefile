test:
	go clean -testcache
	go test -v ./cmd

.PHONY: all
