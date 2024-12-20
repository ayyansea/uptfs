BINARY_DIR=./cmd

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building Go uptfs..."
	go build ./cmd/uptfs

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f ./uptfs

.PHONY: run
run:
	./uptfs