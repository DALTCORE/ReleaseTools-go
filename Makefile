NAME=release-tool
ROOT_DIR=$(shell pwd)
BUILD_DIR=${ROOT_DIR}/build
VERSION=$(shell cat VERSION)
BUILD_NUMBER=$(shell cat BUILD)
COMMIT=$(shell git rev-parse --short HEAD)
LD_FLAGS="-X main.RTVERSION=${VERSION} -X main.NAME=${NAME}"

.PHONY: all build clean fmt run bn

all: pre_build build post_build

pre_build: bn clean fmt
	echo ${BUILD_NUMBER}

build:
	mkdir -p ${ROOT_DIR}/build
	GOOS=linux GOARCH=amd64 go build -ldflags ${LD_FLAGS} -o ${ROOT_DIR}/build/${NAME}-linux-amd64 ${ROOT_DIR}/src/*.go
	GOOS=windows GOARCH=amd64 go build -ldflags ${LD_FLAGS} -o ${ROOT_DIR}/build/${NAME}-windows-amd64.exe ${ROOT_DIR}/src/*.go
	GOOS=darwin GOARCH=amd64 go build -ldflags ${LD_FLAGS} -o ${ROOT_DIR}/build/${NAME}-macos-amd64 ${ROOT_DIR}/src/*.go

post_build:
	@echo
	@echo '##################################'
	@echo '    OUTPUT FOR THE GIT RELEASE'
	@echo '##################################'
	@echo 'Version: `${VERSION}`'
	@echo 'Build number: `${BUILD_NUMBER}`'
	@echo 'Commit: `${COMMIT}`'
	@echo
	@echo 'SHA512 sum:'
	@echo '```'
	@cd ${BUILD_DIR} && sha512sum *
	@echo '```'
	@echo "##################################"

clean:
	rm -rf ${ROOT_DIR}/build

vet:
	go vet ${ROOT_DIR}/src/*.go

fmt:
	go fmt ${ROOT_DIR}/src/*.go

test:
	go test ${ROOT_DIR}/src/*.go -v

run:
	go run ${ROOT_DIR}/src/*.go

bn:
	@echo "${BUILD_NUMBER} + 1" | bc > BUILD
	@BUILD_NUMBER=${BUILD_NUMBER}

