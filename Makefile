BINARY_DIR=./cmd/uptfs

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building Go uptfs..."
	go build $(BINARY_DIR)

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f ./uptfs

.PHONY: run
run:
	./uptfs