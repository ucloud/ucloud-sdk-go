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

# CI 门禁：生成代码必须能编译、且无重复项。
#   G1 `go build ./...` 覆盖 services/ 下全部产品包（生成代码），语法/类型错误即失败。
#   G2 重复 import 与重复声明在 Go 是编译错误，已被 G1 覆盖，无需额外检测。
#      （其它语言不然：php/js/python 的编译器对重复声明一律放行。）
#   另加 gofmtcheck 拦格式漂移——全仓重生成后出现纯空白 diff 的历史问题。
.PHONY: ci-syntax
ci-syntax:
	go build ./...
	@bash $(CURDIR)/scripts/gofmtcheck.sh

.PHONY: lint
lint:
	go vet ./...

.PHONY: test
test: fmtcheck vet
	go test -v ./ucloud/... ./external/... --parallel=16

.PHONY: test-acc
test-acc: fmtcheck vet
	go test -v ./tests/... --parallel=32

.PHONY: test-cov
test-cov: fmtcheck
	go test -cover -coverprofile=coverage.txt ./ucloud/... --parallel=32

.PHONY: cov-preview
cov-preview:
	go tool cover -html=coverage.txt

.PHONY: cyclo
cyclo:
	gocyclo -over 15 ucloud/ services/ external/

# UCloud Tools Path
UCLOUD_TEMPLATE_PATH=../ucloud-api-model-v2/apisdk/lang/go/templates

.PHONY: gen
gen: gen-api gen-test
	@echo "generate code success"

gen-api:
	ucloud-model sdk apis \
		--lang go \
		--type public \
		--template ${UCLOUD_TEMPLATE_PATH}/scripts-api.tpl \
		--output ./scripts/gen-apis.sh

gen-test:
	ucloud-model sdk tests \
		--lang go \
		--template ${UCLOUD_TEMPLATE_PATH}/scripts-test.tpl \
		--output ./scripts/gen-tests.sh
