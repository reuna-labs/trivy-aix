//go:build !aix

// Package sqlite registers a SQLite driver under the name "sqlite" used by
// go-rpmdb and trivy-java-db (sql.Open("sqlite", ...)).
//
// On all platforms except AIX, modernc.org/sqlite is used; it registers
// itself as "sqlite" in its init().  On AIX the ncruces/go-sqlite3 driver
// is used instead; see sqlite_aix.go.
package sqlite

import _ "modernc.org/sqlite" // registers "sqlite" driver via its own init()
