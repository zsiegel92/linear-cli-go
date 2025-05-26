# Linear CLI Go

A Go implementation of a Linear CLI tool for interacting with Linear issues.

## Features

- Load issues from Linear API
- Filter issues by assignee (me) or project
- Display a list of issues with a basic command-line selector
- Allow selection of an issue from the list
- Output the selected issue as formatted JSON

## Installation

### Using Go Install (Recommended)

```bash
# Install the latest version
go install github.com/zsiegel92/linear-cli-go@latest

# Add Go's bin directory to your PATH (one-time setup)
echo 'export PATH="$(go env GOPATH)/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# Setup your Linear API key
export LINEAR_API_KEY="your_api_key_here"
# Or add to ~/.zshrc for persistence
echo 'export LINEAR_API_KEY="your_api_key_here"' >> ~/.zshrc
```

**Note:** After installation, you can run `linear-cli-go` from anywhere in your terminal.

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/zsiegel92/linear-cli-go.git
cd linear-cli-go
go install .
```

## Usage

To run the CLI with the Linear API:

```bash
linear-cli-go
```

Or if using the manual installation:

```bash
go run main.go
```

To use mock data instead of the Linear API:

```bash
linear-cli-go -demo
```

To show only issues assigned to you:

```bash
linear-cli-go -mine
```

To filter issues by project ID:

```bash
linear-cli-go -project="proj_123"
```

For help:

```bash
linear-cli-go -help
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

## Uninstalling

### Go Install Method

```bash
# Remove the binary
rm $(go env GOPATH)/bin/linear-cli-go

# Optionally remove the PATH addition from ~/.zshrc
# Edit ~/.zshrc and remove the line:
# export PATH="$(go env GOPATH)/bin:$PATH"

# Optionally remove the API key from ~/.zshrc
# Edit ~/.zshrc and remove the line:
# export LINEAR_API_KEY="your_api_key_here"
```

### Manual Installation Method

```bash
# Remove the binary (if you ran go install)
rm $(go env GOPATH)/bin/linear-cli-go

# Remove the source code
rm -rf linear-cli-go/
```

## Troubleshooting

### Command not found: linear-cli-go

If you get "command not found" after installation:

1. Check if the binary was installed:
   ```bash
   ls $(go env GOPATH)/bin/linear-cli-go
   ```

2. Add Go's bin directory to your PATH:
   ```bash
   echo 'export PATH="$(go env GOPATH)/bin:$PATH"' >> ~/.zshrc
   source ~/.zshrc
   ```

3. Verify PATH is set correctly:
   ```bash
   echo $PATH | grep $(go env GOPATH)/bin
   ```

### LINEAR_API_KEY not set

If you see an error about missing LINEAR_API_KEY:

1. Get your API key from https://linear.app/settings/api
2. Set it in your shell:
   ```bash
   echo 'export LINEAR_API_KEY="your_api_key_here"' >> ~/.zshrc
   source ~/.zshrc
   ```