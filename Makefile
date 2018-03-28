PROJECT_ROOT := $(realpath $(CURDIR))
GOPATH := $(PROJECT_ROOT)/go

GODEPS := github.com/go-sql-driver/mysql \
          github.com/mattn/go-sqlite3

GOSRC := $(shell find $(GOPATH)/src -name '*.go' | grep -v '/src/github.com/')

.PHONY: all
all:

.PHONY: go
go: $(GOPATH)/bin/fecdumper

$(GOPATH)/bin/fecdumper: $(GODEPS) $(GOSRC)
	GOPATH=$(GOPATH) go install fecdumper

github.com/%:
	GOPATH=$(GOPATH) go get -u $@
