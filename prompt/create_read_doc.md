# Add a new tool to our MCP server called "read_docs" that invokes the "go doc" shell command.
The tool will take a mandatory "package" argument and an optional "symbol" argument.

## TODO:
- create a package `./internal/tools/docs`
- register the tool with the MCP server
- update the client to support the "read_docs" tool by providing arguments to the tool call

## Acceptance Criteria:
- `./bin/client --tools-list` show both hello_world and read_docs
- `./bin/client --tool-call read_docs fmt` returns the documentation for the `fmt` package
- `./bin/client --tool-call read_docs fmt.Println` returns the documentation for the `fmt.Println` function
- `./bin/client --tool-call read_docs github.com/modelcontextprotocol/go-sdk/mcp` returns documentation for the `mcp` package