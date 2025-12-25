# PR Help CLI

A CLI tool designed to help AI assistants (like Cursor) prepare pull request reports and implement fixes based on GitHub PR comments.

## Building the CLI

### Prerequisites
- Go 1.25.4 or later
- GitHub CLI (`gh`) installed and authenticated

### Build Instructions

```bash
go build -o pr-help-cli main.go
```

This will create an executable named `pr-help-cli` in your current directory.

## Adding to PATH

**It is highly recommended to add this CLI to your PATH.** This allows AI tools like Cursor to easily access and execute the CLI commands without needing to specify the full path.

### macOS/Linux

Add the following line to your shell configuration file (`~/.bashrc`, `~/.zshrc`, or `~/.profile`):

```bash
export PATH="$PATH:/path/to/pr-help-cli"
```

Or move the binary to a directory already in your PATH:

```bash
sudo mv pr-help-cli /usr/local/bin/
```

### Windows

1. Copy `pr-help-cli.exe` to a directory (e.g., `C:\tools\`)
2. Add that directory to your PATH environment variable through System Properties

## ⚠️ Important Configuration

**Before using this tool, you MUST customize the compile command at line 139 in `main.go`** to match your project's build system:

```go
fmt.Println("        ./gradlew wallet_solution:compileDebugSources -q")
```

Replace this with your actual compile command, for example:
- `go build ./...` (for Go projects)
- `npm run build` (for Node.js projects)
- `mvn compile` (for Maven projects)
- `cargo build` (for Rust projects)

You can also add additional commands if needed for your workflow (e.g., running tests, linting, etc.).

## Usage with Cursor (or other AI assistants)

Using this CLI with Cursor is simple:

### Step 1: Open Cursor
Open your project in Cursor IDE.

### Step 2: Instruct Cursor
Tell Cursor to use the CLI. For example:

**For generating a PR report:**
```
Please call `pr-help-cli instructions` and follow the instructions to generate a report about all open pull requests.
```

**For assessing PR comments:**
```
Please call `pr-help-cli assess-instructions` and follow the instructions to assess PR comments and determine what changes are needed.
```

**For implementing fixes:**
```
Please call `pr-help-cli fix-instructions` and follow the instructions to implement the fixes based on the assessment.
```

### Step 3: Relax and Wait
Sit back and let Cursor do the work! The AI will:
1. Execute the CLI to get instructions
2. Fetch PR data from GitHub
3. Analyze comments
4. Generate reports or implement fixes as requested

## Available Commands

### `help`, `-h`, `--help`
Show help message with all available commands.

```bash
pr-help-cli help
```

### `instructions`
Get LLM instructions for generating a PR report. This command outputs instructions that tell the AI to:
- List all open pull requests
- Read all comments
- Provide a comprehensive report

```bash
pr-help-cli instructions
```

### `assess-instructions`
Get LLM instructions for assessing PR comments and determining what changes are needed.

```bash
pr-help-cli assess-instructions
```

### `fix-instructions`
Get LLM instructions for implementing code changes based on the assessment.

```bash
pr-help-cli fix-instructions
```

### `prs`
List all open pull requests in the current repository.

```bash
pr-help-cli prs
```

**Output:** JSON array of PRs with number, title, author, URL, head branch, and base branch.

### `comments`
Get unresolved comments for a specific PR.

```bash
pr-help-cli comments --pr <number> [--after-date <date>] [--after-time <datetime>]
```

**Options:**
- `--pr <number>` - PR number (required)
- `--after-date <date>` - Filter comments after date (YYYY-MM-DD format)
- `--after-time <datetime>` - Filter comments after datetime (YYYY-MM-DD HH:MM:SS format)

**Examples:**
```bash
pr-help-cli comments --pr 123
pr-help-cli comments --pr 123 --after-date 2025-12-01
pr-help-cli comments --pr 123 --after-time "2025-12-01 14:30:00"
```

**Output:** JSON array of unresolved comments with file, line, author, comment text, and creation date.

## Workflow Examples

### Basic PR Report Workflow
1. Tell Cursor: "Call `pr-help-cli instructions` and generate a report"
2. Cursor will fetch all open PRs and their comments
3. You'll receive a comprehensive report of all PR activity

### Fix Implementation Workflow
1. Tell Cursor: "Call `pr-help-cli assess-instructions` and assess the PR comments"
2. Review the assessment report
3. Tell Cursor: "Call `pr-help-cli fix-instructions` and implement the fixes"
4. Cursor will checkout branches, make changes, compile, and commit
5. Review the commits and push when ready

## Output Format

All commands output data in JSON format for easy parsing by AI tools and scripts.

## Requirements

- **GitHub CLI (`gh`)**: Must be installed and authenticated
- **Git**: For repository operations
- **Project build tools**: Whatever your project uses (Gradle, Maven, npm, etc.)

## License

This tool is provided as-is for helping with PR management workflows.

