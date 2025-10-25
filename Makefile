BIN_DIR := bin

.PHONY: build clean graph-demo ipfs-hello

build: graph-demo ipfs-hello

graph-demo:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/graph-demo ./cmd/graph-demo

ipfs-hello:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/ipfs-hello ./cmd/ipfs-hello

clean:
	rm -rf $(BIN_DIR)