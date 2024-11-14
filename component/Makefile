# Variables for target and default directory
OUTPUT_DIR := ./build
TARGET := wasip2

# If FILE is not provided, default to "file.wasm"
FILE ?= file.wasm
WIT_WORLD ?= component

# Output file combines directory and the provided or default file name
OUTPUT_FILE := $(OUTPUT_DIR)/$(FILE)

# Default rule to build the project
all: gen-files build-project

gen-files: $(OUTPUT_DIR)
	mkdir -p gen && go generate ./...
# Ensure the output directory exists, then build with TinyGo
build-project: $(OUTPUT_DIR)
	tinygo build -o $(OUTPUT_FILE) -target $(TARGET) --wit-package ./wit --wit-world $(WIT_WORLD)

# Create the output directory if it does not exist
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Clean up the build files
clean:
	rm -rf $(OUTPUT_DIR)

