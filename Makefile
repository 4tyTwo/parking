GO := $(shell which go 2>/dev/null || which ./go)

all: compile

compile:
	$(GO) build
