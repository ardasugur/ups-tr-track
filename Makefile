# Makefile for the UPS Package Tracker

# Variables
BINARY_NAME=upsTrack.exe
MAIN_FILE=main.go

# Default target
.PHONY: all
all: build

# Build target with size optimization
.PHONY: build
build:
	go build -ldflags "-s -w" -o $(BINARY_NAME) $(MAIN_FILE)

# Clean target
.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe



