package docs

import (
	"context"
	"fmt"
	"os/exec"

	"godoctor/internal/mcputil"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	// ToolName defines the name of the read_docs tool.
	ToolName = "read_docs"
	// ToolDescription defines the purpose of the read_docs tool.
	ToolDescription = "Returns the documentation for a package or symbol"
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
		msg := fmt.Sprintf("Error running go doc %s: %v\nOutput: %s", target, err, string(output))
		return mcputil.NewTextResult(msg), nil, nil
	} else {
		msg := string(output)
		return mcputil.NewErrorResult(msg), nil, nil
	}
}
