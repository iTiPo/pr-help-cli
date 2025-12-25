package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		printHelp()
	case "instructions":
		printInstructions()
	case "assess-instructions":
		printAssessInstructions()
	case "fix-instructions":
		printFixInstructions()
	case "prs":
		listPRs()
	case "comments":
		handleCommentsCommand()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("PR Help CLI")
	fmt.Println("===========")
	fmt.Println()
	fmt.Println("A CLI tool to help LLM prepare pull request reports using GitHub CLI")
	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  help, -h, --help")
	fmt.Println("      Show this help message")
	fmt.Println()
	fmt.Println("  instructions")
	fmt.Println("      Get LLM instructions for report generation")
	fmt.Println()
	fmt.Println("  assess-instructions")
	fmt.Println("      Get LLM instructions for assessing comments and determining needed changes")
	fmt.Println()
	fmt.Println("  fix-instructions")
	fmt.Println("      Get LLM instructions for implementing fixes based on assessment")
	fmt.Println()
	fmt.Println("  prs")
	fmt.Println("      List open pull requests")
	fmt.Println()
	fmt.Println("  comments --pr <number> [--after-date <date>] [--after-time <datetime>]")
	fmt.Println("      Get unresolved comments for a specific PR")
	fmt.Println("      Options:")
	fmt.Println("        --pr <number>           PR number (required)")
	fmt.Println("        --after-date <date>     Filter comments after date (YYYY-MM-DD format)")
	fmt.Println("        --after-time <datetime> Filter comments after datetime (YYYY-MM-DD HH:MM:SS format)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  pr-help-cli instructions")
	fmt.Println("  pr-help-cli assess-instructions")
	fmt.Println("  pr-help-cli fix-instructions")
	fmt.Println("  pr-help-cli prs")
	fmt.Println("  pr-help-cli comments --pr 123")
	fmt.Println("  pr-help-cli comments --pr 123 --after-date 2025-12-01")
	fmt.Println("  pr-help-cli comments --pr 123 --after-time \"2025-12-01 14:30:00\"")
	fmt.Println()
}

func printInstructions() {
	fmt.Println("LLM Instructions for PR Report Generation")
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("GOAL:")
	fmt.Println("  Prepare a report about all needed pull requests.")
	fmt.Println()
	fmt.Println("STEPS:")
	fmt.Println("  1. Get list of opened pull requests")
	fmt.Println("  2. Read all comments in the pull requests")
	fmt.Println("  3. Provide a report with the comments")
	fmt.Println()
	fmt.Println("RESTRICTIONS:")
	fmt.Println("  - You shouldn't change the code but just provide the report")
	fmt.Println()
}

func printAssessInstructions() {
	fmt.Println("LLM Instructions for PR Comment Assessment")
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("GOAL:")
	fmt.Println("  Assess pull request comments and determine if code changes are needed.")
	fmt.Println()
	fmt.Println("STEPS:")
	fmt.Println("  1. Get list of opened pull requests")
	fmt.Println("  2. Read all comments in the pull requests")
	fmt.Println("  3. Assess each comment to understand:")
	fmt.Println("     - What changes are being requested")
	fmt.Println("     - Whether the comment requires code modifications")
	fmt.Println("     - The priority and impact of the requested changes")
	fmt.Println("     - Any dependencies between comments")
	fmt.Println("  4. Provide a detailed assessment report that includes:")
	fmt.Println("     - Summary of comments requiring action")
	fmt.Println("     - Recommended changes to make")
	fmt.Println("     - Priority order for addressing comments")
	fmt.Println("     - Estimated complexity of each change")
	fmt.Println()
	fmt.Println("RESTRICTIONS:")
	fmt.Println("  - You should assess and recommend changes, not implement them")
	fmt.Println("  - Focus on actionable insights from the comments")
	fmt.Println()
}

func printFixInstructions() {
	fmt.Println("LLM Instructions for Implementing PR Comment Fixes")
	fmt.Println("====================================================")
	fmt.Println()
	fmt.Println("GOAL:")
	fmt.Println("  Implement code changes based on the assessment from 'assess-instructions' command.")
	fmt.Println()
	fmt.Println("PREREQUISITES:")
	fmt.Println("  - You must have already run 'assess-instructions' to get the assessment report")
	fmt.Println("  - Review the assessment to understand what changes are needed")
	fmt.Println()
	fmt.Println("STEPS:")
	fmt.Println("  1. Review the assessment report from the previous step")
	fmt.Println("  2. For each PR that needs changes:")
	fmt.Println("     a. Checkout the relevant branch for that PR")
	fmt.Println("     b. Make the necessary code changes based on the comments")
	fmt.Println("     c. Compile the code to verify changes using:")
	fmt.Println("        ./gradlew wallet_solution:compileDebugSources -q")
	fmt.Println("     d. If compilation succeeds, commit the changes with a descriptive message")
	fmt.Println("     e. If compilation fails, fix the errors and repeat from step c")
	fmt.Println("     f. Repeat for all changes in this PR")
	fmt.Println("  3. After all changes are committed, provide a summary of:")
	fmt.Println("     - What was changed in each commit")
	fmt.Println("     - Which comments were addressed")
	fmt.Println("     - Any issues encountered")
	fmt.Println()
	fmt.Println("RESTRICTIONS:")
	fmt.Println("  - DO commit changes after successful compilation")
	fmt.Println("  - DO NOT push changes to remote repository")
	fmt.Println("  - Always compile before committing")
	fmt.Println("  - Use clear, descriptive commit messages that reference the PR comments")
	fmt.Println()
}

