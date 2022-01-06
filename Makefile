
.PHONY: help test testgc cover

help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## test: 测试
test:
	go test -v -coverprofile=/tmp/cover.out -covermode=atomic -race ./...

## testgc: gc 测试
testgc:
	go test -v -coverprofile=/tmp/cover.out -covermode=atomic -race -gcflags "all=-N -l" ./...

## cover: 覆盖率
cover:
	go tool cover -html=/tmp/cover.out
