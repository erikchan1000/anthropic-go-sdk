/*
Package anthropic provides a Go client for Anthropic's Computer Use API, allowing programs
to interact with Claude's computer use capabilities.

Overview

The package provides a simple interface to Anthropic's Computer Use API, handling all the
necessary authentication, request formation, and response parsing. It supports all available
tools including computer control, text editor, and bash commands.

Basic Usage

Create a new client and make a request:

    client := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))

    req := &anthropic.ComputerRequest{
        Model:     "claude-3-5-sonnet-20241022",
        MaxTokens: 1024,
        Tools: []anthropic.Tool{
            anthropic.ComputerTool(1024, 768, 1),
        },
        Messages: []anthropic.Message{
            {
                Role:    "user",
                Content: "Open Chrome and navigate to google.com",
            },
        },
    }

    resp, err := client.CreateComputer(req)
    if err != nil {
        log.Fatal(err)
    }

Tools

The package provides helper functions to create the three available tools:

  - ComputerTool: Control computer interactions (mouse, keyboard, screen)
  - TextEditorTool: Text editing capabilities
  - BashTool: Execute bash commands

Example using multiple tools:

    tools := []anthropic.Tool{
        anthropic.ComputerTool(1024, 768, 1),
        anthropic.TextEditorTool(),
        anthropic.BashTool(),
    }

Configuration

The client can be configured with custom settings:

    client := &anthropic.Client{
        APIKey:     "your-api-key",
        BaseURL:    "custom-url",  // Optional
        APIVersion: "2023-06-01",  // Optional
        HTTPClient: customClient,   // Optional
    }

Error Handling

The package provides detailed error information from the API:

    resp, err := client.CreateComputer(req)
    if err != nil {
        var apiErr *anthropic.Error
        if errors.As(err, &apiErr) {
            fmt.Printf("API Error: %s - %s\n", apiErr.Type, apiErr.Message)
        }
        return
    }

Rate Limits and Retries

The client respects Anthropic's rate limits. Users should implement their own retry
logic based on their needs.

Thread Safety

The Client is safe for concurrent use by multiple goroutines.

For more information about the Anthropic API, visit:
https://docs.anthropic.com/
*/
package anthropic
