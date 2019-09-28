SHELL := /bin/bash

OUTPUT_DIR := build
TARGET := $(OUTPUT_DIR)/cdgo

# go source files, ignore vendor directory
SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: build clean install run

$(TARGET): $(SRC)
	@go build -o $(TARGET)

build: TARGET
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install

run: install
	@$(TARGET)
