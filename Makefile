TARGETS = $(BINS)/server $(BINS)/client
SERVER_SRC = $(shell find ./cmd/server -name "*.go")
CLIENT_SRC = $(shell find ./cmd/client -name "*.go")
TOOLS = $(shell find ./internal/tools -name "*.go")
MCP_UTILS = $(shell find ./internal/mcputil -name "*.go")
BINS = bin

all: $(TARGETS)
$(BINS)/server: $(SERVER_SRC) $(TOOLS) $(MCP_UTILS)
	go build -o $(BINS)/server cmd/server/main.go
$(BINS)/client: $(CLIENT_SRC) $(TOOLS) $(MCP_UTILS)
	go build -o $(BINS)/client cmd/client/main.go
clean:
	rm -rf $(TARGETS)
