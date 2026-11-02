//go:build aix && !sqlite3_dotlk

// Building trivy for AIX without the sqlite3_dotlk tag produces a binary
// that compiles cleanly but fails every SQLite open at runtime with
// "disk I/O error".  Catch it here instead.
//
// Fix: add -tags sqlite3_dotlk to your build command.
package sqlite

// sqlite3_dotlk_required_on_aix is intentionally undefined.
// If you see this error, rebuild with: -tags sqlite3_dotlk
var _ = sqlite3_dotlk_required_on_aix //nolint:revive
