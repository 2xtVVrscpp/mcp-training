package main

import (
	"context"
	"log"

	"godoctor/internal/tools/docs"
	"godoctor/internal/tools/hello"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "hello-mcp",
		Version: "1.0.0",
	}, nil)

	// Add the hello_world tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        hello.ToolName,
		Description: hello.ToolDescription,
	}, hello.HelloWorldTool)

	// Add the read_docs tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        docs.ToolName,
		Description: docs.ToolDescription,
	}, docs.ReadDocsTool)

	// Run the server on the stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
