
TOP := $(shell pwd)

all:
	GOPATH=$(TOP) go install all

clean:
	@rm -rf pkg bin

