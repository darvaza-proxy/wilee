# AGENTS.md

This file provides guidance to AI agents when working with code in this
repository. For developers and general project information, please refer to
[README.md](README.md) first.

## Repository Overview

`darvaza.org/wilee` is an ACME (Automatic Certificate Management Environment)
server and proxy implementation. It provides automated TLS certificate
management following the RFC 8555 protocol, serving as part of the darvaza
proxy ecosystem.

## Prerequisites

Before starting development, ensure you have:

- Go 1.23 or later installed (check with `go version`).
- `make` command available (usually pre-installed on Unix systems).
- `$GOPATH` configured correctly (typically `~/go`).
- Git configured for proper line endings.

## Common Development Commands

```bash
# Full build cycle (get deps, generate, tidy, build)
make all

# Run tests for all modules
make test

# Format code, tidy dependencies, and run checks (run before committing)
make tidy

# Check grammar only (without formatting)
make check-grammar

# Check spelling
make check-spelling

# Check shell scripts
make check-shell

# Clean build artifacts
make clean

# Update dependencies
make up

# Run go:generate directives
make generate

# Run tests with coverage
make coverage

# Generate Codecov configuration
make codecov
```

## Build System Features

### Whitespace and EOF Handling

The `internal/build/fix_whitespace.sh` script automatically:

- Removes UTF-8 BOM if present.
- Removes trailing whitespace from all text files.
- Removes trailing empty lines.
- Ensures files end with exactly one newline.
- Excludes binary files and version control directories.
- Integrates with `make fmt` for non-Go files.
- Supports both directory scanning and explicit file arguments.

### Markdownlint Integration

The build system includes automatic Markdown linting:

- Detects markdownlint-cli via `pnpm dlx`.
- Configuration in `internal/build/markdownlint.json`.
- 80-character line limits and strict formatting rules.
- Selective HTML allowlist (comments, br, kbd, etc.).
- Runs automatically with `make fmt` when available.

### CSpell Integration

Spell checking for both Markdown and Go source files:

- Detects cspell via `pnpm dlx`.
- British English configuration in `internal/build/cspell.json`.
- Available via `check-spelling` target.
- Integrated into `make tidy`.
- Custom word list for project-specific terminology.
- Checks both documentation and code comments.

### LanguageTool Integration

Grammar and style checking for Markdown files:

- Detects LanguageTool via `pnpm dlx`.
- British English configuration in `internal/build/languagetool.cfg`.
- Available via `make check-grammar` target.
- Checks for missing articles, punctuation, and proper hyphenation.

### ShellCheck Integration

Shell script analysis for all `.sh` files:

- Detects shellcheck via `pnpm dlx`.
- Available via `check-shell` target.
- Integrated into `make tidy`.
- Uses inline disable directives for SC1007 (empty assignments) and SC3043
  (`local` usage).
- Checks for common shell scripting issues and best practices.

### Test Coverage Collection

Automated coverage reporting across all modules:

- `make coverage` target runs tests with coverage profiling.
- Uses `internal/build/make_coverage.sh` to orchestrate testing.
- Tests each module independently with atomic coverage mode.
- Stores results in `.tmp/coverage/` directory.
- Displays coverage summary after test runs.

### Codecov Integration

Automated coverage reporting with monorepo support:

- `make codecov` target generates Codecov configuration.
- `internal/build/make_codecov.sh` creates:
  - `codecov.yml`: Dynamic configuration with per-module flags.
  - `codecov.sh`: Upload script for bulk submission.
- Module-specific coverage targets (80% default).
- Path mappings for accurate coverage attribution.
- GitHub Actions workflow automatically uploads coverage data.

## Code Architecture

### Key Design Principles

- **RFC 8555 compliance**: Follow the ACME protocol specification.
  See [docs/rfc/](docs/rfc/) for specifications and extensions.
- **Minimal dependencies**: Primarily darvaza.org packages and Go standard
  library.
- **Interface-driven design**: For extensibility and testability.
- **Structured logging**: Use `darvaza.org/slog` for consistent logging.

### Module Structure

- **Root module** (`darvaza.org/wilee`): Core ACME server/proxy library.
- **Command module** (`darvaza.org/wilee/cmd/wilee`): Server binary.

### Code Quality Standards

The project enforces strict linting rules via revive (configuration in
`internal/build/revive.toml`):

- Max function length: 40 lines.
- Max function results: 3.
- Max arguments: 5.
- Cognitive complexity: 7.
- Cyclomatic complexity: 10.

Always run `make tidy` before committing to ensure proper formatting.

### Testing Patterns

- Table-driven tests are preferred.
- Comprehensive coverage for protocol implementations is expected.
- Use `darvaza.org/core` testing utilities where applicable.

## Testing with GOTEST_FLAGS

The `GOTEST_FLAGS` environment variable allows flexible test execution by
passing additional flags to `go test`. This variable is defined in the
Makefile with an empty default value and is used when running tests through
the generated rules.

### Common Usage Examples

```bash
# Run tests with race detection
make test GOTEST_FLAGS="-race"

# Run specific tests by pattern
make test GOTEST_FLAGS="-run TestSpecific"

# Generate coverage profile (alternative to 'make coverage')
make test GOTEST_FLAGS="-coverprofile=coverage.out"

# Run tests with timeout
make test GOTEST_FLAGS="-timeout 30s"

# Combine multiple flags
make test GOTEST_FLAGS="-v -race -coverprofile=coverage.out"

# Run benchmarks
make test GOTEST_FLAGS="-bench=. -benchmem"

# Skip long-running tests
make test GOTEST_FLAGS="-short"

# Test with specific CPU count
make test GOTEST_FLAGS="-cpu=1,2,4"
```

