package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	listTools := flag.Bool("list-tools", false, "List available tools")
	callTool := flag.String("call-tool", "", "Call a tool with the given name")
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
		res, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name: *callTool,
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
