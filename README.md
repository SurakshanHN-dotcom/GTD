# GTD
GTD which stands for "Getting Things Done" is a lightweight cli task manager written in Go. It allows users to manage simple to-do items directly from the terminal with persistent storage.

## Project Structure

```text
cmd/        - CLI entry point (main package)
internal/   - Core logic and storage layer
data/       - Runtime data (ignored in version control)
```


## Usage Process

### Build
Build the CLI binary from the project root:
```bash
go build -o gtd ./cmd
```

### Help
All available commands can be viewed using
```bash
gtd --help      
```
