#!/bin/sh
# shellcheck disable=SC1007,SC3043 # empty assignments and local usage
# fix_whitespace.sh - Find files and fix trailing whitespace and EOF newlines
#
# Usage: fix_whitespace.sh [find arguments]
#        fix_whitespace.sh -- file1 file2 ...
#
# Automatically prunes .git, .tmp and node_modules directories
#
# Environment Variables:
#   SED - sed command to use (default: sed)
#         Set to "gsed" on macOS or "sed -i ''" for BSD compatibility
#   JOBS - number of parallel jobs to run (default: 4)
#   FILES_PER_JOB - number of files per job (default: 4)
#
# Examples:
#   fix_whitespace.sh . -name '*.go' -o -name '*.md'
#   fix_whitespace.sh src/ -name '*.js'
#   fix_whitespace.sh -- README.md LICENCE.txt
#   SED="gsed" fix_whitespace.sh . -name '*.md'
#   JOBS=8 FILES_PER_JOB=2 fix_whitespace.sh .

set -eu

JOBS="${JOBS:-4}"
FILES_PER_JOB="${FILES_PER_JOB:-4}"

# get_bytes FILE COUNT
# COUNT > 0: first COUNT bytes (head)
# COUNT < 0: last -COUNT bytes (tail)
# Outputs hex string
get_bytes() {
	local file="$1" count="$2"
	local bytes

	case "$count" in
	-*)
		bytes=$(tail -c "${count#-}" "$file") ;;
	*)
		bytes=$(head -c "$count" "$file") ;;
	esac

	printf '%s' "$bytes" | od -An -tx1 | tr -d ' \t\n'
}

# quote_arg STRING
# Outputs STRING wrapped in single quotes with embedded quotes escaped
quote_arg() {
	printf "'%s'" "$(printf '%s' "$1" | sed "s/'/'\\\\''/g")"
}

# filter_file FILE COMMAND [ARGS...]
# Applies filter command to file in-place, preserving permissions
filter_file() {
	local file="$1" tmp
	shift

	tmp=$(mktemp "fix_whitespace.XXXXXX")
	"$@" < "$file" > "$tmp"
	cat "$tmp" > "$file"
	rm "$tmp"
}

# Function to fix a single file
fix_file() {
	local file="$1"
	local first_bytes last_content keep_bytes

	# Skip if not a regular file
	[ -f "$file" ] || return 0

	# Remove UTF-8 BOM if present (EF BB BF)
	first_bytes=$(get_bytes "$file" 3)
	[ "$first_bytes" != "efbbbf" ] || filter_file "$file" tail -c +4

	# Remove trailing whitespace from each line
	${SED:-sed} -i 's/[[:space:]]*$//' "$file"

	# Leave empty files alone
	[ -s "$file" ] || return 0

	# Remove trailing empty lines
	last_content=$(grep -n . "$file" | tail -1 | cut -d: -f1) || last_content=""
	if [ -z "$last_content" ]; then
		# No content, truncate
		: > "$file"
		return 0
	fi
	keep_bytes=$(head -n "$last_content" "$file" | wc -c)
	dd if=/dev/null of="$file" bs=1 seek="$keep_bytes" 2>/dev/null

	# Ensure file ends with exactly one newline
	if [ "$(wc -l < "$file")" -lt "$last_content" ]; then
		printf '\n' >> "$file"
	fi
}

# Helper function to run find with auto-pruning
run_find() {
	local paths= quoted=

	# Collect paths until we hit a find option (starts with '-')
	while [ $# -gt 0 ]; do
		case "$1" in
		-*)
			# Found a find option, stop collecting paths
			break
			;;
		*)
			# Add path with proper escaping for spaces and special chars
			quoted=$(quote_arg "$1")
			paths="${paths:+$paths }$quoted"
			shift
			;;
		esac
	done

	# Wrap user conditions in parentheses if they exist
	[ $# -eq 0 ] || set -- \( "$@" \)
	# combine auto-pruning and user conditions
	set -- \( -type d \( -name .git -o -name .tmp -o -name node_modules \) \) -prune -o "$@" -type f
	# combine escaped paths with find options
	eval "set -- ${paths:-.} \"\$@\""

	find "$@" -print0 | xargs -0 -r "-P${JOBS:-4}" "-n${FILES_PER_JOB:-4}" "$0" --
}

if [ "${1:-}" = "--" ]; then
	# Explicit file mode
	shift
	for file; do
		fix_file "$file"
	done
else
	# Find mode with arguments
	run_find "$@"
fi
