GO_TEST_ARGS ?= -race

test:
	go test -v . $(GO_TEST_ARGS) -coverprofile cover.out

