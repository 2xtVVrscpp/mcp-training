# create_read_doc
## commands

```
$ mkdir -p internal/tools/docs
$ go doc github.com/modelcontextprotocol/go-sdk/mcp.Tool
$ go doc github.com/modelcontextprotocol/go-sdk/mcp.AddTool
$ go build -o bin/server cmd/server/main.go && go build -o bin/client cmd/client/main.go && ./bin/client --tools-list
$ ./bin/client --tool-call read_docs fmt
$ ./bin/client --tool-call read_docs fmt.Println
$ ./bin/client --tool-call read_docs github.com/modelcontextprotocol/go-sdk/mcp
$ ./bin/client --tool-call hello_world
```
