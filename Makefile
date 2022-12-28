help:
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-16s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# 单元测试
test:
	go test -v -coverprofile=/tmp/cover.out -covermode=atomic -race ./...

# 单元测试 -gcflags
testgc:
	go test -v -coverprofile=/tmp/cover.out -covermode=atomic -race -gcflags "all=-N -l" ./...

# 生成覆盖率报告
cover: test
	go tool cover -html=/tmp/cover.out