### How It Works

1. The Makefile defines `GOTEST_FLAGS ?=` (empty by default).
2. The generated rules use it in the test target:
   `$(GO) test $(GOTEST_FLAGS) ./...`.
3. Any flags passed via `GOTEST_FLAGS` are forwarded directly to `go test`.

This provides a clean interface for passing arbitrary test flags without
modifying the Makefile, making it easy to run tests with different
configurations for debugging, coverage analysis, or CI/CD pipelines.

## Important Notes

- Go 1.23 is the minimum required version.
- The Makefile dynamically generates rules for subprojects.
- Tool versions (golangci-lint, revive) are selected based on Go version.
- Always use `pnpm` instead of `npm` for any JavaScript/TypeScript tooling.
- Follow existing patterns when adding new functionality.

## Linting and Code Quality

### Documentation Standards

When editing Markdown files, ensure compliance with:

- **Line Length**: Maximum 80 characters per line.
- **LanguageTool**: Check for missing articles ("a", "an", "the"), punctuation,
  and proper hyphenation of compound modifiers.
- **Markdownlint**: Follow standard Markdown formatting rules.

### Common Documentation Issues to Check

1. **Missing Articles**: Ensure proper use of "a", "an", and "the".
   - ❌ "provides simple interface for certificates".
   - ✅ "provides a simple interface for certificates".

2. **Missing Punctuation**: End all list items consistently.
   - ❌ "Comprehensive coverage is expected".
   - ✅ "Comprehensive coverage is expected.".

3. **Compound Modifiers**: Hyphenate when used as modifiers.
   - ❌ "protocol compliant implementation".
   - ✅ "protocol-compliant implementation".

### Writing Documentation Guidelines

When creating or editing documentation files:

1. **File Structure**:
   - Always include a link to related documentation (e.g., AGENTS.md should
     link to README.md).
   - Add prerequisites or setup instructions before diving into commands.
   - Include paths to configuration files when mentioning tools.

2. **Formatting Consistency**:
   - End all bullet points with periods for consistency.
   - Capitalise proper nouns correctly (Go, GitHub, Makefile).
   - Use consistent punctuation in examples and lists.

3. **Clarity and Context**:
   - Provide context for AI agents and developers alike.
   - Include "why" explanations, not just "what" descriptions.
   - Add examples for complex concepts or common pitfalls.

4. **Maintenance**:
   - Update documentation when adding new features or changing workflows.
   - Keep the pre-commit checklist current with project practices.
   - Review documentation changes for the issues listed above.

### Pre-commit Checklist

1. **ALWAYS run `make tidy` first** - Fix ALL issues before committing:
   - Go code formatting and whitespace clean-up.
   - Markdown files checked with CSpell and markdownlint.
   - Shell scripts checked with ShellCheck.
   - If `make tidy` fails, fix the issues and run it again until it passes.
2. Verify all tests pass with `make test`.
3. Ensure no linting violations remain.
4. Update `AGENTS.md` to reflect any changes in development workflow or
   standards.
5. Update `README.md` to reflect significant changes in functionality or API.

## CI/CD and Code Analysis

### DeepSource Configuration

The project uses DeepSource for static code analysis. Configuration is in the
`.deepsource.toml` file:

- Shell analyser is configured for POSIX sh dialect.
- To ignore specific issues for certain files, use `[[issues]]` blocks with
  `paths` (not `exclude_patterns`).
- Common shell issues:
  - SH-1091: "local is undefined in POSIX sh" - excluded for all .sh files.

### GitHub Actions

- **Build workflow** (`.github/workflows/build.yml`): Compilation validation.
- **Test workflow** (`.github/workflows/test.yml`): Multi-version testing.
- **Race workflow** (`.github/workflows/race.yml`): Race condition detection.
- **Codecov workflow** (`.github/workflows/codecov.yml`): Coverage reporting.
- **Renovate workflow** (`.github/workflows/renovate.yml`): Dependency updates.
- Workflows skip branches ending in `-wip`.
- All CI checks must pass before merging PRs.

### Working with Build Tools

When LanguageTool reports issues:

- Custom dictionary is auto-generated from CSpell words in
  `.tmp/languagetool-dict.txt`.
- Technical terms should be added to `internal/build/cspell.json`.
- False positives for code-related punctuation are disabled in
  `languagetool.cfg`.

## Troubleshooting

### Common Issues and Solutions

1. **LanguageTool false positives**:
   - Add technical terms to `internal/build/cspell.json`.
   - Dictionary will auto-regenerate on next `make check-grammar`.
   - For persistent issues, consider adding rules to `languagetool.cfg`.

2. **DeepSource shell issues**:
   - Use ShellCheck disable comments for specific lines.
   - Update `.deepsource.toml` with issue-specific `paths` configurations.
   - Remember: DeepSource uses `paths`, not `exclude_patterns` in
     `[[issues]]` blocks.

3. **Coverage collection failures**:
   - Ensure `.tmp/index` exists by running `make .tmp/index`.
   - Check that all modules have test files.
   - Use `GOTEST_FLAGS` to pass additional flags to tests.

4. **Linting tool detection**:
   - Tools are auto-detected via `pnpm dlx`.
   - If tools aren't found, they're replaced with `true` (no-op).
   - Install tools globally with `pnpm install -g <tool>` if needed.
