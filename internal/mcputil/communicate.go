package mcputil

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// NewTextResult creates a *mcp.CallToolResult with a single TextContent.
func NewTextResult(text string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: text,
			},
		},
	}
}

// NewErrorResult creates a *mcp.CallToolResult with a single TextContent
// and the IsError flag set to true.
func NewErrorResult(text string) *mcp.CallToolResult {
	res := NewTextResult(text)
	res.IsError = true
	return res
}
