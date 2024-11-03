module darvaza.org/wilee

go 1.21

replace (
	darvaza.org/sidecar => ../sidecar
	darvaza.org/x/tls => ../x/tls
)

require (
	darvaza.org/core v0.15.1 // indirect
	darvaza.org/slog v0.5.14
)

require github.com/shaj13/raft v0.0.0-20240423105203-6ba83c31e046

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
