package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"mcp-training/internal/tools/docs"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	listTools := flag.Bool("tools-list", false, "List available tools")
	callTool := flag.String("tool-call", "", "Call a tool with the given name")
	flag.Parse()

	if !*listTools && *callTool == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "hello-client",
		Version: "1.0.0",
	}, nil)

	// In a real scenario, the server binary would be built and placed in bin/
	// For testing, we expect the server binary to be at ./bin/server
	transport := &mcp.CommandTransport{
		Command: exec.Command("./bin/server"),
	}

	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer session.Close()

	if *listTools {
		res, err := session.ListTools(ctx, &mcp.ListToolsParams{})
		if err != nil {
			log.Fatalf("Failed to list tools: %v", err)
		}
		for _, tool := range res.Tools {
			fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		}
	} else if *callTool != "" {
		var args map[string]any
		switch *callTool {
		case docs.ToolName:
			remainingArgs := flag.Args()
			if len(remainingArgs) > 0 {
				target := remainingArgs[0]
				// Split into package and symbol if possible
				// For simplicity, we can just pass the first arg as package if no dot,
				// or split by dot.
				// But wait, the tool expects "package" and "symbol".
				// If target is "fmt.Println", we should split it.
				// However, if the user provides "github.com/mcp/go-sdk/mcp", there's no symbol.
				// Let's use a simple heuristic for now as requested.
				// "fmt" -> package: fmt
				// "fmt.Println" -> package: fmt, symbol: Println

				// Actually, "go doc" can take "fmt.Println".
				// Let's refine the logic to match the requirements.

				lastDot := -1
				for i := len(target) - 1; i >= 0; i-- {
					if target[i] == '.' {
						// Make sure it's not part of a domain (e.g., github.com/...)
						// A dot is a symbol separator if it's after the last slash.
						lastSlash := -1
						for j := i; j >= 0; j-- {
							if target[j] == '/' {
								lastSlash = j
								break
							}
						}
						if i > lastSlash {
							lastDot = i
						}
						break
					}
				}

				if lastDot != -1 {
					args = map[string]any{
						"package": target[:lastDot],
						"symbol":  target[lastDot+1:],
					}
				} else {
					args = map[string]any{
						"package": target,
					}
				}
			}
			val, ok := args["package"].(string)
			if !ok || val == "" {
				fmt.Printf("Package name is invalid.\n")
				fmt.Printf("Usage: $ client -tools-list read_docs ${Package name} \n")
				os.Exit(1)
			}
		}

		res, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name:      *callTool,
			Arguments: args,
		})
		if err != nil {
			log.Fatalf("Failed to call tool: %v", err)
		}
		for _, content := range res.Content {
			if text, ok := content.(*mcp.TextContent); ok {
				fmt.Println(text.Text)
			}
		}
	}
}