func handleCommentsCommand() {
	commentsCmd := flag.NewFlagSet("comments", flag.ExitOnError)
	prNumber := commentsCmd.String("pr", "", "PR number (required)")
	afterDate := commentsCmd.String("after-date", "", "Filter comments after date (YYYY-MM-DD format)")
	afterTime := commentsCmd.String("after-time", "", "Filter comments after datetime (YYYY-MM-DD HH:MM:SS format)")

	commentsCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pr-help-cli comments --pr <number> [--after-date <date>] [--after-time <datetime>]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		commentsCmd.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  pr-help-cli comments --pr 123\n")
		fmt.Fprintf(os.Stderr, "  pr-help-cli comments --pr 123 --after-date 2025-12-01\n")
		fmt.Fprintf(os.Stderr, "  pr-help-cli comments --pr 123 --after-time \"2025-12-01 14:30:00\"\n")
	}

	commentsCmd.Parse(os.Args[2:])

	if *prNumber == "" {
		fmt.Fprintf(os.Stderr, "Error: --pr flag is required\n\n")
		commentsCmd.Usage()
		os.Exit(1)
	}

	// Check if both flags are provided
	if *afterDate != "" && *afterTime != "" {
		fmt.Fprintf(os.Stderr, "Error: cannot use both --after-date and --after-time flags\n\n")
		commentsCmd.Usage()
		os.Exit(1)
	}

	listComments(*prNumber, *afterDate, *afterTime)
}

func listPRs() {
	// Execute: gh pr list --state open --json number,title,author,url,headRefName,baseRefName
	cmd := exec.Command("gh", "pr", "list", "--state", "open", "--json", "number,title,author,url,headRefName,baseRefName")
	output, err := cmd.CombinedOutput()
	if err != nil {
		errorJSON := map[string]string{
			"error":   fmt.Sprintf("failed to list PRs: %v", err),
			"details": string(output),
		}
		out, _ := json.MarshalIndent(errorJSON, "", "  ")
		fmt.Println(string(out))
		os.Exit(1)
	}
	// gh already outputs JSON, we just print it
	fmt.Println(string(output))
}

func listComments(prNumber string, afterDate string, afterTime string) {
	// Convert date/datetime to ISO8601 format if provided
	var filterDate string
	if afterDate != "" {
		// Parse the date in YYYY-MM-DD format
		parsedDate, err := time.Parse("2006-01-02", afterDate)
		if err != nil {
			errorJSON := map[string]string{
				"error": fmt.Sprintf("invalid date format: %s (expected YYYY-MM-DD)", afterDate),
			}
			out, _ := json.MarshalIndent(errorJSON, "", "  ")
			fmt.Println(string(out))
			os.Exit(1)
		}
		// Convert to ISO8601 format with time set to 00:00:00 UTC
		filterDate = parsedDate.Format(time.RFC3339)
	} else if afterTime != "" {
		// Parse the datetime in YYYY-MM-DD HH:MM:SS format
		parsedTime, err := time.Parse("2006-01-02 15:04:05", afterTime)
		if err != nil {
			errorJSON := map[string]string{
				"error": fmt.Sprintf("invalid datetime format: %s (expected YYYY-MM-DD HH:MM:SS)", afterTime),
			}
			out, _ := json.MarshalIndent(errorJSON, "", "  ")
			fmt.Println(string(out))
			os.Exit(1)
		}
		// Convert to ISO8601 format
		filterDate = parsedTime.Format(time.RFC3339)
	}

	// Get repo owner and name
	repoCmd := exec.Command("gh", "repo", "view", "--json", "owner,name")
	repoOutput, err := repoCmd.Output()
	if err != nil {
		errorJSON := map[string]string{
			"error":   fmt.Sprintf("failed to get repo info: %v", err),
			"details": string(repoOutput),
		}
		out, _ := json.MarshalIndent(errorJSON, "", "  ")
		fmt.Println(string(out))
		os.Exit(1)
	}

	var repoInfo struct {
		Name  string `json:"name"`
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
	if err := json.Unmarshal(repoOutput, &repoInfo); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing repo info: %v\n", err)
		os.Exit(1)
	}

	query := fmt.Sprintf(`
{
  repository(owner: "%s", name: "%s") {
    pullRequest(number: %s) {
      reviewThreads(first: 100) {
        nodes {
          isResolved
          comments(first: 100) {
            nodes {
              author { login }
              body
              path
              line
              createdAt
            }
          }
        }
      }
    }
  }
}`, repoInfo.Owner.Login, repoInfo.Name, prNumber)

	// Build jq filter with optional date filtering
	jqFilter := `[.data.repository.pullRequest.reviewThreads.nodes[] | select(.isResolved == false) | .comments.nodes[0] | select(. != null)`
	if filterDate != "" {
		jqFilter += fmt.Sprintf(` | select(.createdAt >= "%s")`, filterDate)
	}
	jqFilter += `] | map({
  file: .path,
  line: .line,
  author: .author.login,
  comment: .body,
  createdAt: .createdAt
})`

	cmd := exec.Command("gh", "api", "graphql", "-f", "query="+query, "--jq", jqFilter)
	output, err := cmd.CombinedOutput()
	if err != nil {
		errorJSON := map[string]string{
			"error":   fmt.Sprintf("failed to get comments for PR %s: %v", prNumber, err),
			"details": string(output),
		}
		out, _ := json.MarshalIndent(errorJSON, "", "  ")
		fmt.Println(string(out))
		os.Exit(1)
	}
	fmt.Println(string(output))
}
