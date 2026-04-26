package hello

import (
	"context"

	"mcp-training/internal/mcputil"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	// ToolName defines the name of the hello_world tool.
	ToolName = "hello_world"
	// ToolDescription defines the purpose of the hello_world tool.
	ToolDescription = "Returns a hello message"
)

// HelloWorldTool handles the hello_world tool call.
func HelloWorldTool(ctx context.Context, req *mcp.CallToolRequest, args any) (*mcp.CallToolResult, any, error) {
	return mcputil.NewTextResult("Hello, MCP world!"), nil, nil
}
