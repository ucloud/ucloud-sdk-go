.PHONY: test validate-dco validate-gofmt example

default: build

test:
	script/test

validate-dco:
	script/validate-dco

validate-gofmt:
	script/validate-gofmt

validate: validate-dco validate-gofmt test

build: clean
	script/build
example: build
	script/example

clean:
	rm -rf Godeps/_workspace/pkg
