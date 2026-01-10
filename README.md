# Wilee, ACME Server/Proxy

[![Go Reference][godoc-badge]][godoc]
[![Go Report Card][goreport-badge]][goreport]
[![codecov][codecov-badge]][codecov]

Wilee is darvaza's official ACME (Automatic Certificate Management Environment)
server and proxy implementation. It provides automated TLS certificate
management following the RFC 8555 protocol.

[godoc]: https://pkg.go.dev/darvaza.org/wilee
[godoc-badge]: https://pkg.go.dev/badge/darvaza.org/wilee.svg
[goreport]: https://goreportcard.com/report/darvaza.org/wilee
[goreport-badge]: https://goreportcard.com/badge/darvaza.org/wilee
[codecov]: https://codecov.io/gh/darvaza-proxy/wilee
[codecov-badge]: https://codecov.io/github/darvaza-proxy/wilee/graph/badge.svg

## Status

This project is under active development. The API and features are subject to
change.

## Development

**Requirements:** Go 1.23 or later.

For development guidelines and AI agent instructions, see
[AGENTS.md](AGENTS.md).

### Quick Start

```bash
make all    # Full build cycle (get deps, generate, tidy, build)
make test   # Run tests
make tidy   # Format and tidy (run before committing)
```

## See also

- [darvaza.org/acmefy][acmefy] - ACME client library.
- [darvaza.org/core][core] - Core utilities for darvaza projects.
- [darvaza.org/darvaza][darvaza] - Darvaza proxy server.
- [darvaza.org/resolver][resolver] - DNS resolver library.
- [darvaza.org/slog][slog] - Structured logging interface.
- [darvaza.org/x][x] - Extended utility packages.
- [Apptly Open-Source Projects][oss-apptly].

[acmefy]: https://pkg.go.dev/darvaza.org/acmefy
[core]: https://pkg.go.dev/darvaza.org/core
[darvaza]: https://pkg.go.dev/darvaza.org/darvaza
[resolver]: https://pkg.go.dev/darvaza.org/resolver
[slog]: https://pkg.go.dev/darvaza.org/slog
[x]: https://pkg.go.dev/darvaza.org/x
[oss-apptly]: https://oss.apptly.co/
