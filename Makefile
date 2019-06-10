GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

all: deps build

build:
	mkdir ./dist
	$(GOCMD) mod vendor
	$(GOBUILD) --mod vendor -o ./dist/main *.go

clean:
	rm -rf ./dist
	rm -rf ./vendor
	rm ./go.sum

deps:
	$(GOMOD) tidy
	$(GOMOD) vendor

