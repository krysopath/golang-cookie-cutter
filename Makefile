GIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null)
GIT_TAG := $(shell git describe --abbrev=0 HEAD 2>/dev/null)
LD_FLAGS := '-s -w -X main.gitTag=$(GIT_TAG) -X main.gitRef=$(GIT_SHA)'

install: bin/cookies
	cp bin/cookies $(GOPATH)/bin

bin/cookies.exe: *.go go.*
	go build \
		-o $@ \
		-ldflags=$(LD_FLAGS) \
		-trimpath

bin/cookies: *.go go.*
	go build \
		-o $@ \
		-ldflags=$(LD_FLAGS) \
		-trimpath

deps.txt: go.mod go.sum
	go get
	go mod graph > deps.txt

