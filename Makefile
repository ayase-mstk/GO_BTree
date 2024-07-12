BINARY_NAME := BTree

MODULE_PATH := BTree
RM          := rm -rf

all: init fmt build

init:
	@if [ ! -f go.mod ]; then \
		go mod init ${MODULE_PATH}; \
	fi

build: ${SRC}
	go build -o ${BINARY_NAME} main.go

run: init fmt build
	./${BINARY_NAME}

fmt:
	gofmt -s -w .

clean:
	go clean
	${RM} go.mod

fclean: clean
	${RM} ${BINARY_NAME}

re: fclean all

.PHONY: init, build, run, fmt, clean, fclean, re
