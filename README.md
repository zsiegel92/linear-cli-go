# Linear CLI Go

A Go implementation of a Linear CLI tool for interacting with Linear issues.

## Features

- Load issues from Linear API
- Filter issues by assignee (me) or project
- Display a list of issues with a basic command-line selector
- Allow selection of an issue from the list
- Output the selected issue as formatted JSON

## Installation

```bash
# Clone the repository
git clone https://github.com/zsiegel92/linear-cli-go.git
cd linear-cli-go

# Setup your Linear API key
export LINEAR_API_KEY="your_api_key_here"
# Or add to ~/.zshrc for persistence
echo 'export LINEAR_API_KEY="your_api_key_here"' >> ~/.zshrc
```

## Usage

To run the CLI with the Linear API:

```bash
go run main.go
```

To use mock data instead of the Linear API:

```bash
go run main.go -demo
```

To show only issues assigned to you:

```bash
go run main.go -mine
```

To filter issues by project ID:

```bash
go run main.go -project="proj_123"
```

For help:

```bash
go run main.go -help
```

## Project Structure

- `cmd/` - Command-line commands and business logic
- `linear/` - Linear API client
- `models/` - Data models for Linear objects
- `mock/` - Mock data for testing
- `ui/` - User interface components for the CLI

## Future Enhancements

Future versions will include:
- Support for multiple actions on issues
- Project selection
- More advanced UI features 