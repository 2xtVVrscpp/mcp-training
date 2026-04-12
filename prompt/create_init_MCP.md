# Create a Model Context Protocol (MCP) server that exposes a "hello_world" tool.
This tool, when called, should return the message "Hello, MCP world!"
For the MCP implementation, you should use the official Go SDK for MCP (github.com/modelcontextprotocol/go-sdk/mcp) and use the stdio transport.

## TODO:
- Download the dependency: `go get github.com/modelcontextprotocol/go-sdk/mcp`
- Inspect the documentation of the SDK: `go doc github.com/modelcontextprotocol/go-sdk/mcp`
- Build a `server` command that supports stdio transport only
- Build a `client` command that connects to the server over command transport to test the server

## Acceptance Criteria:
- `./bin/client --list-tools` returns the list of server tools including "hello_world"
- `./bin/client --call-tool` "hello_world" returns the output "Hello, MCP world!"