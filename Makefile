.PHONY: build ui

all: fmt deps build

deps:
	@go list github.com/mjibson/esc || go get github.com/mjibson/esc/...
	@go list golang.org/x/tools/cmd/goimports || go get golang.org/x/tools/cmd/goimports
	go generate -x
	go get .

fmt:
	goimports -w .
	go vet .

run:
	./bin/eval-server -L debug --ui-dir ui run

build: fmt
	go build -o bin/eval-server ./
