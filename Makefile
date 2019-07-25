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
	@bash $(CURDIR)/scripts/gofmtcheck.sh

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: fmtcheck vet
	go test -v ./ucloud/... ./external/... --parallel=16

.PHONY: test-acc
test-acc: fmtcheck vet
	go test -v ./tests/... --parallel=32

.PHONY: test-cov
test-cov: fmtcheck
	go test -cover -coverprofile=coverage.out ./ucloud/... --parallel=32

.PHONY: cov-preview
cov-preview:
	go tool cover -html=coverage.out

.PHONY: cyclo
cyclo:
	gocyclo -over 15 ucloud/ services/ external/

# UCloud Tools Path
UCLOUD_TEMPLATE_PATH=../ucloud-api-model-v2/apisdk/lang/go/templates

.PHONY: gen
gen:
	ucloud-model sdk apis \
	    --product StepFlow \
		--lang go \
		--type public \
		--template ${UCLOUD_TEMPLATE_PATH}/scripts-api.tpl \
		--output ./scripts/gen-apis.sh
