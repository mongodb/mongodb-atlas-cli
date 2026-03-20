# Atlas CLI — Agent Instructions

## Quick Reference

| Task | Command |
|------|---------|
| Install dependencies | `make setup` |
| Format code | `make fmt` |
| Lint (with auto-fix) | `make fix-lint` |
| Lint (check only) | `make lint` |
| Build binary | `make build` |
| Unit tests | `make unit-test` |
| E2E tests | `make e2e-test` |
| Generate docs | `make gen-docs` |
| Generate mocks | `make gen-mocks` |
| Generate API commands | `make gen-api-commands` |
| Run all checks | `make check` |
| Run single test | `go test -run TestName ./path/to/package/...` |

## Code Style & Conventions

- **Language**: Go. Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- **Linter**: golangci-lint v2 (config in `.golangci.yml`).
- **Branch naming**: JIRA ticket ID as branch name, e.g. `CLOUDP-12345`.
- **Commit messages**: `CLOUDP-12345: Short description`.
- **PRs**: Always include unit tests for new/changed code.
- **Mocks**: Generated with `mockgen` via `//go:generate` directives. Run `make gen-mocks` after interface changes.

## Project Layout

```
cmd/atlas/              — CLI entrypoint
internal/cli/           — Command implementations (Cobra)
internal/store/         — Backend API wrappers (atlas-sdk-go)
internal/api/           — Generated API command definitions
internal/mocks/         — Generated mocks (do not edit)
tools/                  — Code generators and internal tooling
docs/                   — Generated documentation (do not edit manually)
test/e2e/               — End-to-end tests
build/                  — Packaging and release artifacts
scripts/                — Utility scripts
```

## Key Patterns

- Commands use the [Cobra framework](https://github.com/spf13/cobra). Each command has `*Opts` struct + `Run()` method + `*Builder()` function.
- API integration uses [atlas-sdk-go](https://github.com/mongodb/atlas-sdk-go) via `internal/store/` interfaces.
- New commands are registered in `internal/cli/root/atlas/builder.go`.

## Before Submitting Changes

Ensure `.agents/skills/atlas-cli-open-pr/` is used for the full PR workflow and all failures are addressed before submitting changes.

## Further Reading

- [CONTRIBUTING.md](CONTRIBUTING.md) — full contribution guidelines
- `.agents/skills/` — on-demand skills for specialized workflows
