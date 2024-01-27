BINARY=email-scan
BUILD_DIR=./build

all: build-dir build-linux-amd64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64

build-dir:
	- mkdir ${BUILD_DIR}

build-linux-amd64: ${BUILD_DIR}
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY}.linux.amd64
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY}.macos.amd64
	GOOS=darwin GOARCH=arm64 go build -o ${BUILD_DIR}/${BINARY}.macos.arm64
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY}.exe

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean all