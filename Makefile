TARGET_DIR := build
TARGET := breathe
SRC := $(wildcard *.go)

.PHONY: build
build: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET_DIR)/$(TARGET)

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR)
