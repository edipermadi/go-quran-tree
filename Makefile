GO	:= $(shell which go)

.PHONY: test
test: unit

.PHONY: unit
unit:
	$(GO) test -v ./...