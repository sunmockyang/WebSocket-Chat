BIN_FOLDER=Server-side/bin
EXE=ChatServer
SRC_PATH=`pwd`/Server-side
PKG_PATH=${SRC_PATH}/pkg
WEBSOCKET_LIB_PATH=${SRC_PATH}/src/github.com

all:
	@mkdir -p ${BIN_FOLDER}

	@export GOPATH=${SRC_PATH}; \
	go build -o ${BIN_FOLDER}/${EXE} Server-side/src/WebSocket-Chat.go
	
	@echo "Running ${BIN_FOLDER}/${EXE}..."
	
	@./${BIN_FOLDER}/${EXE}

websocket:
	export GOPATH=${SRC_PATH}; \
	go get github.com/gorilla/websocket

clean:
	@echo "Deleting ${BIN_FOLDER}/..."
	rm -rf ${BIN_FOLDER}
	@echo "Deleting ${PKG_PATH}/..."
	rm -rf ${PKG_PATH}
	@echo "Deleting ${WEBSOCKET_LIB_PATH}/..."
	rm -rf ${WEBSOCKET_LIB_PATH}
