# Claude Code Guidelines

## Repository Structure
This is a Go REST API. Files are organized as follows:
- `main.go` — entry point, HTTP server, route registration
- `users.go` — user handlers (package handlers)
- `products.go` — product handlers (package handlers)
- `logger.go` — logger middleware (package middleware)

## Code Conventions
- Follow existing handler patterns (see `products.go` for reference)
- All handlers should handle unsupported methods with `http.StatusMethodNotAllowed`
- Use `encoding/json` for all JSON responses
- Set `Content-Type: application/json` on all JSON responses

## Pull Request Requirements
- Always use the PULL_REQUEST_TEMPLATE.md when creating PRs
- The template is at `.github/PULL_REQUEST_TEMPLATE.md`
- Fill in every section — do not leave placeholders
- Link the Linear issue URL in the Tickets section
- Describe testing steps clearly in "How do we test your code?"
- Set Deadline to "No" unless specified in the issue

## Branch Naming
- Format: `feature/linear-id-short-description`
- Example: `feature/ai-7-health-endpoint`

## Testing
- Verify the package builds with `go build` before opening the PR