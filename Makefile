.DEFAULT_GOAL := test

test:
	GO111MODULE=on go test ./... -race -cover -mod=vendor -tags=component

.PHONY: test