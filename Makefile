ARTIFACT_NAME = bookgo

ifeq ($(OS),Windows_NT)
	EXT = .exe
else
	EXT =
endif

.PHONY: build run clean

build:
	@go build -o bin/${ARTIFACT_NAME}${EXT} cmd/${ARTIFACT_NAME}/main.go

run:
	@go run cmd/${ARTIFACT_NAME}/main.go

clean:
	@rm -rf bin