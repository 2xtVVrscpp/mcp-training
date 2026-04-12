package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "hello-mcp",
		Version: "1.0.0",
	}, nil)

	// Add the hello_world tool
	// We use any for input since it doesn't take arguments
	mcp.AddTool(server, &mcp.Tool{
		Name:        "hello_world",
		Description: "Returns a hello message",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args any) (*mcp.CallToolResult, any, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: "Hello, MCP world!",
				},
			},
		}, nil, nil
	})

	// Run the server on the stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
