VERSION := $(shell go-versioner -defaultVersion=v0)

go-versioner: main.go
	go build -ldflags "-X main.VERSION '$(VERSION)'"
install: main.go
	go install -ldflags "-X main.VERSION '$(VERSION)'"
clean:
	rm -f go-versioner

.PHONY: clean install
