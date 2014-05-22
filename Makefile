GOPATH := $(shell pwd)

all: test

test:
	GOPATH=$(GOPATH) go install $@

clean:
	GOPATH=$(GOPATH) go clean
	${RM} -r pkg bin

.PHONY: all
