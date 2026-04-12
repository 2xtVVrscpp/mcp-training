package docs

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ReadDocsArgs provides the necessary parameters to locate and retrieve 
// documentation for a specific Go package or symbol.
type ReadDocsArgs struct {
	Package string `json:"package" jsonschema:"The package name (e.g., fmt)"`
	Symbol  string `json:"symbol,omitempty" jsonschema:"Optional symbol name (e.g., Println)"`
}

// ReadDocsTool executes the 'go doc' command to fetch and return 
// documentation, enabling the MCP server to provide technical 
// context about Go code to the client.
func ReadDocsTool(ctx context.Context, req *mcp.CallToolRequest, args ReadDocsArgs) (*mcp.CallToolResult, any, error) {
	if args.Package == "" {
		return nil, nil, fmt.Errorf("package is required")
	}

	target := args.Package
	if args.Symbol != "" {
		target = fmt.Sprintf("%s.%s", args.Package, args.Symbol)
	}

	cmd := exec.CommandContext(ctx, "go", "doc", target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Error running go doc %s: %v\nOutput: %s", target, err, string(output)),
				},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(output),
			},
		},
	}, nil, nil
}
