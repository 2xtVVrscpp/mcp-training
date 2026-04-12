# 1st MCP server/client
## commands

```
$ go get github.com/modelcontextprotocol/go-sdk/mcp && go get github.com/modelcontextprotocol/go-sdk/server && go get github.com/modelcontextprotocol/go-sdk/client
# $ go list github.com/modelcontextprotocol/go-sdk/... #
$ go doc github.com/modelcontextprotocol/go-sdk/mcp
$ mkdir -p bin && go build -o bin/server cmd/server/main.go && go build -o bin/client cmd/client/main.go
$ ./bin/client --list-tools
$ ./bin/client --call-tool hello_world
$ gofmt -w cmd/server/main.go cmd/client/main.go
$ go mod tidy
$ gofmt -d cmd/server/main.go cmd/client/main.go
$ ./bin/client --list-tools && ./bin/client --call-tool hello_world
```
