.PHONY: build test clean serve

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=blog
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

serve:
	./$(BINARY_NAME)