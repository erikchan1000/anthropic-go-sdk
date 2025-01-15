// README.md
# Anthropic Computer Use SDK for Go

A Go SDK for interacting with Anthropic's Computer Use API.

## Installation

```bash
go get github.com/yourusername/anthropic-computer
```

## Usage

```go
import "github.com/yourusername/anthropic-computer"

func main() {
    client := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))

    req := &anthropic.ComputerRequest{
        Model:     "claude-3-5-sonnet-20241022",
        MaxTokens: 1024,
        Tools: []anthropic.Tool{
            anthropic.ComputerTool(1024, 768, 1),
            anthropic.TextEditorTool(),
            anthropic.BashTool(),
        },
        Messages: []anthropic.Message{
            {
                Role:    "user",
                Content: "Save a picture of a cat to my desktop.",
            },
        },
    }

    resp, err := client.CreateComputer(req)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Documentation

For more detailed documentation, see the [Anthropic API docs](https://docs.anthropic.com/).


Package anthropic provides a Go client for the Anthropic Computer Use API.

Example usage:

    client := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))

    req := &anthropic.ComputerRequest{
        Model:     "claude-3-5-sonnet-20241022",
        MaxTokens: 1024,
        Tools: []anthropic.Tool{
            anthropic.ComputerTool(1024, 768, 1),
            anthropic.TextEditorTool(),
            anthropic.BashTool(),
        },
        Messages: []anthropic.Message{
            {
                Role:    "user",
                Content: "Save a picture of a cat to my desktop.",
            },
        },
    }

    resp, err := client.CreateComputer(req)
    if err != nil {
        log.Fatal(err)
    }
