GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

.PHONY: help
help:
	@echo "fmt      re-format source codes."
	@echo "build    build binary from source code as './bin/ucloud-cli'."
	@echo "test     run unit test cases."
	@echo "test-acc run acc test cases."
	@echo "test-cov run unit test cases with coverage reporting."

.PHONY: fmt
fmt:
	gofmt -w -s $(GOFMT_FILES)

.PHONY: fmtcheck
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: fmtcheck vet
	go test -v ./ucloud/... --parallel=4

.PHONY: test-acc
test-acc: fmtcheck vet
	go test -v ./tests/... --parallel=32

.PHONY: test-cov
test-cov: fmtcheck vet
	go test -cover -coverprofile=coverage.out ./... --parallel=32

.PHONY: cyclo
cyclo:
	gocyclo -over 15 ucloud/ services/
